package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashString(t *testing.T) {
	plain_password := "mikepassword"

	result, err := HashString(plain_password)
	assert.Nil(t, err)
	assert.NotEqual(t, result, plain_password)
}

func Test_CheckHashedString(t *testing.T) {
	plain_password := "mikepassword"
	result, _ := HashString(plain_password)

	err := CheckHashedString(plain_password, result)
	assert.Nil(t, err)
}
