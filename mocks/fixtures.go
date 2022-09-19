package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/entity"
)


type fixtureMap struct {
	UserCreate1		*entity.UserCreateUpdate
	UserDisplay1	*entity.UserDisplay
	User1Password	string
	UserCreate2		*entity.UserCreateUpdate
	UserDisplay2	*entity.UserDisplay
	User2Password	string
	UserList		[]*entity.UserDisplay
	Server			*gin.Engine
}

func GenerateFixture() (f fixtureMap) {

	f.UserDisplay1 = &entity.UserDisplay{
		ID: "783ed845-387f-4c50-9a64-fef701b4dbb8",
		Email: "mike@mail.com",
		Username: "mike",
		FirstName: "Mike",
		LastName: "Spensor",
		IsActive: false,
		IsStaff: false,
		IsSuperuser: false,
	}
	f.User1Password = "mike_password"

	f.UserCreate1 = &entity.UserCreateUpdate{
		UserDisplay: *f.UserDisplay1,
		Password: f.User1Password,
	}

	f.UserDisplay2 = &entity.UserDisplay{
		ID: "883ed845-387f-4c50-9a64-fef701b4dbb8",
		Email: "john@mail.com",
		Username: "John",
		FirstName: "Alfred",
		LastName: "Goldman",
		IsActive: false,
		IsStaff: false,
		IsSuperuser: false,
	}
	f.User2Password = "john_password"

	f.UserCreate2 = &entity.UserCreateUpdate{
		UserDisplay: *f.UserDisplay1,
		Password: f.User2Password,
	}

	f.UserList = append(f.UserList, f.UserDisplay1)
	f.UserList = append(f.UserList, f.UserDisplay2)

	gin.SetMode(gin.TestMode)
	f.Server = gin.Default()

	return
}