package config

import (
	"Go/cmd/appconfig"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var c appconfig.Conf

type db struct {
	db *sql.DB
}

var host = os.Getenv("POSTGRES_HOST")
var port = 5432
var user = os.Getenv("POSTGRES_USER")
var password = os.Getenv("POSTGRES_PASSWORD")
var dbname = os.Getenv("POSTGRES_DB")

func OppenConnect() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db

}
