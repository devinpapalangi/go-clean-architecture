-- Create "users" table
CREATE TABLE "public"."users" (
  "id" character varying(36) NOT NULL,
  "username" character varying(100) NOT NULL,
  "email" character varying(100) NOT NULL,
  "name" character varying(100) NOT NULL,
  "password" character varying(100) NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
