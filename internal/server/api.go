package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukkak/go-playground/internal/util"
)

func RegisterCommonRoutes(r *mux.Router) {
	r.HandleFunc("/info", GetAppInfo).Methods(http.MethodGet)
	r.HandleFunc("/ping", PingHandler).Methods(http.MethodGet)
}

func GetAppInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	util.WriteJson(w, &AppInfo{Name: "example-service", Version: "develop"})
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	util.WriteText(w, "pong")
}
