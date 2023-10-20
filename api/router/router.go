package router

import (
	"kosei-web/api/router/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// New routes
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(accessControlMiddleware)
	return routes.SetupRoutesWithMiddlewares(r)
}

// access control and  CORS middleware
func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
