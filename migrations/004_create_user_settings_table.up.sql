CREATE TABLE IF NOT EXISTS user_settings (
    user_id UUID PRIMARY KEY,

    -- Уведомления
    email_notifications BOOLEAN DEFAULT true,
    push_notifications BOOLEAN DEFAULT true,
    sms_notifications BOOLEAN DEFAULT false,

    -- Приватность
    show_online_status BOOLEAN DEFAULT true,
    show_last_seen BOOLEAN DEFAULT true,
    show_followers_count BOOLEAN DEFAULT true,
    show_following_count BOOLEAN DEFAULT true,

    -- Безопасность
    two_factor_auth BOOLEAN DEFAULT false,
    login_alerts BOOLEAN DEFAULT true,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);