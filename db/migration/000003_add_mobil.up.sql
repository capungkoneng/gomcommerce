CREATE TABLE "kategori" (
  "id" BIGSERIAL PRIMARY KEY,
  "nama_kategori" varchar NOT NULL,
  "deskripsi" text ,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "mobil" (
  "id" BIGSERIAL PRIMARY KEY,
  "nama" varchar NOT NULL,
  "deskripsi" text ,
  "kategori_id" bigint NOT NULL,
  "gambar" varchar ,
  "user_id" varchar NOT NULL,
  "trf_6jam" bigint NOT NULL,
  "trf_12jam" bigint NOT NULL,
  "trf_24jam" bigint NOT NULL,
  "seat" bigint  ,
  "top_speed" bigint  ,
  "max_power" bigint  ,
  "pintu" bigint  ,
  "gigi" varchar  ,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "mobil" ADD FOREIGN KEY ("kategori_id") REFERENCES "kategori" ("id");

ALTER TABLE "mobil" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

