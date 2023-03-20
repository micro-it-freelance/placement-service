CREATE TYPE placement as ENUM (
    "offer",
    "order"
) 

CREATE TABLE palcements (
    id bigserial primary key,
    
    telegram_id bigint,
    username varchar(32),
    created_at timestamp DEFAULT NOW(),


    type placement,
    title text NOT NULL,
    description text NOT NULL,
    contacts text NOT NULL
);