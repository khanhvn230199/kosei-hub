// package main

// import (
// 	"kosei-web/api"
// )

// func main() {
// 	api.Run()
// }

// docker run -d -p 3306:3306 --name mysql-docker-container -e MYSQL_ROOT_PASSWORD=gotest -e MYSQL_DATABASE=gotest -e MYSQL_USER=sergey -e MYSQL_PASSWORD=gotest mysql:8.0.34

package main

import (
	"fmt"
	"kosei-web/api/auth"
	"kosei-web/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

func createJWT() (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = 123
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Sign the token with a secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func main() {
	// Create a JWT
	user := models.User{
		Email:    "abc@gmail.com",
		Password: "aaaaaaaa",
	}
	token, err := auth.GenerateJWT(user)
	if err != nil {
		return
	}

	fmt.Println(token)
}
