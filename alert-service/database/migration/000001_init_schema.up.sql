CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "Alerts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "crypro" varchar NOT NULL,
  "price" float NOT NULL,
  "direction" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "Alerts" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");