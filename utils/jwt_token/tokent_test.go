package jwttoken

import (
	//"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateToken(t *testing.T) {
	user_id := "8fe62a1c-e44e-47ce-b719-b4eef7f271f2"
	result, err := GenerateToken(user_id)
	assert.Nil(t, err)
	assert.NotNil(t, result)

	/*for key, _ := range result {
		assert.Equal(t, strings.HasSuffix(key, "_token"), true)
	}*/
}