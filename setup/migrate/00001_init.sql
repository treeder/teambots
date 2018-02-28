-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
