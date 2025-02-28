CREATE TABLE
    IF NOT EXISTS sessions (
        id SERIAL PRIMARY KEY,
        time_slot TIMESTAMP NOT NULL,
        team VARCHAR(100) NOT NULL,
        topic VARCHAR(255) NOT NULL DEFAULT '',
        presenter VARCHAR(100) NOT NULL DEFAULT ''
    );