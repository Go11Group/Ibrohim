package token

import (
	"fmt"
	"log"
	"time"
	"user_service/config"
	pbAu "user_service/genproto/authentication"

	"github.com/golang-jwt/jwt"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateJWT(user *pbAu.UserToken) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	config := config.Load()
	access, err := accessToken.SignedString([]byte(config.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access Token not generated: %v", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.Id
	rftClaims["username"] = user.Username
	rftClaims["email"] = user.Email
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := refreshToken.SignedString([]byte(config.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Refresh Token not generated: %v", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func GenerateAccessToken(refreshToken string) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)

	claims, err := ExtractClaims(refreshToken)
	if err != nil {
		return "", err
	}

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	accessToken.Claims = claims
	token, err := accessToken.SignedString([]byte(config.Load().SIGNING_KEY))

	return token, err
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaims(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Load().SIGNING_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
