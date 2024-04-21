-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY username LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  username, email, hasshed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- -- name: UpdateAuthor :exec
-- UPDATE authors
--   set name = $2,
--   bio = $3
-- WHERE id = $1;

-- -- name: DeleteAuthor :exec
-- DELETE FROM authors
-- WHERE id = $1;