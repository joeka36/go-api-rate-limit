package api

import (
	"net/http"
	"encoding/json"
)

func GetDockerName(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract scope to properly rate limiting for user type

	dockerName := APIService.GetDockerName()
	response, err := json.Marshal(dockerName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}