CREATE TABLE IF NOT EXISTS posts (
  id SERIAL PRIMARY KEY,
  title text,
  user_id integer REFERENCES users (id) ON DELETE CASCADE,
  inserted_at timestamp default now(),
  updated_at timestamp default now()
);
