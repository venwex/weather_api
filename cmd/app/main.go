package main

import (
	"log"
	"net/http"
	"weather_api/internal/handler"
	"weather_api/internal/repository"
	"weather_api/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "postgres://weather:weatherpass@localhost:5433/weather_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Connection error to the DB: %v", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	mux := initRoutes(handler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func initRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", h.Users.GetUsers) // user logic
	mux.HandleFunc("GET /users/{id}", h.Users.GetUserByID)
	mux.HandleFunc("POST /users", h.Users.CreateUser)
	mux.HandleFunc("PUT /users/{id}", h.Users.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", h.Users.DeleteUser)

	mux.HandleFunc("GET /users/{id}/cities", h.Cities.GetCities) // city logic
	mux.HandleFunc("POST /users/{id}/cities", h.Cities.AddCity)
	mux.HandleFunc("DELETE /users/{id}/cities/{city_id}", h.Cities.DeleteCity)

	//mux.HandleFunc("GET /users/{id}/weather")
	//mux.HandleFunc("GET /users/{id}/weather/history?city=Almaty")
	return mux
}
