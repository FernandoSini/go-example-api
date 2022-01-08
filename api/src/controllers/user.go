package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//controllers
//responsavel por controlar as operacoes, não necessariamente vao executar elas
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.User
	if erro = json.Unmarshal(reqBody, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.Preparar("registerUser"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	//passando a conexão pro repositório
	repository := repository.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repository.Criar(usuario)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuario)

}
func FindingUsers(w http.ResponseWriter, r *http.Request) {
	//usando o parametro na url?
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NovoRepositorioDeUsuarios(db)
	user, erro := repository.Find(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, user)

}
func FindUser(w http.ResponseWriter, r *http.Request) {
	paramteros := mux.Vars(r)

	userId, erro := strconv.ParseUint(paramteros["userId"], 10, 64)
	if erro != nil {
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
	user, erro := repository.FindUserById(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	paramteros := mux.Vars(r)

	userId, erro := strconv.ParseUint(paramteros["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIdInToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if userId != userIdInToken {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden Action"))
		return
	}

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
	if erro = user.Preparar("updateUser"); erro != nil {
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
	if erro = repository.Update(userId, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	userIdInToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdInToken {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden Action"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NovoRepositorioDeUsuarios(db)
	if erro = repository.DeleteUser(userId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden action"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NovoRepositorioDeUsuarios(db)

	if erro = repository.FollowUser(userId, followerId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	folllowerId, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if folllowerId == userId {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden action"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	defer db.Close()

	repository := repository.NovoRepositorioDeUsuarios(db)
	if erro = repository.UnfollowUser(userId, folllowerId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}

func FindFollowers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	followers, erro := repository.FindFollowers(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}

func UsersFollowing(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	followingUsers, erro := repository.FindFollowingUsers(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, followingUsers)

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdInToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if userIdInToken != userId {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden action"))
		return
	}

	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var password model.Password
	if erro = json.Unmarshal(reqBody, &password); erro != nil {
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
	//buscando a senha no banco para comparar sua hash com o valor vindo do json
	savedPassword, erro := repository.FindPassword(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(savedPassword, password.CurrentPassword); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New(" Not correct password"))
		return
	}

	hashedPassword, erro := security.Hash(password.NewPassword)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	//inserido a senha atualizada no banco
	if erro = repository.UpdatePassword(userId, string(hashedPassword)); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}
	responses.JSON(w, http.StatusNoContent, nil)

}
