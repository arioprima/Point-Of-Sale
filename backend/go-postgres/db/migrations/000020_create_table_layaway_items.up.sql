CREATE TABLE layaway_items (
                               id VARCHAR(36)PRIMARY KEY,
                               sale_id VARCHAR(36) NOT NULL,
                               product_id VARCHAR(36) NOT NULL,
                               quantity INT NOT NULL,
                               reserved_until DATE NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
