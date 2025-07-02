CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL UNIQUE,
    github_username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    avatar_url TEXT NOT NULL,
    user_role VARCHAR(50) NOT NULL DEFAULT 'Member',
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    last_collected_at BIGINT NOT NULL
);

CREATE TABLE teams (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE team_members (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    created_at BIGINT NOT NULL,
    UNIQUE (user_id, team_id)
);

CREATE TABLE github_integrations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    access_token BYTEA NOT NULL,
    refresh_token BYTEA NOT NULL,
    scopes TEXT,
    github_username INTEGER NOT NULL REFERENCES users(github_username) ON DELETE CASCADE,
    last_processed_event_id TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE slack_integrations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    access_token BYTEA NOT NULL,
    refresh_token BYTEA NOT NULL,
    scopes TEXT,
    slack_user_id VARCHAR(255) NOT NULL UNIQUE,
    team_id VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE slack_integration_channels (
    id BIGSERIAL PRIMARY KEY,
    slack_integration_id BIGINT NOT NULL REFERENCES slack_integrations(id) ON DELETE CASCADE,
    channel_id VARCHAR(255) NOT NULL
);

CREATE TABLE trello_integrations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    secret TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE TABLE github_activities (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    event_id VARCHAR(255) NOT NULL UNIQUE,
    type VARCHAR(100) NOT NULL,
    url TEXT NOT NULL,
    repository_name VARCHAR(255),
    
    -- APIで取得した生データ
    payload JSONB NOT NULL,

    created_at BIGINT NOT NULL
);

CREATE TABLE github_polling_states (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    -- 最後に処理したイベントのID（文字列）を保存します。
    -- 差分イベントを正しく識別し、重複処理を防ぐために使用します。
    -- 初回実行時にはNULLになります。
    last_processed_event_id TEXT,

    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);


CREATE TABLE slack_activities (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    slack_user_id VARCHAR(255) REFERENCES slack_integrations(slack_user_id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    url TEXT NOT NULL,
    channel_id VARCHAR(100),
    message_timestamp VARCHAR(100),
    created_at BIGINT NOT NULL
);

CREATE TABLE trello_activities (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    url TEXT NOT NULL,
    board_id VARCHAR(255),
    list_id VARCHAR(255),
    card_id VARCHAR(255),
    created_at BIGINT NOT NULL
);

CREATE TABLE private_reflections (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    reflection_date DATE NOT NULL,
    did_content TEXT,
    thought_content TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    UNIQUE (user_id, team_id, reflection_date)
);

-- 日報テーブル
CREATE TABLE daily_reports (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    report_date DATE NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'draft',
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    UNIQUE (user_id, team_id, report_date)
);
