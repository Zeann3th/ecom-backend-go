DROP TABLE IF EXISTS products CASCADE;

CREATE TABLE products(
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  image TEXT NOT NULL,
  price FLOAT NOT NULL,
  stock INT NOT NULL,
  sellerId INT NOT NULL,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CHECK ( stock >= 0 ),
  FOREIGN KEY(sellerId) REFERENCES users(id)
)
