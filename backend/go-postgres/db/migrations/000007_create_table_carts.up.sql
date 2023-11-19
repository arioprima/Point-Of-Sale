CREATE TABLE carts (
                       id VARCHAR(36)PRIMARY KEY,
                       user_id VARCHAR(36) NOT NULL,
                       store_id VARCHAR(36),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);