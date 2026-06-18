CREATE TABLE tasks (
    title VARCHAR(200) PRIMARY KEY,
    description VARCHAR(200),
    completed BOOLEAN NOT NULL,
    created_at DATE NOT NULL,
    completed_at DATE
)