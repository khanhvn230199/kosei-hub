package routes

import (
	"net/http"

	"kosei-web/api/controllers"
)

var loginRoutes = []Route{
	Route{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
