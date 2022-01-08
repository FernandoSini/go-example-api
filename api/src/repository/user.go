package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

//representa o repositorio de usuarios
type User struct {
	db *sql.DB
}

//faz a comunicacao direta com o banco de dados
func NovoRepositorioDeUsuarios(db *sql.DB) *User {
	return &User{db}
}

//insere um usuario no banco de dados
func (repositorio User) Criar(user model.User) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (name, nick,email, password) VALUES (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastIDInserted), nil

}

//buscando usuarios por nome ou nick
func (repositorio User) Find(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nomeounick%

	linhas, erro := repositorio.db.Query("SELECT id, name, nick, email, createdAt from usuarios WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []model.User

	for linhas.Next() {
		var user model.User
		if erro = linhas.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil

}

//find user by id
func (repositorio User) FindUserById(ID uint64) (model.User, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, name, nick, email, createdAt from usuarios WHERE id = ? ",
		ID,
	)
	if erro != nil {
		return model.User{}, erro
	}
	defer linhas.Close()

	var user model.User

	if linhas.Next() {
		if erro = linhas.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return model.User{}, erro
		}

	}
	return user, nil
}

func (repositorio User) Update(ID uint64, user model.User) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios set name = ?, nick = ?, email = ? where id = ? ",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(
		&user.Name,
		&user.Nick,
		&user.Email,
		ID,
	); erro != nil {
		return erro
	}
	return nil

}

func (repositorio User) DeleteUser(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE from usuarios WHERE id =?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil

}
func (repositorio User) FindUserByEmail(email string) (model.User, error) {
	linha, erro := repositorio.db.Query("select id, password from usuarios where email = ?", email)
	if erro != nil {
		return model.User{}, erro
	}
	defer linha.Close()

	var user model.User
	if linha.Next() {
		if erro = linha.Scan(&user.ID, &user.Password); erro != nil {
			return model.User{}, erro
		}
	}
	return user, nil
}

func (repository User) FollowUser(userId, followerId uint64) error {
	statement, erro := repository.db.Prepare("INSERT ignore INTO followers (usuario_id, seguidor_id) values(?,?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(&userId, &followerId); erro != nil {
		return erro
	}

	return nil
}

func (repository User) UnfollowUser(userId, followerId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM followers where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(&userId, &followerId); erro != nil {
		return erro
	}
	return nil
}

func (repository User) FindFollowers(userId uint64) ([]model.User, error) {
	linhas, erro := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.createdAt 
		from usuarios u inner join 
		followers f on u.id = f.seguidor_id where f.usuario_id = ?`, userId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var followers []model.User
	if linhas.Next() {
		var follower model.User
		if erro = linhas.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		followers = append(followers, follower)
	}
	return followers, nil

}
func (repository User) FindFollowingUsers(userId uint64) ([]model.User, error) {
	linhas, erro := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.createdAt 
		from usuarios u inner join 
		followers f on u.id = f.usuario_id where f.seguidor_id = ?`, userId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var followingUsers []model.User
	if linhas.Next() {
		var followingUser model.User
		if erro = linhas.Scan(
			&followingUser.ID,
			&followingUser.Name,
			&followingUser.Nick,
			&followingUser.Email,
			&followingUser.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		followingUsers = append(followingUsers, followingUser)
	}
	return followingUsers, nil

}

func (repository User) FindPassword(userId uint64) (string, error) {
	linha, erro := repository.db.Query("select password from usuarios where id =? ", userId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var user model.User
	if linha.Next() {
		if erro = linha.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}
	return user.Password, nil
}

//altera a senha de um usuario
func (repository User) UpdatePassword(userId uint64, password string) error {
	statement, erro := repository.db.Prepare("UPDATE usuarios set password = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(password, userId); erro != nil {

		return erro
	}
	return nil
}
