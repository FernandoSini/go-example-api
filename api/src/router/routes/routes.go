package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//struct que vai representar todas as rotas da api,
// a estrutura sera a mesma
type Route struct {
	URI                string
	Metodo             string
	Function           func(w http.ResponseWriter, r *http.Request)
	NeedAuthentication bool
}

//configurando as rotas para passar pro arquivou router.go
func Configurar(r *mux.Router) *mux.Router {
	routes := routesUser
	routes = append(routes, routeLogin)
	routes = append(routes, routePosts...)

	//percorrendo o slice de rotas
	for _, route := range routes {
		//adicionando middleware
		if route.NeedAuthentication {

			//fazendo o middleware igual no javascript(node.js)
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Metodo)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Metodo)
		}
	}

	return r

}
