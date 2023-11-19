CREATE TABLE received_items (
                                id VARCHAR(36)PRIMARY KEY,
                                purchase_order_id VARCHAR(36) NOT NULL,
                                product_id VARCHAR(36) NOT NULL,
                                store_id VARCHAR(36) NOT NULL,
                                quantity INT NOT NULL,
                                received_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);