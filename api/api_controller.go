package api

import (
	"encoding/json"
	"fmt"
	"go-rate-limit/transport"
	"net/http"
)

func GetDockerName(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract scope to properly rate limiting for user type
	token := r.Header.Get("X-Go-API-Token")

	email, err := transport.ValidateUserEmail(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(email)

	dockerName := APIService.GetDockerName()
	response, err := json.Marshal(dockerName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
