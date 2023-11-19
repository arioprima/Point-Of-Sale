CREATE TABLE orders (
                        id VARCHAR(36)PRIMARY KEY,
                        user_id VARCHAR(36) NOT NULL,
                        sale_id VARCHAR(36) NOT NULL,
                        shipping_method_id VARCHAR(36) NOT NULL,
                        address VARCHAR(255) NOT NULL,
                        status VARCHAR(50) NOT NULL,
                        customers_member_id VARCHAR(36) NOT NULL,
                        customers_non_member_id VARCHAR(36) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
