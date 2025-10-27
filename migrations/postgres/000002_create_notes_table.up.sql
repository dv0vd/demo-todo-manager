CREATE TABLE notes (
    id            BIGSERIAL                  NOT NULL                   PRIMARY KEY,
    title         VARCHAR(255)               NOT NULL,
    description   TEXT,
    user_id       BIGSERIAL                  NOT NULL                   REFERENCES users(id)   ON UPDATE RESTRICT   ON DELETE RESTRICT,
    done          BOOLEAN                    NOT NULL   DEFAULT FALSE,
    created_at    TIMESTAMP WITH TIME ZONE   NOT NULL   DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE   NOT NULL   DEFAULT NOW()
);