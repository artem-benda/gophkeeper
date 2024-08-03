CREATE TABLE secrets
(
    id bigserial NOT NULL,
    guid varchar(100) NOT NULL UNIQUE,
    name TEXT NOT NULL,
    enc_payload BYTEA,
    created_at datetime with timezone,
    updated_at datetime with timezone,
    user_id bigint REFERENCES users,
    PRIMARY KEY (id)
);
