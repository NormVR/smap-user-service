CREATE TABLE IF NOT EXISTS users
(
    id             UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    email          VARCHAR(255) UNIQUE NOT NULL,
    password_hash  VARCHAR(255)        NOT NULL,
    email_verified BOOLEAN                  DEFAULT false,
    is_active      BOOLEAN                  DEFAULT true,
    mfa_enabled    BOOLEAN                  DEFAULT false,
    mfa_secret     VARCHAR(100),
    last_login     TIMESTAMP WITH TIME ZONE,
    login_attempts INTEGER                  DEFAULT 0,
    locked_until   TIMESTAMP WITH TIME ZONE,
    created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW()

        CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);