-- name: CreateRunner :one
INSERT INTO runners (
    id, user_id, monitor_id, name, region, interval_second, is_active
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
    name = $2,
    region = $3,
    interval_second = $4,
    is_active = $5,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteRunner :exec
DELETE FROM runners
WHERE id = $1;
    
