package rest

import (
	"log/slog"
	"net/http"

	"github.com/Konil-Startup/go-backend/internal/service"
	"github.com/go-chi/chi"
)

type RestHandler struct {
	Service service.Service
	l       *slog.Logger
}

func New(s service.Service, l *slog.Logger) *RestHandler {
	return &RestHandler{
		Service: s,
		l:       l,
	}
}

func Routes(h *RestHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/user/{user_id}", h.UserByID)
	r.Get("/user/email/{email}", h.UserByEmail)
	r.Post("/user", h.SaveUser)

	return WriteToConsole(r)
}
