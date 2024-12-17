package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Hash implements root.Hash
type Hash struct{}

// Generate a salted hash for the input string
func (c *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hash := string(hashedBytes[:])
	return hash, nil
}

// Compare string to generated hash
func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	//return bcrypt.CompareHashAndPassword(existing, incoming)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	if err != nil {
		fmt.Printf("OK got error: %v\n", err)
	} else {
		fmt.Printf("Error got error: %v\n", nil)
	}
	fmt.Println()
	return err
}
