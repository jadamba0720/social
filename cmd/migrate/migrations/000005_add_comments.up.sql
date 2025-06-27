CREATE TABLE IF NOT EXISTS comments (
  id BIGSERIAL PRIMARY KEY,
  post_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  content TEXT,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
