// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: kategori.sql

package db

import (
	"context"
	"database/sql"
)

const createKategori = `-- name: CreateKategori :one
INSERT INTO kategori (
  nama_kategori,
  deskripsi
) VALUES (
  $1, $2
) RETURNING id, nama_kategori, deskripsi, created_at
`

type CreateKategoriParams struct {
	NamaKategori string         `json:"nama_kategori"`
	Deskripsi    sql.NullString `json:"deskripsi"`
}

func (q *Queries) CreateKategori(ctx context.Context, arg CreateKategoriParams) (Kategori, error) {
	row := q.db.QueryRowContext(ctx, createKategori, arg.NamaKategori, arg.Deskripsi)
	var i Kategori
	err := row.Scan(
		&i.ID,
		&i.NamaKategori,
		&i.Deskripsi,
		&i.CreatedAt,
	)
	return i, err
}
