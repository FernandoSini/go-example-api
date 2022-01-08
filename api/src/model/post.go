package model

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// vai chamar os metodos validar e formatar a publicacao recebida
func (post *Post) Preparar() error {
	//formatar o post
	post.formatar()

	//chamando o validar e verificar se os campos estão preenchidos
	if erro := post.validar(); erro != nil {
		return erro
	}

	return nil

}

func (post *Post) validar() error {
	if post.Title == "" {
		return errors.New(" Title needed for post")
	}
	if post.Content == "" {
		return errors.New(" Content needed for post")
	}
	return nil
}

func (post *Post) formatar() {
	post.Title = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(post.Title, " "))
	post.Content = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(post.Content, " "))

}
