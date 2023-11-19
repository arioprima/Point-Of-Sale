CREATE TABLE transaction_history (
                                     id VARCHAR(36)PRIMARY KEY,
                                     user_id VARCHAR(36) NOT NULL,
                                     sale_id VARCHAR(36) NOT NULL,
                                     payment_amount DECIMAL(10, 2) NOT NULL,
                                     payment_method VARCHAR(50) NOT NULL,
                                     transaction_status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);