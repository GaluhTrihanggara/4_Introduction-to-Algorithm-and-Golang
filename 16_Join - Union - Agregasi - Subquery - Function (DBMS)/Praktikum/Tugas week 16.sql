CREATE DATABASE altaonlineshop;
USE altaonlineshop;
DROP DATABASE altaonlineshop;
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
-- 2a Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name FROM users WHERE gender = 'Male';

-- 2c Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT *
FROM users
WHERE name LIKE '%a%'
  AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY);

-- 2d Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) AS jumlah_user_perempuan FROM users WHERE gender = 'Female';

-- 2e Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY name ASC;

-- 1h Insert 5 user pada tabel user.
INSERT INTO users (name, addres, date_of_birth, status_user, gender)
VALUES 
    ('John Doe', '123 Main St', '1990-01-01', 'active', 'Male'),
    ('Jane Doe', '456 Maple Ave', '1995-05-05', 'active', 'Female'),
    ('Bob Smith', '789 Oak St', '1985-03-15', 'inactive', 'Male'),
    ('Alice Johnson', '234 Elm Rd', '2000-12-31', 'active', 'Female'),
    ('Tom Jones', '567 Pine Blvd', '1975-08-20', 'inactive', 'Male');
    
CREATE TABLE product (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (100),
code VARCHAR (50) NOT NULL,
status SMALLINT (5),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- 2b Tampilkan product dengan id = 3.
SELECT name, code, status
FROM product
WHERE id = 3;

-- 1c Insert 2 product dengan product type id = 1, dan operators id = 3. 
INSERT INTO product (name, code, status, product_types_id, operator_id)
VALUES
  ('Product A', 'A01', 1, 1, 3),
  ('Product B', 'B01', 1, 1, 3);

-- 1d Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO product (name, code, status, product_types_id, operator_id)
VALUES
('Product C', 'C01', 1, 2, 1),
('Product D', 'D01', 1, 2, 1),
('Product E', 'E01', 1, 2, 1);

-- 1e Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO product (name, code, status, product_types_id, operator_id)
VALUES
  ('Product C', 'C01', 1, 3, 4),
  ('Product D', 'D01', 1, 3, 4),
  ('Product E', 'E01', 1, 3, 4);

-- 2f Tampilkan 5 data pada data product
SELECT * FROM product LIMIT 5;

-- 3a Ubah data product id 1 dengan nama ‘product dummy’.
-- namun sayangnya pada data ini terjadi kesalahan id nya yang dimana pada id nya tidak dimulai dari 1 namun dimulai dari 3
UPDATE product
SET name = 'product dummy'
WHERE id = 3;

-- 4a DELETE FROM product WHERE id = 1;
-- namun sayangnya pada data ini terjadi kesalahan id nya yang dimana pada id nya tidak dimulai dari 1 namun dimulai dari 3
DELETE FROM product WHERE id = 3;

CREATE TABLE product_descriptions (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
description TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- 1f Insert product description pada setiap product.
INSERT INTO product_descriptions (product_id, description)
VALUES 
(3, 'Ini adalah deskripsi untuk product dengan id 1'),
(4, 'Ini adalah deskripsi untuk product dengan id 2'),
(5, 'Ini adalah deskripsi untuk product dengan id 3'),
(6, 'Ini adalah deskripsi untuk product dengan id 4');

CREATE TABLE product_types (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (100),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 1b Insert 3 product type.
INSERT INTO product_types (name) VALUES
  ('Product Type 1'),
  ('Product Type 2'),
  ('Product Type 3');

-- 4b Delete pada pada tabel product dengan product type id 1.
DELETE FROM product
WHERE product_types_id = 1;


CREATE TABLE operators (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR (100),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
Drop TABLE operators;
-- 1a Insert 5 operators pada table operators.
INSERT INTO operators (name) VALUES
  ('Operator 1'),
  ('Operator 2'),
  ('Operator 3'),
  ('Operator 4'),
  ('Operator 5');


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
  
ALTER TABLE product_descriptions
ADD COLUMN product_id INT(11),
ADD CONSTRAINT FK_product_description_product
  FOREIGN KEY (product_id)
  REFERENCES product(id);


CREATE TABLE payment_methods (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name varchar (100) not null,
status SMALLINT (5),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 1g Insert 3 payment methods.
INSERT INTO payment_methods (name, status)
VALUES 
  ('Credit Card', 1),
  ('Bank Transfer', 1),
  ('PayPal', 1);
select * from payment_methods;
CREATE TABLE transactions (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
status VARCHAR (25),
total_qty INT (15) NOT NULL,
total_price NUMERIC (25,2),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 1i Insert 3 transaksi di masing-masing user.
INSERT INTO transactions (user_id, payment_methods_id, status, total_qty, total_price)
VALUES
  (1, 1, 'completed', 2, 250000),
  (2, 1, 'cancelled', 1, 100000),
  (3, 1, 'completed', 3, 450000);
UPDATE transactions
SET user_id = 3, status = 'completed'
WHERE id = 3;


CREATE TABLE transaction_details (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  status VARCHAR(20),
  qty INT(15),
  price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 1j Insert 3 product di masing-masing transaksi.
INSERT INTO transaction_details (product_id, status, qty, price)
VALUES
(3, 'pending', 1, 1000.00),
(4, 'pending', 2, 2000.00),
(5, 'pending', 1, 1500.00);

-- 3b Update qty = 3 pada transaction detail dengan product id = 1. 
-- namun sayangnya pada data ini terjadi kesalahan id nya yang dimana pada id nya tidak dimulai dari 1 namun dimulai dari 3
UPDATE transaction_details SET qty = 3 WHERE product_id = 3;

-- membuat foreignkey agar terhubung dengan tabel product
ALTER TABLE transaction_details
ADD COLUMN product_id INT(11),
ADD CONSTRAINT FK_transaction_details_product
  FOREIGN KEY (product_id)
  REFERENCES product(id);

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

-- Join, Union, Sub query, Function
-- 1
SELECT transactions.id, transactions.user_id, users.name, transactions.created_at
FROM transactions
JOIN users ON transactions.user_id = users.id
WHERE transactions.user_id IN (1, 2)
ORDER BY transactions.created_at DESC;

-- 2 Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) as total_harga
FROM transactions
WHERE user_id = 1;

-- 3 Tampilkan total transaksi dengan product type 2.
SELECT 
  SUM(td.price) AS total_price
FROM 
  transaction_details td
  INNER JOIN product p ON td.product_id = p.id
WHERE 
  p.product_types_id = 2;

-- 4 Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT product.*, product_types.name AS product_types_name
FROM product
JOIN product_types ON product.product_types_id = product_types.id;

-- 5 Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT transactions.*, users.*
FROM transactions
INNER JOIN users ON transactions.user_id = users.id;

SELECT td.*, prod.name AS product_name
FROM transaction_details td
INNER JOIN product prod ON td.product_id = prod.id;

-- 6 
DELIMITER $$
CREATE FUNCTION delete_transaction(trans_id INT) RETURNS BOOLEAN 
BEGIN 
    DECLARE row_count INT;
    DELETE FROM transaction_details WHERE transaction_id = trans_id;
    SET row_count = ROW_COUNT();
    IF row_count > 0 THEN
        RETURN TRUE;
    ELSE
        RETURN FALSE;
    END IF;
END$$
DELIMITER ;

-- 7
DELIMITER $$
CREATE FUNCTION update_total_qty(trx_id INT)
RETURNS INT
BEGIN
  DECLARE total_qty INT;
  
  SELECT SUM(qty) INTO total_qty
  FROM transaction_detail
  WHERE transaction_id = trx_id;
  
  UPDATE transactions
  SET total_qty = total_qty
  WHERE id = trx_id;
  
  RETURN total_qty;
END$$
DELIMITER ;

-- 8 
SELECT *
FROM product
WHERE id NOT IN (
  SELECT product_id
  FROM transaction_details
);
