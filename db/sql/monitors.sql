-- name: CreateMonitor :one
INSERT INTO monitors (
    id, user_id, name, url, type, settings, is_enabled
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
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
    name = $2,
    url = $3,
    type = $4,
    settings = $5,
    is_enabled = $6,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteMonitor :exec
DELETE FROM monitors
WHERE id = $1;

-- name: UpdateMonitorIsEnabled :one
UPDATE monitors
SET 
  is_enabled = $2,
  updated_at = now()
WHERE id = $1
RETURNING *;