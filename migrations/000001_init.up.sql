CREATE locations (
    id SERIAL PRIMARY KEY,
    title VARCHAR (50) NOT NULL,
    description TEXT NOT NULL,
    address VARCHAR(140) NOT NULL,
);

CREATE INDEX locations_id_idx ON locations (id);
