package api

import (
	"encoding/json"
	"go-rate-limit/errs"
	"go-rate-limit/transport"
	"net/http"
)

// GetDockerName filters the request and return a randomize docker container name
func GetDockerName(w http.ResponseWriter, r *http.Request) {
	// Validate JWT Token
	token := r.Header.Get("X-Go-API-Token")

	email, err := transport.ValidateUserEmail(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate API Key
	apiKey := r.Header.Get("X-Go-API-Key")

	err = transport.ValidateAPIKey(email, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check Rate Limit
	err = transport.CheckRateLimit(email)
	if err != nil {
		requestType := http.StatusInternalServerError

		if err == errs.MaxUsageErr {
			requestType = http.StatusTooManyRequests
		}

		http.Error(w, err.Error(), requestType)
		return
	}

	dockerName := APIService.GetDockerName()
	response, err := json.Marshal(dockerName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
