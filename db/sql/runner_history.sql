-- name: SaveRunnerHistory :exec
INSERT INTO runner_histories (
    id,
    runner_id,
    runner_name,
    status,
    message,
    started_at,
    ended_at,
    response_time_ms,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, now()
);

-- name: FindRunnerHistoriesByRunnerID :many
SELECT
    id,
    runner_id,
    runner_name,
    status,
    message,
    started_at,
    ended_at,
    response_time_ms,
    created_at
FROM runner_histories
WHERE runner_id = $1
ORDER BY created_at DESC;

-- name: SearchRunnerHistories :many
SELECT
    rh.id,
    rh.runner_id,
    rh.runner_name,
    rh.status,
    rh.message,
    rh.started_at,
    rh.ended_at,
    rh.response_time_ms,
    rh.created_at
FROM runner_histories AS rh
JOIN runners AS r ON rh.runner_id = r.id
WHERE 
    r.user_id = $1
    AND rh.status = $2
    AND rh.created_at BETWEEN (NOW() - (sqlc.arg('minutes')::int * INTERVAL '1 minute')) AND NOW()
ORDER BY rh.created_at DESC;;