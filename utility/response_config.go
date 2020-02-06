package utility

import (
	"net/http"
)

// SetContentTypeToJSON is use to set the content type of our api response to json
func SetContentTypeToJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}