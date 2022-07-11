-- name: CreateMobil :one
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
) RETURNING *;


-- name: GetMobilJoinMany :many
SELECT DISTINCT t.*, t.kategori_id
FROM (SELECT m.* FROM mobil m) as t
INNER JOIN kategori a ON a.id = t.kategori_id
INNER JOIN (SELECT o.username from users o)
as t2 ON t2.username = t.user_id
LIMIT $1
OFFSET $2;

-- name: GetMobilJoinOne :one
SELECT t1.*, t2.username, t.nama_kategori
FROM mobil t1
JOIN users t2 ON t2.username = t1.user_id
JOIN kategori t ON t.id = t1.kategori_id
WHERE t1.id = $1;

