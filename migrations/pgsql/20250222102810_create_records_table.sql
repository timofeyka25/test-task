-- +goose Up
CREATE TABLE IF NOT EXISTS records
(
    id         SERIAL PRIMARY KEY,
    data       TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

INSERT INTO records (data)
VALUES ('Record 1'),
       ('Record 2'),
       ('Record 3'),
       ('Record 4'),
       ('Record 5'),
       ('Record 6'),
       ('Record 7'),
       ('Record 8'),
       ('Record 9'),
       ('Record 10'),
       ('Record 11'),
       ('Record 12'),
       ('Record 13'),
       ('Record 14'),
       ('Record 15'),
       ('Record 16'),
       ('Record 17'),
       ('Record 18'),
       ('Record 19'),
       ('Record 20');

-- +goose Down
DROP TABLE IF EXISTS records;
