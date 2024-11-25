-- name: CreateUser :one
INSERT INTO users (phone_number, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE phone_number = $1;