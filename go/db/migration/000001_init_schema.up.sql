CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "text" varchar NOT NULL,
  "is_done" boolean NOT NULL DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);