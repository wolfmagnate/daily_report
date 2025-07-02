-- name: CreateSlackActivities :copyfrom
INSERT INTO slack_activities (user_id, slack_user_id, message, url, channel_id, message_timestamp, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateSlackIntegrationToken :exec
UPDATE slack_integrations
SET access_token = $1, refresh_token = $2, updated_at = $3
WHERE user_id = $4;
