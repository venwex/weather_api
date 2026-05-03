package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"weather_api/internal/auth"
	"weather_api/internal/client"
	"weather_api/internal/handler"
	"weather_api/internal/middleware"
	"weather_api/internal/repository"
	"weather_api/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	ttlRaw := os.Getenv("ACCESS_TOKEN_TTL")
	if ttlRaw == "" {
		ttlRaw = "30m"
	}

	accessTTL, err := time.ParseDuration(ttlRaw)
	if err != nil {
		log.Fatal(err)
	}

	tokenManager := auth.NewTokenManager(jwtSecret, accessTTL)

	db, err := sqlx.Open("postgres", "postgres://weather:weatherpass@localhost:5433/weather_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Connection error to the DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	repo := repository.NewRepository(db)

	weatherClient := client.NewOpenMeteoClient(nil)

	svc := service.NewService(repo, weatherClient, tokenManager)

	h := handler.NewHandler(svc)

	mux := initRoutes(h, tokenManager)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func initRoutes(h *handler.Handler, tokenManager *auth.TokenManager) *http.ServeMux {
	mux := http.NewServeMux()

	authMW := middleware.AuthMiddleware(tokenManager)
	adminMW := middleware.RequireRole("admin")

	protected := func(fn http.HandlerFunc) http.Handler {
		return authMW(http.HandlerFunc(fn))
	}

	adminOnly := func(fn http.HandlerFunc) http.Handler {
		return authMW(adminMW(http.HandlerFunc(fn)))
	}

	mux.HandleFunc("POST /auth/register", h.Auth.Register)
	mux.HandleFunc("POST /auth/login", h.Auth.Login)

	mux.Handle("GET /users/me", protected(h.Users.Me))

	mux.Handle("POST /cities", protected(h.Cities.AddCity))
	mux.Handle("GET /cities", protected(h.Cities.GetCities))
	mux.Handle("DELETE /cities/{city_id}", protected(h.Cities.DeleteCity))

	mux.Handle("GET /weather", protected(h.Weather.GetUserWeather))
	mux.Handle("GET /weather/history", protected(h.Weather.GetWeatherHistory))

	mux.Handle("GET /users", adminOnly(h.Users.GetUsers))
	mux.Handle("GET /users/{id}", adminOnly(h.Users.GetUserByID))
	mux.Handle("DELETE /users/{id}", adminOnly(h.Users.DeleteUser))

	return mux
}
