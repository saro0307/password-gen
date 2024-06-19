# Password Generator

A simple Golang program to generate random passwords.

## Description

This program generates a random password with a specified length and a combination of lowercase letters, uppercase letters, digits, and special characters. It uses cryptographic random number generation to ensure the password's randomness.

## Features

- Generate random passwords
- Include lowercase letters, uppercase letters, digits, and special characters
- Cryptographically secure random number generation

## Usage

1. Clone the repository:

   ```sh
   git clone https://github.com/saro0307/password-generator
   ```

2. Navigate to the project directory:

   ```sh
   cd password-generator
   ```

3. Run the program:

   ```sh
   go run passwordgen.go
   ```

4. The generated password will be printed to the console.

## Customization

You can customize the program by editing the `passwordgen.go` file:

- Change the length of the password by modifying the `passwordLength` variable.
- Modify the character sets (`lowercase`, `uppercase`, `digits`, `special`) to include or exclude specific characters.

## Example

```go
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
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

- [GitHub](https://github.com/saro0307)
