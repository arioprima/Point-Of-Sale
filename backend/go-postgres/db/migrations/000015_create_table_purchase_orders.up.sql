CREATE TABLE purchase_orders (
                                 id VARCHAR(36)PRIMARY KEY,
                                 supplier_id VARCHAR(36) NOT NULL,
                                 order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);