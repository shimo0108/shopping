CREATE TABLE IF NOT EXISTS restaurants (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  fee int NOT NULL DEFAULT 0,
  time_required int NOT NULL,
  created_at  timestamp not null default current_timestamp,
  updated_at  timestamp not null default current_timestamp
);
