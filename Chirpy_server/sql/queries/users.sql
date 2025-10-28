-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1
)
RETURNING *;

-- name: StorePassword :exec
INSERT INTO users (hashed_password)
VALUES (
  $1
);

-- name: GetUserbyEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: DeleteUser :exec
DELETE FROM users;