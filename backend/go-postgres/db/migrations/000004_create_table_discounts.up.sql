CREATE TABLE discounts (
                           id VARCHAR(36)PRIMARY KEY,
                           percentage INT NOT NULL,
                           start_date DATE NOT NULL,
                           end_date DATE NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);