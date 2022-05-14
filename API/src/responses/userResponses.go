package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Returns a JSON response to a request
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// Returns an error JSON response to a request
func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	JSONResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
