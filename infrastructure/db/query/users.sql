-- name: GetAuthor :one
SELECT * FROM users
WHERE username = $1 LIMIT username;

-- name: ListAuthors :many
SELECT * FROM users;

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