CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    teams_id INT,
    name TEXT NOT NULL,
    jersey_number TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
