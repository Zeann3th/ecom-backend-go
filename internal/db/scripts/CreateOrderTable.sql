DROP TABLE IF EXISTS orders;

CREATE TABLE orders(
  user_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  total FLOAT NOT NULL,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY(user_id, product_id),
  FOREIGN KEY(user_id) REFERENCES users(user_id),
  FOREIGN KEY(product_id) REFERENCES products(product_id)
)
