CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "role" AS ENUM (
  'Registration',
  'Support'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "is_admin" boolean DEFAULT 'FALSE'
);

CREATE TABLE "voters" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "registered_at" timestamptz PRIMARY KEY NOT NULL DEFAULT 'now()',
  "voted_at" timestamptz,
  "voted" boolean,
  "voters_public_address" varchar NOT NULL
);

CREATE TABLE "admins" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" role NOT NULL
);

CREATE TABLE "contestants" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "position" varchar NOT NULL,
  "registered_at" timestamptz DEFAULT 'now()',
  "description" varchar NOT NULL,
  "region" varchar NOT NULL,
  "ethereum_address" varchar NOT NULL,
  "national_id_number" bigint NOT NULL,
  "image_address" varchar NOT NULL
);

CREATE INDEX ON "admins" ("full_name");

CREATE INDEX ON "contestants" ("full_name");

ALTER TABLE "voters" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");

ALTER TABLE "admins" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");

ALTER TABLE "contestants" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");
