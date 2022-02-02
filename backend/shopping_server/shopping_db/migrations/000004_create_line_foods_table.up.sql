CREATE TABLE IF NOT EXISTS line_foods(
  id serial PRIMARY KEY,
  restaurant_id INT NOT NULL,
  food_id INT NOT NULL,
  order_id INT NOT NULL,
  count int NOT NULL DEFAULT 0,
  active  boolean default false,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp,
  foreign key(restaurant_id) references restaurants(id),
  foreign key(food_id) references foods(id),
  foreign key(order_id) references orders(id)
);
