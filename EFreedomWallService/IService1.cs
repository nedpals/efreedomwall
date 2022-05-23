using CoreWCF;
using System.Runtime.Serialization;
using System.Diagnostics.CodeAnalysis;

namespace EFreedomWallService
{
	[ServiceContract]
	public interface IService1
	{
		[OperationContract]
		public bool IsPostLocked(int postId);

		[OperationContract]
		public bool CreatePost(Post post);

		[OperationContract]
		public Post GetPost(int id, string? password = null);

		[OperationContract]
		public bool UpdatePost(Post post, string? password = null);

		[OperationContract]
		public bool DeletePost(int id, string? password = null);

		[OperationContract]
		public int RecordView(string sessionId, int postId);

		[OperationContract]
		public int LikePost(string sessionId, int postId);

		[OperationContract]
		public int UnlikePost(string sessionId, int postId);

		[OperationContract]
		public bool IsPostLiked(string sessionId, int postId);

		[OperationContract]
		public int GetPostLikes(int postId);

		[OperationContract]
		public int GetPostViews(int postId);

        [OperationContract]
		public Posts RecentPosts(int howMany, int page = 1);
	}

	[DataContract]
	public class Posts
    {
		private List<Post> _posts = new List<Post>();
		private int _total = 0;
		private int _currentPage = 1;
		private int? _prevPage;
		private int? _nextPage;

		[DataMember]
		public List<Post> Results
        {
			get => _posts;
        }

		[DataMember]
		public int Total
        {
			get => _total;
			set => _total = Math.Max(value, 0);
        }

		[DataMember]
		public int CurrentPage
        {
			get => _currentPage;
			set => _total = Math.Max(value, 1);
        }

		[DataMember]
		[AllowNull]
		public int? PrevPage
        {
			get => _prevPage != null && _prevPage == 0 ? null : _prevPage;
			set => _prevPage = value == null ? value : Math.Max(value ?? 1, 1);
        }

		[DataMember]
		[AllowNull]
		public int? NextPage
		{
			get => _nextPage != null && _nextPage == 0 ? null : _nextPage;
			set => _nextPage = value == null ? value : Math.Max(value ?? 1, 1);
		}

		public void FromReader(Microsoft.Data.Sqlite.SqliteDataReader reader)
        {
			while (reader.Read())
			{
				var idIndex = reader.GetOrdinal("id");
				var contentIndex = reader.GetOrdinal("content");
				var posterIndex = reader.GetOrdinal("poster");
				var createdAtIndex = reader.GetOrdinal("created_at");
				var passwordIndex = reader.GetOrdinal("password");

				var hasPassword = false;
				if (!reader.IsDBNull(passwordIndex))
                {
					var password = reader.GetString(passwordIndex);
					hasPassword = password.Length != 0;
				}

				_posts.Add(new Post
				{
					Id = reader.GetInt32(idIndex),
					Content = !hasPassword ? reader.GetString(contentIndex) : "",
					Poster = !hasPassword && !reader.IsDBNull(posterIndex) ? reader.GetString(posterIndex) : null,
					CreatedAt = reader.GetDateTime(createdAtIndex),
					IsLocked = hasPassword
				});
			}
		}
    }

	[DataContract]
	public class Post
    {
		private int _id = -1;
		private DateTime _createdAt = DateTime.Now;
		private string _content = "";
		private bool _isLocked = false;

		[AllowNull]
		private string _poster;

		[AllowNull]
		private string _password;

		[DataMember]
		public int Id
        {
			get => _id;
			set => _id = value;
        }

		[DataMember]
		public string Content {
			get => _content;
			set => _content = value;
		}

		[DataMember]
		public DateTime CreatedAt
        {
			get => _createdAt;
			set => _createdAt = value;
        }

		[DataMember]
		[AllowNull]
		public string Poster
        {
			get => _poster;
			set => _poster = value;
        }

		[DataMember]
		[AllowNull]
		public string Password
        {
			get => _password;
			set => _password = value;
        }

		[DataMember]
		public bool IsLocked
        {
			get => _isLocked;
			set => _isLocked = value;
        }

		public void Unlock(string? password)
        {
			if (_password == null || _password.Length == 0)
				return;

			if (password != _password)
            {
				throw new FaultException("Password mismatch.");
            }

			_password = null;
        }
    }

}

