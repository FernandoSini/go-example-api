package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//vai gerenciar as rotas
//vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	return routes.Configurar(router)
}
