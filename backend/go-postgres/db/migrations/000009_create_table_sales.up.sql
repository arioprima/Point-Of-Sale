CREATE TABLE sales (
                       id VARCHAR(36)PRIMARY KEY,
                       user_id VARCHAR(36) NOT NULL,
                       cart_id VARCHAR(36) NOT NULL,
                       store_id VARCHAR(36),
                       total_amount DECIMAL(10, 2) NOT NULL,
                       discount_amount DECIMAL(10, 2) NOT NULL,
                       payment_gateway_id VARCHAR(36) NOT NULL,
                       transaction_id VARCHAR(50),
                       payment_status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                       tax_id VARCHAR(36),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
