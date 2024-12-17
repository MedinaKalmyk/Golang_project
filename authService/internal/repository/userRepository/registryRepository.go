package userRepository

import (
	connection "Go/connection"
	"fmt"
	"github.com/google/uuid"
)

func Registry(username, password, surname, name, middleName, email, phoneNumber, age string) error {
	user_id := uuid.New()
	fmt.Println(connection.OppenConnect().Begin())
	tx, err := connection.OppenConnect().Begin()
	fmt.Println("Привет")
	fmt.Println(err)

	if err != nil {
		fmt.Println("Transaction is fail")
		return err
	}
	_, err = tx.Exec("INSERT INTO users VALUES (default, $1, $2, $3)",
		user_id, username, password)
	if err != nil {
		fmt.Println("Rollback table user")
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("INSERT INTO user_information (user_id, surname, name, middle_name, age, phone_number, email) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user_id, surname, name, middleName, age, phoneNumber, email)
	if err != nil {
		fmt.Println("Rollback table user_information")
		fmt.Println(err)

		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
