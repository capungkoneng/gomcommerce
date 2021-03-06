CREATE TABLE "akun" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "akun_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_akun" bigint NOT NULL,
  "to_akun" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("akun_id") REFERENCES "akun" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_akun") REFERENCES "akun" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_akun") REFERENCES "akun" ("id");

CREATE INDEX ON "akun" ("owner");

CREATE INDEX ON "entries" ("akun_id");

CREATE INDEX ON "transfers" ("from_akun");

CREATE INDEX ON "transfers" ("to_akun");

CREATE INDEX ON "transfers" ("from_akun", "to_akun");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
