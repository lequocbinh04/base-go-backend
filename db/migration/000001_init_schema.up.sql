CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "email" varchar NOT NULL,
                         "password" varchar,
                         "crondata_token" json,
                         "balance" bigint NOT NULL DEFAULT 0,
                         "balance_profile" bigint NOT NULL DEFAULT 0,
                         "role" varchar NOT NULL,
                         "status" varchar,
                         "created_at" timestamp DEFAULT (now()),
                         "updated_at" timestamp
);

CREATE TABLE "profile" (
                           "id" bigserial PRIMARY KEY,
                           "user_id" bigserial,
                           "name" varchar NOT NULL,
                           "information" json,
                           "expired_at" timestamp,
                           "status" varchar,
                           "created_at" timestamp DEFAULT (now()),
                           "updated_at" timestamp
);

ALTER TABLE "profile" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");