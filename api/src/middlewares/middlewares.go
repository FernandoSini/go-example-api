package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

//log escreve informacoes da requisicao no terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

//middleware é uma camada entre a requisição e a resposta
//handlerFunc é uma outra forma de chamar (w http.ResponseWriter, r*http.Request)
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidateToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunction(w, r)
	}
}
