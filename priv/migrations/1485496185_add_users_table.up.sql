CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text UNIQUE NOT NULL,
  inserted_at timestamp default now(),
  updated_at timestamp default now()
);
