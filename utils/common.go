package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/entity"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Function used to parse API response
func ResponseJSON(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:    errCode,
		Message: msg,
		Data:    data,
	})
}

// Allow to cypher a given word
func HashString(password string) (string, error) {
	if password == "" {
		return "", entity.ErrInvalidPassword
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare a cyphered word and a plain word
func CheckHashedString(plain_word, hashed_word string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_word), []byte(plain_word))

	if err ==  bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password associated with identifier")
	}
	return err
}
