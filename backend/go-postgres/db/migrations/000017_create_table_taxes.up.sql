CREATE TABLE taxes (
                       id VARCHAR(36)PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       rate DECIMAL(5, 2) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
