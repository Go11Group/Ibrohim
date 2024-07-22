package tokens

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const accesssigningkey = "ssecca"

func GenerateAccessToken(id, username, email string) (string, error) {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["username"] = username
	claims["email"] = email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(accesssigningkey))

	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "failed to generate access token")
	}

	return newToken, nil
}
