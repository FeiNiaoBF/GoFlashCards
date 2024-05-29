BEGIN;
CREATE TABLE "cards" (
                         "id" bigserial PRIMARY KEY NOT NULL,
                         "front" text NOT NULL,
                         "back" text NOT NULL,
                         "know" bool NOT NULL DEFAULT FALSE,
                         "tags_id" integer NOT NULL,
                         "add_Time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
                        "id" bigserial  PRIMARY KEY NOT NULL,
                        "name" varchar(255) NOT NULL
);

COMMIT;