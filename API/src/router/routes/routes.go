package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Represents the API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

// Configure all user routes to the given router
func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {

		if route.RequiresAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method) // This nests the function given as a parameter to the Authenticate method from middlewares
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
