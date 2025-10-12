CREATE TABLE users (
    id           BIGSERIAL                  NOT NULL                   PRIMARY KEY,
    email        VARCHAR(255)               NOT NULL                   UNIQUE,
    password     VARCHAR(255)               NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE   NOT NULL   DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE   NOT NULL   DEFAULT NOW()
);