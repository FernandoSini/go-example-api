package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	userId, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post model.Post
	if erro = json.Unmarshal(reqBody, &post); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	post.AuthorId = userId
	if erro = post.Preparar(); erro != nil {

		responses.Erro(w, http.StatusBadRequest, erro)
		return

	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NovoRepositorioDePublicacoes(db)
	post.ID, erro = repository.CriarPost(post)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, post)

}

func FindPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, erro := strconv.ParseUint(params["postId"], 10, 64)
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
	repository := repository.NovoRepositorioDePublicacoes(db)
	post, erro := repository.FindPostById(postId)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, post)

}

func FindPosts(w http.ResponseWriter, r *http.Request) {
	userId, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NovoRepositorioDePublicacoes(db)
	posts, erro := repository.FindPosts(userId)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	postId, erro := strconv.ParseUint(params["postId"], 10, 64)
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

	repository := repository.NovoRepositorioDePublicacoes(db)
	postSavedInDb, erro := repository.FindPostById(postId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postSavedInDb.AuthorId != userId {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden action"))
		return
	}
	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post model.Post
	if erro = json.Unmarshal(reqBody, &post); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePost(postId, post); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

	userId, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	postId, erro := strconv.ParseUint(params["postId"], 10, 64)
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

	repository := repository.NovoRepositorioDePublicacoes(db)
	postSavedInDb, erro := repository.FindPostById(postId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postSavedInDb.AuthorId != userId {
		responses.Erro(w, http.StatusForbidden, errors.New(" Forbidden action"))
		return
	}

	if erro = repository.DeletePost(postId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func FindPostsByUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repository.NovoRepositorioDePublicacoes(db)
	posts, erro := repository.FindPostsByUser(userId)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, erro := strconv.ParseUint(params["postId"], 10, 64)
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

	repository := repository.NovoRepositorioDePublicacoes(db)
	if erro = repository.LikePost(postId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, erro := strconv.ParseUint(params["postId"], 10, 64)
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

	repository := repository.NovoRepositorioDePublicacoes(db)
	if erro = repository.UnlikePost(postId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
