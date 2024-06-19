package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Define the length of the password
	passwordLength := 12

	// Define character sets
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_+=<>?"

	// Combine all character sets
	allCharacters := lowercase + uppercase + digits + special

	// Generate password
	password := generatePassword(passwordLength, allCharacters)

	// Print the generated password
	fmt.Println("Generated Password:", password)
}

func generatePassword(length int, characters string) string {
	password := make([]byte, length)
	charSetLength := big.NewInt(int64(len(characters)))

	for i := range password {
		// Generate a random index
		randomIndex, err := rand.Int(rand.Reader, charSetLength)
		if err != nil {
			fmt.Println("Error generating random index:", err)
			return ""
		}

		// Select the character from the character set
		password[i] = characters[randomIndex.Int64()]
	}

	return string(password)
}
