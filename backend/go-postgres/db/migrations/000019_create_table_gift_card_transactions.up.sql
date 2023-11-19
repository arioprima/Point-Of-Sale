CREATE TABLE gift_card_transactions (
                                        id VARCHAR(36)PRIMARY KEY,
                                        sale_id VARCHAR(36) NOT NULL,
                                        gift_card_id VARCHAR(36) NOT NULL,
                                        amount DECIMAL(10, 2) NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);