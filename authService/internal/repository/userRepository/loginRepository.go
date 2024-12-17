package userRepository

import (
	_ "Go/connection"
	connection "Go/connection"
	"Go/internal/models"
)

// авторизация пользователя
func LogIn(username, password string) models.User {
	stmt := "SELECT username, password FROM users WHERE username = $1 "
	row := connection.OppenConnect().QueryRow(stmt, username)

	bk := new(models.User)
	err := row.Scan(&bk.Username, &bk.Password)

	if err != nil {
		return models.User{}
	}
	return models.User{bk.Username, bk.Password}
}
