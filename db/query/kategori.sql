-- name: CreateKategori :one
INSERT INTO kategori (
  nama_kategori,
  deskripsi
) VALUES (
  $1, $2
) RETURNING *;