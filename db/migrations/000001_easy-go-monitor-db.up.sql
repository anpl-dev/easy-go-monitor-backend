---- Initial Schema
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS monitors (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    url TEXT NOT NULL,
    type VARCHAR(50) NOT NULL,
    settings JSONB NOT NULL,
    is_enabled BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE(user_id, name)
);

CREATE TABLE IF NOT EXISTS runners (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    monitor_id UUID NOT NULL REFERENCES monitors(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    region VARCHAR(50) NOT NULL,
    interval_second INT NOT NULL,
    is_enabled BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE(user_id, name)
);

CREATE TABLE IF NOT EXISTS runner_histories (
    id UUID PRIMARY KEY,
    runner_id uuid NOT NULL REFERENCES runners(id) ON DELETE CASCADE,
    runner_name VARCHAR(32) NOT NULL,
    status VARCHAR(32) NOT NULL,
    message TEXT,
    started_at TIMESTAMPTZ NOT NULL,
    ended_at TIMESTAMPTZ,
    response_time_ms INT,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

/* CREATE TABLE IF NOT EXISTS notifiers (
    id INT PRIMARY KEY,
    type VARCHAR(50) NOT NULL UNIQUE,
    display_name VARCHAR(100) NOT NULL
) */;

/* CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY,
    runner_id UUID NOT NULL REFERENCES runners(id) ON DELETE CASCADE,
    notifier_id INT NOT NULL REFERENCES notifiers(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    trigger VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    is_enabled BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);
 */
CREATE INDEX idx_monitors_user_id ON monitors(user_id);
CREATE INDEX idx_runners_user_id ON runners(user_id);
CREATE INDEX idx_runner_histories_runner_id ON runner_histories(runner_id);
-------

/* INSERT INTO notifiers (id, type, display_name) VALUES
(1, 'email', 'Email'),
(2, 'slack', 'Slack'),
(3, 'webhook', 'Webhook')
ON CONFLICT DO NOTHING */;