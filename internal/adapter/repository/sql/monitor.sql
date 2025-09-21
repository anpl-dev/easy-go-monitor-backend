-- name: CreateMonitor :one
INSERT INTO monitors (id, user_id, name, url, interval_seconds)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMonitorByID :one
SELECT * FROM monitors
WHERE id = $1;

-- name: ListMonitorsByUser :many
SELECT * FROM monitors
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateMonitor :one
UPDATE monitors
SET name = $2,
    url = $3,
    interval_seconds = $4,
    updated_at = now()
WHERE id = $1
RETURNING *;;

-- name: DeleteMonitor :exec
DELETE FROM monitors
WHERE id = $1;