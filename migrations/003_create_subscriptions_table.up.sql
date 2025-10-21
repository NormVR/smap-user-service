CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    follower_id UUID NOT NULL,
    following_id UUID NOT NULL,

    -- Статус подписки (pending для приватных аккаунтов)
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'pending', 'blocked')),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    -- Уникальность связи
    UNIQUE(follower_id, following_id),

    -- Внешние ключи
    CONSTRAINT fk_follower FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_following FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Индексы для быстрого доступа к подпискам
CREATE INDEX IF NOT EXISTS idx_subscriptions_follower_id ON subscriptions(follower_id);
CREATE INDEX IF NOT EXISTS idx_subscriptions_following_id ON subscriptions(following_id);
CREATE INDEX IF NOT EXISTS idx_subscriptions_status ON subscriptions(status);