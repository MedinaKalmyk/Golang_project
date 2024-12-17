package controllers

import (
	_ "Go/connection"
	"Go/internal/repository/userRepository"
	"Go/internal/services"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

var jwtSecretKey = []byte("very-secret-key")
var (
	errBadCredentials = errors.New("email or password is incorrect")
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	regReq := RegisterRequest{}

	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}
	user := userRepository.LogIn(regReq.Username, regReq.Password)
	if user.Username != regReq.Username {
		return errBadCredentials
	}
	compare := services.Hash{}
	fmt.Println(compare.Compare(user.Password, regReq.Password))
	fmt.Println(user.Password)
	err := compare.Compare(user.Password, regReq.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Wrong Password")
	}
	payload := jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		logrus.WithError(err).Error("JWT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(LoginResponse{AccessToken: t})
}
