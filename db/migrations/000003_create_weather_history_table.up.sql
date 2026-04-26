CREATE TABLE IF NOT EXISTS weather_history (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    city TEXT NOT NULL,
    temperature DOUBLE PRECISION NOT NULL,
    description TEXT NOT NULL,
    requested_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_weather_history_user_city_requested_at
    ON weather_history (user_id, city, requested_at DESC);

INSERT INTO weather_history (user_id, city, temperature, description, requested_at)
VALUES
    (1, 'Almaty', 18.5, 'cloudy', now()),
    (1, 'Almaty', 20.0, 'sunny', now() - interval '1 day'),
    (1, 'Astana', 10.2, 'windy', now()),
    (2, 'London', 13.7, 'rainy', now());