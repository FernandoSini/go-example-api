package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//Retorna uma resposta em json em uma requisicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

//retorna um erro em json
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})

}
