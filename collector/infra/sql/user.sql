-- name: FindCandidates :many
SELECT
    u.id AS u_id,
    u.last_collected_at AS u_last_collected_at,
    u.github_id AS u_github_id,
    u.github_username AS u_github_username,
    gi.access_token AS gi_access_token,
    gi.refresh_token AS gi_refresh_token,
    gi.scopes AS gi_scopes,
    gi.last_processed_event_id AS gi_last_processed_event_id,
    si.access_token AS si_access_token,
    si.refresh_token AS si_refresh_token,
    si.scopes AS si_scopes,
    si.slack_user_id AS si_slack_user_id,
    si.team_id AS si_team_id,
    ARRAY_AGG(sic.channel_id) AS si_channel_ids
FROM
    users u
LEFT JOIN
    github_integrations gi ON u.id = gi.user_id
LEFT JOIN
    slack_integrations si ON u.id = si.user_id
LEFT JOIN
    slack_integration_channels sic ON si.id = sic.slack_integration_id
WHERE
    u.id NOT IN (sqlc.slice('exclude_ids'))
GROUP BY
    u.id, gi.id, si.id
ORDER BY
    u.last_collected_at ASC
LIMIT
    sqlc.arg('lim');
