-- name: CreateUser :one
INSERT INTO users (id, username, email, first_name, last_name, password_hash, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserWithUsername :one
SELECT * FROM users WHERE username = $1;
