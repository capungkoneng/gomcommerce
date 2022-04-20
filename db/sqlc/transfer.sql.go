// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfers = `-- name: CreateTransfers :one
INSERT INTO transfers (
  from_akun, 
  to_akun,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_akun, to_akun, amount, created_at
`

type CreateTransfersParams struct {
	FromAkun int64 `json:"from_akun"`
	ToAkun   int64 `json:"to_akun"`
	Amount   int64 `json:"amount"`
}

func (q *Queries) CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfers, arg.FromAkun, arg.ToAkun, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAkun,
		&i.ToAkun,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_akun, to_akun, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAkun,
		&i.ToAkun,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
