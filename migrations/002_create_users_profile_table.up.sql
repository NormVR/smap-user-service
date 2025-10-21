CREATE TABLE IF NOT EXISTS user_profiles
(
    id                         UUID PRIMARY KEY,
    firstname                  VARCHAR(100),
    lastname                   VARCHAR(100),
    username                   VARCHAR(100) UNIQUE NOT NULL,
    bio                        TEXT,
    avatar_url                 VARCHAR(500),
    website                    VARCHAR(500),
    location                   VARCHAR(200),
    birth_date                 DATE,
    gender                     VARCHAR(20),
    telephone                  VARCHAR(50),

    -- Privacy
    is_public                  BOOLEAN                  DEFAULT true,
    allow_messages_from_anyone BOOLEAN                  DEFAULT true,

    -- Statistic
    followers_count            INTEGER                  DEFAULT 0,
    following_count            INTEGER                  DEFAULT 0,
    posts_count                INTEGER                  DEFAULT 0,

    -- Timestamps
    created_at                 TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at                 TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_seen_at               TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_users_username ON user_profiles (username);