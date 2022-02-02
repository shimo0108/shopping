CREATE TABLE IF NOT EXISTS restaurants (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  fee int NOT NULL DEFAULT 0,
  time_required int NOT NULL,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp
);
SELECT setval('restaurants_id_seq', 4, false);


CREATE TABLE IF NOT EXISTS foods (
  id serial PRIMARY KEY,
  restaurant_id INT NOT NULL,
  name varchar(100) NOT NULL,
  price int NOT NULL DEFAULT 0,
  description varchar(255) NOT NULL,
  created_at  timestamp NOT NULL DEFAULT current_timestamp,
  updated_at  timestamp NOT NULL default current_timestamp,
  foreign key(restaurant_id) references restaurants(id)
);
SELECT setval('foods_id_seq', 37, false);


CREATE TABLE IF NOT EXISTS orders(
  id serial PRIMARY KEY,
  total_price int NOT NULL DEFAULT 0,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp
);

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
