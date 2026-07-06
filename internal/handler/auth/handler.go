package auth

import (
	"net/http"

	"github.com/amirabaris/go-auth/internal/handler/httputil"
	"github.com/amirabaris/go-auth/internal/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

type credentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req credentialsRequest
	if err := httputil.DecodeJSON(r, &req); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.svc.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "error")
		return
	}

	httputil.WriteJSON(w, http.StatusCreated, AuthResponse{Token: token})
}
