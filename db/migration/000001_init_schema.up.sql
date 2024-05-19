BEGIN;
CREATE TABLE "cards" (
                         "id" bigserial PRIMARY KEY NOT NULL,
                         "front" text NOT NULL,
                         "back" text NOT NULL,
                         "know" bool NOT NULL DEFAULT FALSE,
                         "add_Time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
                        "id" integer PRIMARY KEY NOT NULL,
                        "name" text NOT NULL
);

CREATE TABLE "card_tags" (
                             "cards_id" bigserial NOT NULL,
                             "tags_id" integer NOT NULL,
                             PRIMARY KEY ("cards_id", "tags_id")
);

ALTER TABLE "card_tags" ADD FOREIGN KEY ("cards_id") REFERENCES "cards" ("id");
ALTER TABLE "card_tags" ADD FOREIGN KEY ("tags_id") REFERENCES "tags" ("id");

COMMIT;