CREATE TYPE placement as ENUM ("OFFER", "ORDER") 

CREATE TABLE placements (
    id bigserial primary key,
    telegram_id bigint,
    typ placement NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    created_at timestamp DEFAULT NOW(),
);