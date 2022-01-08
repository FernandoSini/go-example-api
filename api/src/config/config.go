package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//string conexao com mysql
	UrlConexao = ""
	Porta      = 0
	//chave que será usada pra assinar o token
	SecretKey []byte
)

//inicializar variaveis de ambiente
func Carregar() {
	//lendo as variaveis de ambiente
	var erro error

	//matando a aplicacao caso não consiga ler as variaveis de ambiente
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) //parseint pegando a string e convertendo pra int
	if erro != nil {
		Porta = 9000
	}
	UrlConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
