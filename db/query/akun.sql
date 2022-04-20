-- name: CreateAuthor :one
INSERT INTO akun (
  owner, 
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM akun
WHERE id = $1 LIMIT 1;

-- name: GetAkunForUpdate :one
SELECT * FROM akun
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAuthors :many
SELECT * FROM akun
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: UpdateAkun :one
UPDATE akun 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAkunBalance :one
UPDATE akun 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAkun :exec
DELETE FROM akun 
WHERE id = $1;