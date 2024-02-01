package util

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiStatusResponse struct {
	success bool
}

type ApiErrorResponse struct {
	message string
}

func ReadRequestPathVars(r *http.Request, v any) error {
	bytes, err := json.Marshal(mux.Vars(r))
	if err == nil {
		slog.Info("mux vars", "bytes", bytes)
		return json.Unmarshal(bytes, v)
	} else {
		return err
	}
}

func ReadRequestBody(r *http.Request, v any) error {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	slog.Info("Request Body", "bytes", bytes)
	return json.Unmarshal(bytes, v)
}

//func ReadRequestBody(r *http.Request, v any) error {
//	return json.NewDecoder(r.Body).Decode(v)
//}

func WriteStatusJson(w http.ResponseWriter, success bool) {
	WriteJson(w, ApiStatusResponse{success: success})
}

func WriteJson(w http.ResponseWriter, value any) {
	bytes, err := json.Marshal(value)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	} else {
		WriteError(w, "Error serializing the response")
	}
}

func WriteText(w http.ResponseWriter, value string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(value))
}

func WriteError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(420)
	bytes, _ := json.Marshal(ApiErrorResponse{message: message})
	w.Write(bytes)
}

func WriteErr(w http.ResponseWriter, e error) {
	WriteError(w, e.Error())
}
