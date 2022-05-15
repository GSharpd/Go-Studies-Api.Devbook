package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// Verifies if client making a request is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		nextFunction(w, r)
	}
}

//Logger logs informations from requests in the terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s, %s, %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}
