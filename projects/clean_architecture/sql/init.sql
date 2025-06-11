USE orders;

CREATE TABLE orders (
    id INT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price  decimal(10,2)  NOT NULL,
    tax decimal(10,2)  NOT NULL,
    final_price decimal(10,2)  NOT NULL
);

INSERT INTO orders (price, tax, final_price)
VALUES 
  (100.00, 10.00, 110.00),
  (200.00, 20.00, 220.00),
  (150.50, 15.05, 165.55);