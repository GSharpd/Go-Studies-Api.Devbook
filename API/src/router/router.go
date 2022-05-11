package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Returns a new router with the configured routes
func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	return routes.ConfigureRoutes(router)
}
