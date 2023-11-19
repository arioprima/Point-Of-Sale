CREATE TABLE stores (
                        id VARCHAR(36)PRIMARY KEY,
                        name VARCHAR(50) NOT NULL,
                        location VARCHAR(255),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);