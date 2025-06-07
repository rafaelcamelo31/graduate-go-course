CREATE TABLE orders (
    id INT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price  decimal(10,2)  NOT NULL,
    tax decimal(10,2)  NOT NULL,
    final_price decimal(10,2)  NOT NULL
)