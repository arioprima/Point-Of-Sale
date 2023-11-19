CREATE TABLE consignment_items (
                                   id VARCHAR(36)PRIMARY KEY,
                                   product_id VARCHAR(36) NOT NULL,
                                   quantity INT NOT NULL,
                                   consignor_name VARCHAR(50) NOT NULL,
                                   status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);