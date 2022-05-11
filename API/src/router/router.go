package router

import (
	"github.com/gorilla/mux"
)

// Returns a new router with the configured routes
func CreateRouter() *mux.Router {
	return mux.NewRouter()
}
