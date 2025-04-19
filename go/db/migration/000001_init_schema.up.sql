CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);