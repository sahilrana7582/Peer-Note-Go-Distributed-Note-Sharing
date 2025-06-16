CREATE TABLE IF NOT EXISTS peers (
    id SERIAL PRIMARY KEY,
    ip TEXT NOT NULL,
    port INT NOT NULL
);


CREATE TABLE IF NOT EXISTS files (
    id SERIAL PRIMARY KEY,
    file_name TEXT NOT NULL,
    course_code TEXT NOT NULL,
    professor VARCHAR(255) NOT NULL,
    peer_id INT REFERENCES peers(id),
    keywords TEXT[],
    file_path TEXT NOT NULL
);