CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  amount FLOAT,
  status VARCHAR(255),
  method VARCHAR(255)
);