CREATE TABLE IF NOT EXISTS restaurants (
  id varchar(255) PRIMARY KEY,
  name varchar(255) NOT NULL,
  fee int NOT NULL DEFAULT 0,
  time_required int NOT NULL,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp
);


CREATE TABLE IF NOT EXISTS foods (
  id varchar(255) PRIMARY KEY,
  restaurant_id varchar(255) NOT NULL,
  name varchar(100) NOT NULL,
  price int NOT NULL DEFAULT 0,
  description varchar(255) NOT NULL,
  created_at  timestamp NOT NULL DEFAULT current_timestamp,
  updated_at  timestamp NOT NULL default current_timestamp,
  foreign key(restaurant_id) references restaurants(id)
);


CREATE TABLE IF NOT EXISTS orders(
  id varchar(255) PRIMARY KEY,
  total_price int NOT NULL DEFAULT 0,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS line_foods(
  id varchar(255) PRIMARY KEY,
  restaurant_id varchar(255) NOT NULL,
  food_id varchar(255) NOT NULL,
  order_id varchar(255),
  count int NOT NULL DEFAULT 0,
  active  boolean default false,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp,
  foreign key(restaurant_id) references restaurants(id),
  foreign key(food_id) references foods(id)
);
