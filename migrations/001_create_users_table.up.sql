CREATE TABLE IF NOT EXISTS users_profile (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    bio TEXT,
    avatar_url VARCHAR(500),
    website VARCHAR(500),
    location VARCHAR(200),
    birth_date DATE,
    gender VARCHAR(20),
    phone VARCHAR(50),

    -- Настройки приватности
    is_public BOOLEAN DEFAULT true,
    allow_messages_from_anyone BOOLEAN DEFAULT true,

    -- Статистика
    followers_count INTEGER DEFAULT 0,
    following_count INTEGER DEFAULT 0,
    posts_count INTEGER DEFAULT 0,

    -- Таймстампы
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_seen_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    -- Индексы
    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
    );

-- Индексы для быстрого поиска
CREATE INDEX IF NOT EXISTS idx_users_email ON users_profile(email);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users_profile(created_at);