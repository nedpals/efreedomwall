DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS post_views;
DROP TABLE IF EXISTS post_likes;

CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  content TEXT NOT NULL,
  poster TEXT,
  password TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post_views (
  session_id TEXT NOT NULL,
  post_id INTEGER NOT NULL REFERENCES posts(id)
);

CREATE TABLE post_likes (
  session_id TEXT NOT NULL,
  post_id INTEGER NOT NULL REFERENCES posts(id)
);
