CREATE TABLE "user_role" (
    "user_id" int,
    "role_id" int
);

ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");