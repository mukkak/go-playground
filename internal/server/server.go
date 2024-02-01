package server

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() error {
	router := mux.NewRouter()
	RegisterCommonRoutes(router)

	server := &http.Server{
		Addr:      "0.0.0.0:8080",
		Handler:   router,
		TLSConfig: nil,
	}
	slog.Info("Starting server at " + server.Addr)
	return server.ListenAndServe()
}
