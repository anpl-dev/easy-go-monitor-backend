-- name: SaveRunnerHistory :exec
INSERT INTO runner_histories (
    id,
    runner_id,
    status,
    message,
    started_at,
    ended_at,
    duration_ms,
    response_time_ms,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, now()
);

-- name: FindHistoryByRunnerID :many
SELECT
    id,
    runner_id,
    status,
    message,
    started_at,
    ended_at,
    duration_ms,
    response_time_ms,
    created_at
FROM runner_histories
WHERE runner_id = $1
ORDER BY created_at DESC;