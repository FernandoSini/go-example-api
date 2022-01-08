package repository

import (
	"api/src/model"
	"database/sql"
)

//repositorio de publicacoes
type Post struct {
	db *sql.DB
}

//cria um novo repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Post {
	return &Post{db}
}

//insere publicacao no db
func (repositorio Post) CriarPost(post model.Post) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(&post.Title, &post.Content, &post.AuthorId)
	if erro != nil {
		return 0, erro
	}
	lastIdInserted, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInserted), nil
}

func (repositorio Post) FindPostById(postId uint64) (model.Post, error) {
	//ficar atento pra passar as ordens corretas no scan
	linha, erro := repositorio.db.Query(
		`SELECT p.*, u.nick from posts p 
		 inner join usuarios u on u.id = p.author_id
		  WHERE p.id =?`, postId)
	if erro != nil {
		return model.Post{}, erro
	}
	defer linha.Close()
	var post model.Post

	if linha.Next() {
		if erro = linha.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt, &post.AuthorNick); erro != nil {
			return model.Post{}, erro
		}

	}
	return post, nil

}

//tras os posts dos users seguidos e do proprio usuario
func (repositorio Post) FindPosts(userId uint64) ([]model.Post, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from posts p 
	inner join usuarios u on u.id = p.author_id 
	inner join followers f on p.author_id = f.usuario_id 
	where u.id = ? or f.seguidor_id = ?
	`, userId, userId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()
	var posts []model.Post

	for linhas.Next() {
		var post model.Post
		if erro = linhas.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt, &post.AuthorNick); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (repositorio Post) UpdatePost(postId uint64, post model.Post) error {
	statement, erro := repositorio.db.Prepare("UPDATE posts set title= ?, content= ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(&post.Title, &post.Content, postId); erro != nil {
		return erro
	}
	return nil
}
func (repositorio Post) DeletePost(postId uint64) error {
	statement, erro := repositorio.db.Prepare("Delete from posts where id = ? ")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(postId); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Post) FindPostsByUser(userId uint64) ([]model.Post, error) {
	linhas, erro := repositorio.db.Query(
		`select p.*, u.nick from posts p 
	join usuarios u on u.id = p.author_id
	 where p.author_id = ?`, userId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var posts []model.Post

	for linhas.Next() {
		var post model.Post
		if erro = linhas.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt, &post.AuthorNick); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (repositorio Post) LikePost(postId uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE posts set likes = likes + 1 where id =?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(&postId); erro != nil {
		return erro
	}
	return nil
}
func (repositorio Post) UnlikePost(postId uint64) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE posts set likes = 
	 CASE WHEN likes >0 THEN 
	 likes -1 ELSE likes END 
	 where id =?`)
	 
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(&postId); erro != nil {
		return erro
	}
	return nil
}
