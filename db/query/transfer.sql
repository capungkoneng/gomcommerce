-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;


-- name: CreateTransfers :one
INSERT INTO transfers (
  from_akun, 
  to_akun,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;