// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: mobil.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMobil = `-- name: CreateMobil :one
INSERT INTO mobil (
  nama, 
  deskripsi,
  kategori_id,
  user_id,
  gambar,
  trf_6jam,
  trf_12jam,
  trf_24jam,
  seat,
  top_speed,
  max_power,
  pintu,
  gigi
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING id, nama, deskripsi, kategori_id, gambar, user_id, trf_6jam, trf_12jam, trf_24jam, seat, top_speed, max_power, pintu, gigi, created_at
`

type CreateMobilParams struct {
	Nama       string         `json:"nama"`
	Deskripsi  sql.NullString `json:"deskripsi"`
	KategoriID int64          `json:"kategori_id"`
	UserID     string         `json:"user_id"`
	Gambar     sql.NullString `json:"gambar"`
	Trf6jam    int64          `json:"trf_6jam"`
	Trf12jam   int64          `json:"trf_12jam"`
	Trf24jam   int64          `json:"trf_24jam"`
	Seat       sql.NullInt64  `json:"seat"`
	TopSpeed   sql.NullInt64  `json:"top_speed"`
	MaxPower   sql.NullInt64  `json:"max_power"`
	Pintu      sql.NullInt64  `json:"pintu"`
	Gigi       sql.NullString `json:"gigi"`
}

func (q *Queries) CreateMobil(ctx context.Context, arg CreateMobilParams) (Mobil, error) {
	row := q.db.QueryRowContext(ctx, createMobil,
		arg.Nama,
		arg.Deskripsi,
		arg.KategoriID,
		arg.UserID,
		arg.Gambar,
		arg.Trf6jam,
		arg.Trf12jam,
		arg.Trf24jam,
		arg.Seat,
		arg.TopSpeed,
		arg.MaxPower,
		arg.Pintu,
		arg.Gigi,
	)
	var i Mobil
	err := row.Scan(
		&i.ID,
		&i.Nama,
		&i.Deskripsi,
		&i.KategoriID,
		&i.Gambar,
		&i.UserID,
		&i.Trf6jam,
		&i.Trf12jam,
		&i.Trf24jam,
		&i.Seat,
		&i.TopSpeed,
		&i.MaxPower,
		&i.Pintu,
		&i.Gigi,
		&i.CreatedAt,
	)
	return i, err
}

const getMobilJoinMany = `-- name: GetMobilJoinMany :many
SELECT DISTINCT t.nama, a.nama_kategori, a.id, t.kategori_id
FROM (SELECT m.id, m.nama, m.deskripsi, m.kategori_id, m.gambar, m.user_id, m.trf_6jam, m.trf_12jam, m.trf_24jam, m.seat, m.top_speed, m.max_power, m.pintu, m.gigi, m.created_at FROM mobil m) as t
INNER JOIN kategori a ON a.id = t.kategori_id
INNER JOIN (SELECT o.username from users o)
as t2 ON t2.username = t.user_id
LIMIT $1
OFFSET $2
`

type GetMobilJoinManyParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetMobilJoinManyRow struct {
	Nama         string `json:"nama"`
	NamaKategori string `json:"nama_kategori"`
	ID           int64  `json:"id"`
	KategoriID   int64  `json:"kategori_id"`
}

func (q *Queries) GetMobilJoinMany(ctx context.Context, arg GetMobilJoinManyParams) ([]GetMobilJoinManyRow, error) {
	rows, err := q.db.QueryContext(ctx, getMobilJoinMany, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMobilJoinManyRow{}
	for rows.Next() {
		var i GetMobilJoinManyRow
		if err := rows.Scan(
			&i.Nama,
			&i.NamaKategori,
			&i.ID,
			&i.KategoriID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMobilJoinOne = `-- name: GetMobilJoinOne :one
SELECT t1.id, t1.nama, t1.deskripsi, t1.kategori_id, t1.gambar, t1.user_id, t1.trf_6jam, t1.trf_12jam, t1.trf_24jam, t1.seat, t1.top_speed, t1.max_power, t1.pintu, t1.gigi, t1.created_at, t2.username, t.nama_kategori
FROM mobil t1
JOIN users t2 ON t2.username = t1.user_id
JOIN kategori t ON t.id = t1.kategori_id
WHERE t1.id = $1
`

type GetMobilJoinOneRow struct {
	ID           int64          `json:"id"`
	Nama         string         `json:"nama"`
	Deskripsi    sql.NullString `json:"deskripsi"`
	KategoriID   int64          `json:"kategori_id"`
	Gambar       sql.NullString `json:"gambar"`
	UserID       string         `json:"user_id"`
	Trf6jam      int64          `json:"trf_6jam"`
	Trf12jam     int64          `json:"trf_12jam"`
	Trf24jam     int64          `json:"trf_24jam"`
	Seat         sql.NullInt64  `json:"seat"`
	TopSpeed     sql.NullInt64  `json:"top_speed"`
	MaxPower     sql.NullInt64  `json:"max_power"`
	Pintu        sql.NullInt64  `json:"pintu"`
	Gigi         sql.NullString `json:"gigi"`
	CreatedAt    time.Time      `json:"created_at"`
	Username     string         `json:"username"`
	NamaKategori string         `json:"nama_kategori"`
}

func (q *Queries) GetMobilJoinOne(ctx context.Context, id int64) (GetMobilJoinOneRow, error) {
	row := q.db.QueryRowContext(ctx, getMobilJoinOne, id)
	var i GetMobilJoinOneRow
	err := row.Scan(
		&i.ID,
		&i.Nama,
		&i.Deskripsi,
		&i.KategoriID,
		&i.Gambar,
		&i.UserID,
		&i.Trf6jam,
		&i.Trf12jam,
		&i.Trf24jam,
		&i.Seat,
		&i.TopSpeed,
		&i.MaxPower,
		&i.Pintu,
		&i.Gigi,
		&i.CreatedAt,
		&i.Username,
		&i.NamaKategori,
	)
	return i, err
}
