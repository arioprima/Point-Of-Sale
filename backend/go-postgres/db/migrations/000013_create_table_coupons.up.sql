CREATE TABLE coupons (
                         id VARCHAR(36)PRIMARY KEY,
                         code VARCHAR(20) UNIQUE NOT NULL,
                         discount_percentage INT NOT NULL,
                         start_date DATE NOT NULL,
                         end_date DATE NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);