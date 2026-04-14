package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type healthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Message  string `json:"message,omitempty"`
}

// Endpoint used to check the readiness and/or liveness (health) of the server.
func (a *AcmednsAPI) healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	if err := a.DB.GetBackend().Ping(); err != nil {
		a.Logger.Warnw("Health check: database ping failed", "error", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(healthResponse{
			Status:   "degraded",
			Database: "error",
			Message:  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(healthResponse{
		Status:   "ok",
		Database: "ok",
	})
}
