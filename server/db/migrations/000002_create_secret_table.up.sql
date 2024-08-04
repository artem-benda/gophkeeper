CREATE TABLE secrets
(
    id bigserial NOT NULL,
    guid varchar(100) NOT NULL UNIQUE,
    name TEXT NOT NULL,
    enc_payload BYTEA,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    user_id bigint REFERENCES users,
    PRIMARY KEY (id)
);
