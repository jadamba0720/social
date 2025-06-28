CREATE TABLE IF NOT EXISTS followers (
    user_id bigint NOT NULL,
    followers_id bigint NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, followers_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (followers_id) REFERENCES users (id) ON DELETE CASCADE
);

