CREATE TABLE tasks (
    title VARCHAR(200) PRIMARY KEY,
    description VARCHAR(200),
    completed BOOLEAN NOT NULL,
    create_at DATE NOT NULL,
    completed_at DATE
)