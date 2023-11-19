CREATE TABLE gift_cards (
                            id VARCHAR(36)PRIMARY KEY,
                            code VARCHAR(20) UNIQUE NOT NULL,
                            balance DECIMAL(10, 2) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
