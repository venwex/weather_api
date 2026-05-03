package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"weather_api/internal/auth"
	"weather_api/internal/models"
	"weather_api/internal/service"
	u "weather_api/internal/utils"
)

type UserHandler struct {
	Users *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		Users: svc,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.Users.GetUsers(ctx)
	if err != nil {
		log.Printf("error getting users: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Users.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("error getting user: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error decoding user: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Users.CreateUser(ctx, user)
	if err != nil {
		log.Printf("error creating user: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error decoding user: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.Users.UpdateUser(ctx, id, user)
	if err != nil {
		log.Printf("error updating user: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) { // soft Delete
	ctx := r.Context()

	id, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		u.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.Users.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("error deleting user: %v", err)
		u.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusNoContent, nil)
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	currentUser, err := auth.GetCurrentUser(r.Context())
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Users.Me(r.Context(), currentUser.ID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	u.WriteJSON(w, http.StatusOK, user)
}
