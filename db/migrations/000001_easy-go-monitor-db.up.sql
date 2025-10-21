CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS monitor_groups (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    is_active BOOLEAN,
    description TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE (user_id, name)
);

CREATE TABLE IF NOT EXISTS monitors (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id UUID NULL REFERENCES monitor_groups(id) ON DELETE SET NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    url TEXT NOT NULL,
    method VARCHAR(10) NOT NULL CHECK (method IN ('GET', 'POST', 'PUT', 'DELETE', 'HEAD')),
    timeout_ms int DEFAULT 5000,
    is_active BOOLEAN DEFAULT false,
    header JSONB,
    body text,
    expected_status INT DEFAULT 200,
    description TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS runners (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    monitor_id UUID NULL REFERENCES monitors(id) ON DELETE SET NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    region VARCHAR(50) NOT NULL,
    interval_second int NOT NULL,
    is_active BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS runner_histories (
    id UUID PRIMARY KEY,
    runner_id uuid NOT NULL REFERENCES runners(id),
    status VARCHAR(32) NOT NULL CHECK (status IN ('success', 'failure', 'timeout', 'error')),
    message TEXT,
    started_at TIMESTAMPTZ DEFAULT now() ,
    ended_at TIMESTAMPTZ NULL,
    duration_ms int NULL,
    response_time_ms int NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS notifiers (
    id int PRIMARY KEY,
    type VARCHAR(50) NOT NULL UNIQUE,
    display_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY,
    runner_id uuid NOT NULL REFERENCES runners(id) ON DELETE CASCADE,
    notifier_id int NOT NULL REFERENCES notifiers(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    trigger VARCHAR(50) NOT NULL DEFAULT 'on_failure' CHECK (trigger IN ('on_failure', 'on_recovery', 'always')),
    message TEXT NOT NULL,
    description TEXT DEFAULT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE INDEX idx_monitors_user_id ON monitors(user_id);
CREATE INDEX idx_runners_user_id ON runners(user_id);
CREATE INDEX idx_runner_histories_runner_id ON runner_histories(runner_id);