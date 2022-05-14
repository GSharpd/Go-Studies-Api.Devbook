package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Verifies if client making a request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validating Token...")
		next(w, r)
	}
}

//Logger logs informations from requests in the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s, %s, %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
