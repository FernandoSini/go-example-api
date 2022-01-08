package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeLogin = Route{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Function:           controllers.Login,
	NeedAuthentication: false,
}
