package models

import "time"

type User struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type UserCity struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	City      string    `db:"city"`
	CreatedAt time.Time `db:"created_at"`
}

type WeatherHistory struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	City        string    `db:"city"`
	Temperature float64   `db:"temperature"`
	Description string    `db:"description"`
	RequestedAt time.Time `db:"requested_at"`
}