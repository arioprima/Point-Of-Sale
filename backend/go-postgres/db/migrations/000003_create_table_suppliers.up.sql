CREATE TABLE suppliers (
                           id VARCHAR(36)PRIMARY KEY,
                           name VARCHAR(50) NOT NULL,
                           contact_person VARCHAR(50),
                           contact_email VARCHAR(50),
                           contact_phone VARCHAR(20),
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
