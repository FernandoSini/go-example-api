package model

import (
	"api/src/security"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"` //quando vc for passar o usuario pra um json e o id estiver vazio, ele não sera passado pelo json
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func (user *User) Preparar(step string) error {
	//ai ele vai formatar os dados
	if erro := user.formatar(step); erro != nil {
		return erro
	}
	//validando se os campos não estão em branco
	if erro := user.validar(step); erro != nil {
		return erro
	}

	return nil
}

//controlando os erros de campo no json
func (user *User) validar(step string) error {

	if user.Name == "" {
		return errors.New(" Nome obrigatório e não pode ser vazio")
	}

	if user.Nick == "" {
		return errors.New(" Nick obrigatório e não pode ser vazio")
	}

	if user.Email == "" {
		return errors.New(" Email obrigatório e não pode ser vazio")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New(" Email Inserido é invalido")
	}
	if step == "registerUser" && user.Password == "" {
		return errors.New(" Password obrigatório e não pode ser vazio")
	}
	return nil
}

func (user *User) formatar(step string) error {
	user.Name = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(user.Name, " "))
	user.Nick = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(user.Nick, " "))
	user.Email = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(user.Email, " "))

	if step == "registerUser" {
		passwordHashed, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHashed)
	}
	return nil
}
