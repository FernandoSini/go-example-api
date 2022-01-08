package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//é responsável por autenticar um usuário na api
func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user model.User
	if erro = json.Unmarshal(reqBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repository.NovoRepositorioDeUsuarios(db)
	savedUserInDB, erro := repository.FindUserByEmail(user.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	//verificando se o hash bate com a senha vindo na requisicao
	if erro = security.VerifyPassword(savedUserInDB.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CriarToken(savedUserInDB.ID)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
	//responses.JSON(w, http.StatusAccepted, token)

}
