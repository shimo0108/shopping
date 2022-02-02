CREATE TABLE IF NOT EXISTS orders(
  id serial PRIMARY KEY,
  total_price int NOT NULL DEFAULT 0,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp
);
