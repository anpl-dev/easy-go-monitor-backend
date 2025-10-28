-- name: CreateMonitor :one
INSERT INTO monitors (
    id, user_id, group_id, name, url, type, settings, is_enabled
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: FindMonitorByID :one
SELECT * 
FROM monitors
WHERE id = $1;

-- name: FindAllMonitors :many
SELECT * 
FROM monitors
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateMonitor :one
UPDATE monitors
SET 
    group_id = $3,
    name = $4,
    url = $5,
    type = $6,
    settings = $7,
    is_enabled = $8,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteMonitor :exec
DELETE FROM monitors
WHERE id = $1;