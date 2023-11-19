CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);



CREATE TABLE roles (
                       id VARCHAR(36)PRIMARY KEY,
                       name VARCHAR(50) UNIQUE NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE suppliers (
                           id VARCHAR(36)PRIMARY KEY,
                           name VARCHAR(50) NOT NULL,
                           contact_person VARCHAR(50),
                           contact_email VARCHAR(50),
                           contact_phone VARCHAR(20),
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE discounts (
                           id VARCHAR(36)PRIMARY KEY,
                           percentage INT NOT NULL,
                           start_date DATE NOT NULL,
                           end_date DATE NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_discounts (
                                   id VARCHAR(36)PRIMARY KEY,
                                   product_id VARCHAR(36) NOT NULL,
                                   discount_id VARCHAR(36) NOT NULL,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   FOREIGN KEY (product_id) REFERENCES products(id),
                                   FOREIGN KEY (discount_id) REFERENCES discounts(id)
);


CREATE TABLE payment_gateways (
                                  id VARCHAR(36)PRIMARY KEY,
                                  name VARCHAR(50) NOT NULL,
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE carts (
                       id VARCHAR(36)PRIMARY KEY,
                       user_id VARCHAR(36) NOT NULL,
                       store_id VARCHAR(36),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (user_id) REFERENCES users(id),
                       FOREIGN KEY (store_id) REFERENCES stores(id)
);


CREATE TABLE cart_items (
                            id VARCHAR(36)PRIMARY KEY,
                            cart_id VARCHAR(36) NOT NULL,
                            product_id VARCHAR(36) NOT NULL,
                            quantity INT NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            FOREIGN KEY (cart_id) REFERENCES carts(id),
                            FOREIGN KEY (product_id) REFERENCES products(id)
);


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
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (user_id) REFERENCES users(id),
                        FOREIGN KEY (cart_id) REFERENCES carts(id),
                        FOREIGN KEY (store_id) REFERENCES stores(id),
                        FOREIGN KEY (payment_gateway_id) REFERENCES payment_gateways(id),
                        FOREIGN KEY (transaction_id) REFERENCES transaction_history(id),
                        FOREIGN KEY (tax_id) REFERENCES taxes(id)
);


CREATE TABLE shipping_methods (
                                  id VARCHAR(36)PRIMARY KEY,
                                  name VARCHAR(50) NOT NULL,
                                  cost DECIMAL(10, 2) NOT NULL,
                                  estimated_delivery_time INT NOT NULL, -- in days
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


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
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (user_id) REFERENCES users(id),
                        FOREIGN KEY (sale_id) REFERENCES sales(id),
                        FOREIGN KEY (shipping_method_id) REFERENCES shipping_methods(id),
                        FOREIGN KEY (customers_member_id) REFERENCES customers_members(id),
                        FOREIGN KEY (customers_non_member_id) REFERENCES customers_non_members(id)
);


CREATE TABLE order_items (
                             id VARCHAR(36)PRIMARY KEY,
                             sale_id VARCHAR(36) NOT NULL,
                             product_id VARCHAR(36) NOT NULL,
                             quantity INT NOT NULL,
                             unit_price DECIMAL(10, 2) NOT NULL,
                             total_amount DECIMAL(10, 2) NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (sale_id) REFERENCES sales(id),
                             FOREIGN KEY (product_id) REFERENCES products(id)
);


CREATE TABLE coupons (
                         id VARCHAR(36)PRIMARY KEY,
                         code VARCHAR(20) UNIQUE NOT NULL,
                         discount_percentage INT NOT NULL,
                         start_date DATE NOT NULL,
                         end_date DATE NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE transaction_history (
                                     id VARCHAR(36)PRIMARY KEY,
                                     user_id VARCHAR(36) NOT NULL,
                                     sale_id VARCHAR(36) NOT NULL,
                                     payment_amount DECIMAL(10, 2) NOT NULL,
                                     payment_method VARCHAR(50) NOT NULL,
                                     transaction_status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     FOREIGN KEY (user_id) REFERENCES users(id),
                                     FOREIGN KEY (sale_id) REFERENCES sales(id)
);


CREATE TABLE purchase_orders (
                                 id VARCHAR(36)PRIMARY KEY,
                                 supplier_id VARCHAR(36) NOT NULL,
                                 order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 FOREIGN KEY (supplier_id) REFERENCES suppliers(id)
);


CREATE TABLE received_items (
                                id VARCHAR(36)PRIMARY KEY,
                                purchase_order_id VARCHAR(36) NOT NULL,
                                product_id VARCHAR(36) NOT NULL,
                                store_id VARCHAR(36) NOT NULL,
                                quantity INT NOT NULL,
                                received_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id),
                                FOREIGN KEY (product_id) REFERENCES products(id),
                                FOREIGN KEY (store_id) REFERENCES stores(id)
);


CREATE TABLE taxes (
                       id VARCHAR(36)PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       rate DECIMAL(5, 2) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE gift_cards (
                            id VARCHAR(36)PRIMARY KEY,
                            code VARCHAR(20) UNIQUE NOT NULL,
                            balance DECIMAL(10, 2) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE gift_card_transactions (
                                        id VARCHAR(36)PRIMARY KEY,
                                        sale_id VARCHAR(36) NOT NULL,
                                        gift_card_id VARCHAR(36) NOT NULL,
                                        amount DECIMAL(10, 2) NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        FOREIGN KEY (sale_id) REFERENCES sales(id),
                                        FOREIGN KEY (gift_card_id) REFERENCES gift_cards(id)
);


CREATE TABLE layaway_items (
                               id VARCHAR(36)PRIMARY KEY,
                               sale_id VARCHAR(36) NOT NULL,
                               product_id VARCHAR(36) NOT NULL,
                               quantity INT NOT NULL,
                               reserved_until DATE NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               FOREIGN KEY (sale_id) REFERENCES sales(id),
                               FOREIGN KEY (product_id) REFERENCES products(id)
);


CREATE TABLE consignment_items (
                                   id VARCHAR(36)PRIMARY KEY,
                                   product_id VARCHAR(36) NOT NULL,
                                   quantity INT NOT NULL,
                                   consignor_name VARCHAR(50) NOT NULL,
                                   status VARCHAR(50) NOT NULL DEFAULT 'Pending',
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   FOREIGN KEY (product_id) REFERENCES products(id)
);


CREATE TABLE stores (
                        id VARCHAR(36)PRIMARY KEY,
                        name VARCHAR(50) NOT NULL,
                        location VARCHAR(255),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


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
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (supplier_id) REFERENCES suppliers(id),   
                          FOREIGN KEY (store_id) REFERENCES stores(id)
);

CREATE TABLE customers_members (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    membership_level VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE customers_non_members (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



