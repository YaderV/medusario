CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR ( 150 ) NOT NULL,
    password VARCHAR ( 1024 ) NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
);

CREATE INDEX users_id_idx ON users (id);

CREATE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR ( 150 ) NOT NULL,
    start_date TIMESTAMPTZ,
    is_active BOOLEAN DEFAULT TRUE,
    user_id REFERENCES users (id) ON DELETE CASCADE NOT NULL,
);

CREATE INDEX courses_id_idx ON courses (id);
CREATE INDEX courses_user_id_idx ON courses (user_id);

CREATE lessons (
    id SERIAL PRIMARY KEY,
    topic VARCHAR ( 150 ) NOT NULL,
    num INT NOT NULL DEFAULT 0,
    text TEXT NOT NULL,
    course_id REFERENCES courses (id) ON DELETE
);

CREATE INDEX lessons_id_idx ON lessons (id);
CREATE INDEX lessons_course_id_idx ON lessons (course_id);
