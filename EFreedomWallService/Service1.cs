using System;
using CoreWCF;
using Microsoft.Data.Sqlite;

namespace EFreedomWallService
{
    [ServiceBehavior(IncludeExceptionDetailInFaults = true)]
    public class Service1 : IService1
    {
        private SqliteConnection newConnection()
        {
            return new SqliteConnection("Data Source=freedom_wall.db");
        }

        public bool CreatePost(Post post)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "INSERT INTO posts (content, poster, password) VALUES ($content, $poster, $password);";
                cmd.Parameters.AddWithValue("$content", post.Content);
                cmd.Parameters.AddWithValue("$poster", post.Poster);
                Console.Out.WriteLine("PASSWORD " + post.Password);
                cmd.Parameters.AddWithValue("$password", string.IsNullOrEmpty(post.Password) ? System.DBNull.Value : post.Password);
                var n = cmd.ExecuteNonQuery();
                if (n == 0)
                {
                    return false;
                }

                connection.Close();
                return true;
            }
        }

        public Posts RecentPosts(int n, int page = 1)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var totalPages = 1;
                var posts = new Posts()
                {
                    CurrentPage = page
                };

                // for list count
                var countCmd = new SqliteCommand("SELECT COUNT(*) FROM posts", connection);
                using (var reader = countCmd.ExecuteReader())
                {
                    if (reader.Read())
                    {

                        var total = reader.GetInt32(0);
                        totalPages = (int)Math.Floor((decimal)(total / n));
                        if (total % n > 0)
                        {
                            totalPages++;
                        }

                        posts.Total = total;
                    }
                    if (page + 1 <= totalPages)
                    {
                        posts.NextPage = page + 1;
                    }

                    posts.PrevPage = page - 1;
                }

                // for post list
                var cmd = new SqliteCommand(
                    "SELECT id, content, poster, password, created_at FROM posts ORDER BY created_at DESC LIMIT $count OFFSET $page",
                    connection
                   );
                cmd.Parameters.AddWithValue("$count", n);
                cmd.Parameters.AddWithValue("$page", Math.Max(page - 1, 0) * n);

                using (var reader = cmd.ExecuteReader())
                {
                    posts.FromReader(reader);
                }

                connection.Close();
                return posts;
            }
        }

        public Post GetPost(int id, string? password = null)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "SELECT id, content, poster, password, created_at FROM posts WHERE id = $id";
                cmd.Parameters.AddWithValue("$id", id);

                using (var reader = cmd.ExecuteReader())
                {
                    if (!reader.Read())
                        throw new FaultException("Post not found!");

                    var idIndex = reader.GetOrdinal("id");
                    var contentIndex = reader.GetOrdinal("content");
                    var posterIndex = reader.GetOrdinal("poster");
                    var createdAtIndex = reader.GetOrdinal("created_at");
                    var passwordIndex = reader.GetOrdinal("password");

                    var post = new Post
                    {
                        Id = reader.GetInt32(idIndex),
                        Content = reader.GetString(contentIndex),
                        Poster = !reader.IsDBNull(posterIndex) ? reader.GetString(posterIndex) : null,
                        CreatedAt = reader.GetDateTime(createdAtIndex),
                        Password = !reader.IsDBNull(passwordIndex) ? reader.GetString(passwordIndex) : null,
                    };

                    post.Unlock(password);
                    connection.Close();
                    return post;
                }
            }
        }

        public bool UpdatePost(Post post, string? password = null)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                var passClause = password == null || password.Length == 0 ? "password IS NULL" : "password = $password";
                var newPassClause = post.Password != null && post.Password.Length > 0 ? ", password = $new_password " : "";
                cmd.CommandText = $"UPDATE posts SET content = $content, poster = $poster {newPassClause} WHERE id = $id AND {passClause}";
                cmd.Parameters.AddWithValue("$id", post.Id);
                cmd.Parameters.AddWithValue("$content", post.Content);
                cmd.Parameters.AddWithValue("$poster", post.Poster);
                cmd.Parameters.AddWithValue("$new_password", post.Password);
                cmd.Parameters.AddWithValue("$password", post.Password);
                var n = cmd.ExecuteNonQuery();

                connection.Close();
                return n != 0;
            }
        }

        public bool DeletePost(int id, string? password = null)
        {
           using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();

                // delete first post views then post likes
                var deletePostViewsCmd = new SqliteCommand("DELETE FROM post_views WHERE post_id = $id", connection);
                deletePostViewsCmd.Parameters.AddWithValue("$id", id);
                deletePostViewsCmd.ExecuteNonQuery();

                var deletePostLikesCmd = new SqliteCommand("DELETE FROM post_likes WHERE post_id = $id", connection);
                deletePostLikesCmd.Parameters.AddWithValue("$id", id);
                deletePostLikesCmd.ExecuteNonQuery();

                var passClause = password == null || password.Length == 0 ? "password IS NULL" : "password = $password";
                cmd.CommandText = $"DELETE FROM posts WHERE id = $id AND {passClause}";
                cmd.Parameters.AddWithValue("$id", id);
                cmd.Parameters.AddWithValue("$password", password);

                var n = cmd.ExecuteNonQuery();
                connection.Close();
                return n != 0;
            }
        }

        public int RecordView(string sessionId, int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "INSERT INTO post_views (session_id, post_id) VALUES ($session_id, $post_id)";
                cmd.Parameters.AddWithValue("$session_id", sessionId);
                cmd.Parameters.AddWithValue("$post_id", postId);
                cmd.ExecuteNonQuery();
                connection.Close();
                return GetPostViews(postId);
            }
        }

        public int LikePost(string sessionId, int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "INSERT INTO post_likes (session_id, post_id) VALUES ($session_id, $post_id)";
                cmd.Parameters.AddWithValue("$session_id", sessionId);
                cmd.Parameters.AddWithValue("$post_id", postId);
                cmd.ExecuteNonQuery();
                connection.Close();
                return GetPostLikes(postId);
            }
        }

        public int GetPostLikes(int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "SELECT COUNT(*) FROM post_likes WHERE post_id = $id";
                cmd.Parameters.AddWithValue("$id", postId);

                using (var reader = cmd.ExecuteReader())
                {
                    var likesCount = 0;
                    if (reader.Read() && !reader.IsDBNull(0))
                    {
                        likesCount = reader.GetInt32(0);
                    }
                    connection.Close();
                    return likesCount;
                }

            }
        }

        public int GetPostViews(int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "SELECT COUNT(*) FROM post_views WHERE post_id = $id";
                cmd.Parameters.AddWithValue("$id", postId);

                using (var reader = cmd.ExecuteReader())
                {
                    var viewsCount = 0;
                    if (reader.Read() && !reader.IsDBNull(0))
                    {
                        viewsCount = reader.GetInt32(0);
                    }
                    connection.Close();
                    return viewsCount;
                }

            }
        }

        public bool IsPostLocked(int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "SELECT id FROM posts WHERE id = $id AND password IS NOT NULL";
                cmd.Parameters.AddWithValue("$id", postId);

                using (var reader = cmd.ExecuteReader())
                {
                    var isLocked = false;
                    if (reader.Read())
                    {
                        isLocked = true;
                    }

                    connection.Close();
                    return isLocked;
                }
            }
        }

        public bool IsPostLiked(string sessionId, int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "SELECT COUNT(*) FROM post_likes WHERE post_id = $id AND session_id = $sessionId";
                cmd.Parameters.AddWithValue("$id", postId);
                cmd.Parameters.AddWithValue("$sessionId", sessionId);
                var isLiked = false;

                using (var reader = cmd.ExecuteReader())
                {
                    if (reader.Read() && reader.GetInt32(0) == 1)
                    {
                        isLiked = true;
                    }
                    connection.Close();
                }
                return isLiked;
            }
        }

        public int UnlikePost(string sessionId, int postId)
        {
            using (var connection = newConnection())
            {
                connection.Open();
                var cmd = connection.CreateCommand();
                cmd.CommandText = "DELETE FROM post_likes WHERE post_id = $post_id AND session_id = $session_id";
                cmd.Parameters.AddWithValue("$session_id", sessionId);
                cmd.Parameters.AddWithValue("$post_id", postId);
                cmd.ExecuteNonQuery();
                connection.Close();
                return GetPostLikes(postId);
            }
        }
    }
}

