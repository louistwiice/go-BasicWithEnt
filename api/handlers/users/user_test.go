package handler_users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/mocks"
	"github.com/louistwiice/go/basicwithent/mocks/user"
	"github.com/louistwiice/go/basicwithent/utils"
	"github.com/stretchr/testify/assert"
)

func Test_getUser(t *testing.T) {
	t.Run("No error when ID exist", func(t *testing.T) {
		fixture := mocks.GenerateFixture()
		router := fixture.Server
		u := fixture.UserDisplay1
		u_password := fixture.User1Password
		service := user.MockUserService{}
		service.On("GetByID", u.ID).Return(u, u_password, nil)

		controller := NewUserController(&service)
		router.GET("/:id", controller.getUser)

		rr := httptest.NewRecorder()
		expected_body_resp, err := json.Marshal(utils.Response{
			Code: http.StatusFound,
			Message: "successful",
			Data: u,
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", u.ID), nil)
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)

		assert.Equal(t, expected_body_resp, rr.Body.Bytes())
	})

	t.Run("Error returned when service return an error", func(t *testing.T) {
		fixture := mocks.GenerateFixture()
		router := fixture.Server
		u := fixture.UserDisplay1
		service := user.MockUserService{}
		service.On("GetByID", u.ID).Return(u, "", entity.ErrNotFound)

		controller := NewUserController(&service)
		router.GET("/:id", controller.getUser)

		rr := httptest.NewRecorder()
		expected_body_resp, err := json.Marshal(utils.Response{
			Code: http.StatusBadRequest,
			Message: entity.ErrNotFound.Error(),
			Data: nil,
		})
		assert.NoError(t, err)
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", u.ID), nil)
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)

		assert.Equal(t, expected_body_resp, rr.Body.Bytes())
	})
}
