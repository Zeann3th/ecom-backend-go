DROP TABLE IF EXISTS orders;

CREATE TABLE orders(
  userId INT NOT NULL,
  productId INT NOT NULL,
  quantity INT NOT NULL,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY(userId, productId),
  FOREIGN KEY(userId) REFERENCES users(id),
  FOREIGN KEY(productId) REFERENCES products(id)
)
