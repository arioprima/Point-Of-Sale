CREATE TABLE users (
                       id VARCHAR(36)PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       email VARCHAR(50) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role_id VARCHAR(36),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);