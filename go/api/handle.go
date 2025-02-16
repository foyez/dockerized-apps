package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/health" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rsp := map[string]any{
			"status":    "ok",
			"timestamp": time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(rsp)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		rsp := map[string]any{
			"error": "Not Found",
		}
		json.NewEncoder(w).Encode(rsp)
	}
}
