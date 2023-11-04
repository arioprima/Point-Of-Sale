create table users (
    user_id varchar(36) primary key,
    firstname varchar(255) not null,
    lastname varchar(255),
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    role enum('admin', 'staff', 'employee') default 'employee',
    image varchar(255),
    is_deleted tinyint(1) default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;