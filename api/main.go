package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//gerando a secret key
/* func init() {
	key := make([]byte, 64)
	if _, erro := rand.Read(key); erro != nil {
		log.Fatal(erro)
	}
	//convertendo o slice de bytes em uma string base64 e salvar no env
	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
} */

func main() {
	config.Carregar()

	fmt.Println("rodando")
	r := router.Gerar()

	fmt.Printf("Server running on port %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
