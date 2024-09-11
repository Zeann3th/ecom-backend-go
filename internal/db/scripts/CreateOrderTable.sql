DROP TABLE IF EXISTS orders;

CREATE TABLE orders(
  user_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY(user_id, product_id),
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(product_id) REFERENCES products(id)
)
