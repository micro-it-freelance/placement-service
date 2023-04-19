CREATE TYPE placement as ENUM ("OFFER", "ORDER") 

CREATE TABLE placement (
    id bigserial primary key,
    telegram_id bigint,
    username varchar(32),
    created_at timestamp DEFAULT NOW(),
    typ placement,
    title text NOT NULL,
    content text NOT NULL,
);