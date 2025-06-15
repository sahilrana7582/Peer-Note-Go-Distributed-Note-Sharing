CREATE TABLE IF NOT EXISTS peers (
    id SERIAL PRIMARY KEY,
    ip TEXT NOT NULL,
    port INT NOT NULL
);


CREATE TABLE IF NOT EXISTS files (
    id SERIAL PRIMARY KEY,
    file_name TEXT NOT NULL,
    peer_id INT REFERENCES peers(id),
    keywords TEXT[]
);