CREATE TABLE webhook
(
    id         VARCHAR PRIMARY KEY,
    city_id       VARCHAR NOT NULL,
    callback_url  VARCHAR NOT NULL
);

