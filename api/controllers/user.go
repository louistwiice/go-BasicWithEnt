package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/domain"
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/utils"
)

type controller struct {
	service domain.UserService
}

func NewUserController(svc domain.UserService) *controller {
	return &controller{
		service: svc,
	}
}

func (c *controller) listUsers(ctx *gin.Context) {
	users, err := c.service.List()
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "successful", users)
}

func (c *controller) getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	
	user, _, err := c.service.GetByID(id)
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusFound, "successful", user)
}

func (c *controller) updateUser(ctx *gin.Context) {
	var data *entity.UserCreateUpdate
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, _, err := c.service.GetByID(id)
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrNotFound.Error(), nil)
		return
	}

	data = entity.ValidateUpdate(data, user)

	err = c.service.UpdateUser(data)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response := entity.UserDisplayFormater(data)
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusAccepted, "successful", response)
}


// Update user password
func (c *controller) updatePassword(ctx *gin.Context) {
	var data entity.ChangePassword
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, password, err := c.service.GetByID(id)
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrNotFound.Error(), nil)
		return
	}

	err = utils.CheckHashedString(data.OldPassword, password)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, "old password does not match", err.Error())
		return
	}

	udapte_user := &entity.UserCreateUpdate{
		UserDisplay: *user,
		Password: data.NewPassword,
	}
	
	err = c.service.UpdatePassword(udapte_user)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "successful", "Password reset successfully")
}

/*
**
**
*/

func (c *controller) MakeUserHandlers(app *gin.RouterGroup) {
	app.GET("", c.listUsers)
	app.GET(":id", c.getUser)
	app.PUT(":id", c.updateUser)
	app.POST(":id/reset_password", c.updatePassword)
}