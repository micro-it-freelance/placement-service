CREATE TYPE placement as ENUM (
    "offer",
    "order"
) 

CREATE TABLE palcements (
    id bigserial primary key,
    
    telegram_id bigint,
    username varchar(32),
    created_at timestamp DEFAULT NOW(),

    typ placement,
    title text NOT NULL,
    content text NOT NULL,
);