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
