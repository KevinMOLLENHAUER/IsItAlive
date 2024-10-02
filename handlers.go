package main

import(
  "net/http"
  "log/slog"
  "io"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A simple healthcheck handler

	slog.Info("Receive call on healthcheck handler")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func RegisterTargetHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteTargetHandler(w http.Response)

func TargetStatusHandler(w http.ResponseWriter, r *http.Request) {
  
}
