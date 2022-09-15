package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateUpdate_With_Same_Name(t *testing.T) {
	u := &UserDisplay{
		ID: "2721d5df-59eb-4d29-bd04-79ef3d346c8e",
		Email: "mike@mail.com",
		FirstName: "Mike",
		LastName: "Spenser",
	}

	user_to_update := &UserCreateUpdate{
		UserDisplay: UserDisplay{
			ID: "2721d5df-59eb-4d29-bd04-79ef3d346c8e",
			Email: "",
			FirstName: "Mike",
			LastName: "Spenser",
		},
		Password: "my_password",
	}

	result := ValidateUpdate(user_to_update, u)
	assert.NotEqual(t, result.Email, "")
	assert.Equal(t, result.Email, u.Email) // new email should be mike@mail.com

	user_to_update.Email = "newmail@gmail.com"

	result = ValidateUpdate(user_to_update, u)
	assert.Equal(t, result.Email, "newmail@gmail.com") // new email should be newmail@mail.com
}

func Test_ValidateUpdate_With_Different_Name(t *testing.T) {
	u := &UserDisplay{
		ID: "2721d5df-59eb-4d29-bd04-79ef3d346c8e",
		Email: "mike@mail.com",
		FirstName: "Mike",
		LastName: "Spenser",
	}

	user_to_update := &UserCreateUpdate{
		UserDisplay: UserDisplay{
			ID: "2721d5df-59eb-4d29-bd04-79ef3d346c8e",
			Email: "mike@mail.com",
			FirstName: "John",
			LastName: "Spenser",
		},
		Password: "my_password",
	}

	result := ValidateUpdate(user_to_update, u)
	assert.NotEqual(t, result.FirstName, u.FirstName) 
	assert.Equal(t, result.FirstName, "John")

	user_to_update.LastName = "Lennon"
	result = ValidateUpdate(user_to_update, u)
	assert.NotEqual(t, result.LastName, u.LastName) 
	assert.Equal(t, result.LastName, "Lennon")
}