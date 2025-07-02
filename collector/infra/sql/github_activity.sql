-- name: CreateGithubActivities :copyfrom
INSERT INTO github_activities (user_id, event_id, type, url, repository_name, payload, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateGithubIntegrationLastProcessedEventID :exec
UPDATE github_polling_states
SET last_processed_event_id = $1, updated_at = $2
WHERE user_id = $3;

-- name: UpdateGithubIntegrationToken :exec
UPDATE github_integrations
SET access_token = $1, refresh_token = $2, updated_at = $3
WHERE user_id = $4;