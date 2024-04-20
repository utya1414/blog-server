CREATE TABLE IF NOT EXISTS "users" (
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hasshed_password" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar UNIQUE NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "sessions" ("username");

CREATE INDEX ON "sessions" ("refresh_token");

CREATE INDEX ON "sessions" ("expires_at");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
