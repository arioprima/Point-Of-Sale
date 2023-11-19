CREATE TABLE shipping_methods (
                                  id VARCHAR(36)PRIMARY KEY,
                                  name VARCHAR(50) NOT NULL,
                                  cost DECIMAL(10, 2) NOT NULL,
                                  estimated_delivery_time INT NOT NULL, -- in days
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);