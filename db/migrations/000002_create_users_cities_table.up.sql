CREATE TABLE IF NOT EXISTS user_cities (
   id SERIAL PRIMARY KEY,
   user_id INTEGER NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
   city TEXT NOT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

   UNIQUE (user_id, city)
);

INSERT INTO user_cities (user_id, city)
VALUES
    (1, 'Almaty'),
    (1, 'Astana'),
    (2, 'London');