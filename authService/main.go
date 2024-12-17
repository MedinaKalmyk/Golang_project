package main

import (
	"Go/cmd/app/api"
	"Go/cmd/appconfig"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ServerID := os.Getenv("SERVERID")
	io.WriteString(w, "msg from server"+ServerID)
}

func main() {

	var c appconfig.Conf

	app := fiber.New()

	api.Routes(app)

	server := c.GetConf().Server

	app.Listen(":" + server.Port)
}
