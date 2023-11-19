CREATE TABLE products (
                          id VARCHAR(36)PRIMARY KEY,
                          name VARCHAR(50) NOT NULL,
                          price DECIMAL(10, 2) NOT NULL,
                          stock INT NOT NULL,
                          category VARCHAR(50) NOT NULL,
                          description TEXT,
                          image_url VARCHAR(255),
                          supplier_id VARCHAR(36),
                          store_id VARCHAR(36),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);