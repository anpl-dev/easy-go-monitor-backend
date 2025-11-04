-- name: CreateRunner :one
INSERT INTO runners (
    id, user_id, monitor_id, name, region, interval_second, is_enabled
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;
-- name: FindRunnerByID :one
SELECT * 
FROM runners
WHERE id = $1;

-- name: FindAllRunners :many
SELECT * 
FROM runners
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateRunner :one
UPDATE runners
SET 
    monitor_id = $2,
    name = $3,
    region = $4,
    interval_second = $5,
    is_enabled = $6,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteRunner :exec
DELETE FROM runners
WHERE id = $1;
    
