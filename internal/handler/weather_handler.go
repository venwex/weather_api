package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"weather_api/internal/service"
	u "weather_api/internal/utils"
)

type WeatherHandler struct {
	Weather *service.WeatherService
}

func NewWeatherHandler(weather *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		Weather: weather,
	}
}

func (h *WeatherHandler) GetUserWeather(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.Weather.GetUserWeather(ctx, userID)
	if err != nil {
		log.Printf("error getting user weather: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, response)
}

func (h *WeatherHandler) GetWeatherHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	city := strings.TrimSpace(r.URL.Query().Get("city"))
	if city == "" {
		u.WriteError(w, http.StatusBadRequest, "city query param is required")
		return
	}

	limit := 10
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			u.WriteError(w, http.StatusBadRequest, "invalid limit")
			return
		}

		limit = parsedLimit
	}

	response, err := h.Weather.GetWeatherHistory(ctx, userID, city, limit)
	if err != nil {
		log.Printf("error getting weather history: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, response)
}
