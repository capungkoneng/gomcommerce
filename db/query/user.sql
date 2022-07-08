-- name: CreateUser :one
INSERT INTO users (
  username, 
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserJoin :many
SELECT DISTINCT t.*, a.owner
FROM (SELECT u.* FROM users u) as t
INNER JOIN akun a ON a.owner = t.username
ORDER BY t.username, t.full_name, t.hashed_password, t.email 
LIMIT $1
OFFSET $2;
