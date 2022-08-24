
-- +migrate Up
CREATE TABLE "persons" (
  "id" BIGSERIAL PRIMARY KEY,
  "createdAt" TIMESTAMP,
  "updatedAt" TIMESTAMP,
  "firstname" VARCHAR(50) NOT NULL,
  "lastname" VARCHAR(50) NOT NULL,
  "birthdate" DATE, 
  "emailaddress" VARCHAR(100) UNIQUE
);

CREATE TABLE "stamps" (
  "id" BIGSERIAL PRIMARY KEY,
  "createdAt" TIMESTAMP,
  "updatedAt" TIMESTAMP,
  "person_id" BIGINT,
  "checkin" BOOLEAN NOT NULL DEFAULT FALSE,
  "stamp" TIMESTAMP NOT NULL,
  FOREIGN KEY ("person_id") REFERENCES "persons" ("id")
);

-- ALTER TABLE "stamps" ADD FOREIGN KEY ("person_id") REFERENCES "persons" ("id");

CREATE INDEX ON "persons" ("emailaddress");

COMMENT ON TABLE "stamps" IS 'Timestamps for persons';


INSERT INTO persons(firstname, lastname, birthdate, emailaddress) 
  VALUES ('Sigi', 'Reinhardt', '1977-01-01', 'sigi@reinhardt.de');

INSERT INTO persons(firstname, lastname, birthdate, emailaddress) 
  VALUES ('Michi', 'Sofia', '1971-05-07', 'michzuerch@gmail.com');

INSERT INTO stamps(person_id, checkin, stamp) 
  VALUES (1, TRUE, '1997-01-01 10:58:19');

-- +migrate Down
DROP TABLE "stamps";
DROP TABLE "persons";
