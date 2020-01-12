CREATE TABLE temperature
(
    id         VARCHAR PRIMARY KEY,
    city_id       VARCHAR NOT NULL,
    max int NOT NULL,
    min int NOT NULL,
    timestamp int NOT NULL
);
