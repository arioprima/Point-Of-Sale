CREATE TABLE product_discounts (
                                   id VARCHAR(36)PRIMARY KEY,
                                   product_id VARCHAR(36) NOT NULL,
                                   discount_id VARCHAR(36) NOT NULL,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);