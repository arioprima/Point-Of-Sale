create table product(
    product_id varchar(36) PRIMARY KEY ,
    product_name varchar(100) not null ,
    category_id varchar(36) not null ,
    price int(10) not null,
    description varchar(255),
    quantity int(10) not null,
    condition varchar(255) not null ,
    image varchar(255),
    supplier_id varchar(36),
    date_of_arrival datetime,
    expiry_date date,
    is_deleted tinyint(1) default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)