package jwttoken

import (
	"fmt"
	"log"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/configs"
)

// GenerateToken generates token when user connects himself
func GenerateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["expire"] = time.Now().Add(time.Hour * time.Duration(configs.GetInt("TOKEN_HOUR_LIFESPAN"))).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println(token.SignedString([]byte(configs.GetString("API_SECRET"))))

	return token.SignedString([]byte(configs.GetString("API_SECRET")))
}

// ExtractToken retrieves the token send by the user in every request
func ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	// If the token doesn not start with a good prefix: like Bearer, Token, ..., we return empty string
	if !strings.HasPrefix(bearerToken, configs.GetString("TOKEN_PREFIX")) {
		return ""
	}

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// TokenValid check if a token is valid and not expired
func TokenValid(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)

	token , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetString("API_SECRET")), nil
	})

	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return err
}

// ExtractTokenID extracts the user ID on a Token
func ExtractClaims(ctx *gin.Context) (jwt.MapClaims, error) {
	tokenString := ExtractToken(ctx)

	token , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetString("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if err != nil {
			return nil, err
		}
		return claims, nil
	}

	return nil, nil
}