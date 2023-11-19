CREATE TABLE roles (
                       id VARCHAR(36)PRIMARY KEY,
                       name VARCHAR(50) UNIQUE NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
