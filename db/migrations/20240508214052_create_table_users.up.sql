CREATE TABLE users (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    token VARCHAR(100),
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);
