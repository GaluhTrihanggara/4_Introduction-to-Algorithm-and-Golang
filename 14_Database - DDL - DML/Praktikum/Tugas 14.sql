-- 1 Create database alta_online_shop.
CREATE DATABASE alta_onlineshop;
USE alta_online_shop;
-- 2a membuat tabel user
CREATE TABLE users (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (100) NOT NULL,
addres VARCHAR (200),
date_of_birth DATE,
status_user VARCHAR (30),
gender ENUM('Male','Female','Other'),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- 2b membuat tabel product
CREATE TABLE product (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
code VARCHAR (50),
name VARCHAR (100),
status SMALLINT (5),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2b membuat tabel product types
CREATE TABLE product_types (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (100),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2b membuat tabel operators
CREATE TABLE operators (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (100),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- membuat foreign keynya 
ALTER TABLE product
ADD COLUMN product_types_id INT(11),
ADD COLUMN operator_id INT(11),
ADD CONSTRAINT FK_product_product_types
  FOREIGN KEY (product_types_id)
  REFERENCES product_types(id),
ADD CONSTRAINT FK_product_operator
  FOREIGN KEY (operator_id)
  REFERENCES operators(id);
  
-- 2b membuat tabel product description
CREATE TABLE product_descriptions (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
description TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- 2b membuat tabel paymen method
CREATE TABLE payment_methods (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name varchar (100) not null,
status SMALLINT (5),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2c membuat tabel transaction
CREATE TABLE transactions (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
status VARCHAR (25),
total_qty INT (15) NOT NULL,
total_pricw NUMERIC (25,2),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2c membuat tabel transaction details
CREATE TABLE transaction_details (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  status VARCHAR(20),
  qty INT(15),
  price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- membuat foreignkey agar terhubung dengan tabel product
ALTER TABLE transaction_details
ADD COLUMN product_id INT(11),
ADD CONSTRAINT FK_transaction_details_product
  FOREIGN KEY (product_id)
  REFERENCES product(id);
  
-- membuat foreign key nya didalam tabel transaction
ALTER TABLE transactions
ADD COLUMN payment_methods_id INT(11),
ADD COLUMN user_id INT(11),
ADD COLUMN transaction_details_id INT(11),
ADD CONSTRAINT FK_transactions_payment_methods
FOREIGN KEY (payment_methods_id)
REFERENCES payment_methods(id),
ADD CONSTRAINT FK_transactions_users
FOREIGN KEY (user_id)
REFERENCES users(id),
ADD CONSTRAINT FK_transactions_transaction_details
FOREIGN KEY (transaction_details_id)
REFERENCES transaction_details(id);
  
-- 7a 1-to-1 Payment Method Description
CREATE TABLE payment_method_description (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
description TEXT,
payment_method_id INT NOT NULL UNIQUE,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- 7b 1-to-many user dengan alamat
CREATE TABLE alamat (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    jalan VARCHAR(255) NOT NULL,
    kota VARCHAR(100) NOT NULL,
    provinsi VARCHAR(100) NOT NULL,
    kode_pos VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 7c many-to-many user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE payment_method_detail (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- 3 membuat tabel kurir
CREATE TABLE kurir (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (50) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 4 menambahkan column ongkos_dasar pada tabel kurir
ALTER TABLE kurir
ADD COLUMN ongkos_dasar NUMERIC(10,2);

-- 5 merename tabel kurir menjadi shipping
RENAME TABLE kurir TO shipping;

-- 6 drop tabel shipping karena tidak dibutuhkan
DROP TABLE shipping;
