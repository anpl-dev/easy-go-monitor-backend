-- name: CreateUser :one
INSERT INTO users (id, name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: FindAllUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    email = $3,
    password = $4,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;