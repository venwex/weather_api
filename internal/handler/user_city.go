package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"weather_api/internal/handler/dto"
	"weather_api/internal/service"
	u "weather_api/internal/utils"
)

type CityHandler struct {
	Cities *service.CityService
}

func NewCityHandler(cities *service.CityService) *CityHandler {
	return &CityHandler{Cities: cities}
}

func (h *CityHandler) AddCity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := u.GetID(r) // user_id
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	var req dto.CreateCityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error decoding city: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	city, err := h.Cities.AddCity(ctx, id, req.City)
	if err != nil {
		log.Printf("error adding city: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusCreated, city)
}

func (h *CityHandler) GetCities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cities, err := h.Cities.GetCities(ctx, id)
	if err != nil {
		log.Printf("error getting cities: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, cities)
}

func (h *CityHandler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cityID, err := u.GetCityID(r)
	if err != nil || cityID <= 0 {
		log.Printf("error getting city id: %v", err)
		u.WriteError(w, http.StatusBadRequest, "invalid city id")
		return
	}

	if err := h.Cities.DeleteCity(ctx, userID, cityID); err != nil {
		log.Printf("error deleting city: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusNoContent, nil)
}
