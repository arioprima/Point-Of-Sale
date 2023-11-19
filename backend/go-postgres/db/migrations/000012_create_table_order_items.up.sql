CREATE TABLE order_items (
                             id VARCHAR(36)PRIMARY KEY,
                             sale_id VARCHAR(36) NOT NULL,
                             product_id VARCHAR(36) NOT NULL,
                             quantity INT NOT NULL,
                             unit_price DECIMAL(10, 2) NOT NULL,
                             total_amount DECIMAL(10, 2) NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);