package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//abrindo conexao com banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.UrlConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
