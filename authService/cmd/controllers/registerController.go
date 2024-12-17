package controllers

import (
	"Go/internal/models"
	"Go/internal/repository/userRepository"
	"Go/internal/services"
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var successCredentials = "Success Register"

type (
	AuthHandler struct{}
	AuthStorage struct {
		users map[string]models.User
	}
)

type RegistryResponse struct {
	Message string `json:"message"`
}
type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Surname     string `json:"surname"`
	Name        string `json:"name"`
	MiddleName  string `json:"middle_name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (h *AuthHandler) Registration(c *fiber.Ctx) error {
	regReq := RegisterRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}
	fmt.Println(regReq.Username)
	compare := services.Hash{}
	var password, _ = compare.Generate(regReq.Password)
	fmt.Println(password)

	userRepository.Registry(regReq.Username, password, regReq.Surname, regReq.Name, regReq.MiddleName, regReq.Email, regReq.PhoneNumber, regReq.Age)
	return c.JSON(RegistryResponse{Message: successCredentials})

}
