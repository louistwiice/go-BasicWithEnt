package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/configs"
	"github.com/louistwiice/go/basicwithent/domain"
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/utils"
	jwttoken "github.com/louistwiice/go/basicwithent/utils/jwt_token"
)

type authcontroller struct {
	service domain.AuthService
}

func NewAuthController(svc domain.AuthService) *authcontroller {
	return &authcontroller{
		service: svc,
	}
}

// Register/Create a new user or account
func (c *authcontroller) register(ctx *gin.Context) {
	var user entity.UserCreateUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := c.service.Create(&user)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "successful", user)
}

// Login is used to connect to the API
func (c *authcontroller) login(ctx *gin.Context) {
	conf := configs.LoadConfigEnv()
	var input entity.UserLogin

	if err := ctx.ShouldBindJSON(&input); err !=nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, hashed_password, err := c.service.SearchUser(input.Identifier)
	if err!= nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUserNotFound.Error(), nil)
		return
	}

	err = utils.CheckHashedString(input.Password, hashed_password)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tokens, err := jwttoken.GenerateToken(user.ID)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ctx.SetCookie("access_token", tokens["access_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", tokens["refresh_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", 1, "/", "localhost", false, false)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "Login successfully", gin.H{"token": tokens, "duration": conf.AccessTokenHourLifespan, "token_prefix": conf.TokenPrefix})
}

func (c *authcontroller) refreshToken(ctx *gin.Context) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	conf := configs.LoadConfigEnv()
	tokenReq := tokenReqBody{}

	if err := ctx.ShouldBindJSON(&tokenReq); err !=nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	claim, err := jwttoken.ExtractClaimsFromRefresh(tokenReq.RefreshToken)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, _, err := c.service.GetByID(claim["sub"].(string))
	if err != nil || user == nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, entity.ErrUserNotFound.Error(), nil)
		return
	}

	tokens, err := jwttoken.RefreshToken(ctx, tokenReq.RefreshToken)
	if err != nil {
		utils.ResponseJSON(ctx, http.StatusOK, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ctx.SetCookie("access_token", tokens["access_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", tokens["refresh_token"], 1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", 1, "/", "localhost", false, false)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "refresh successfully", gin.H{"token": tokens, "duration": conf.AccessTokenHourLifespan, "token_prefix": conf.TokenPrefix})
	
}

func (c *authcontroller) logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	utils.ResponseJSON(ctx, http.StatusOK, http.StatusOK, "logout successfully", nil)

}


/*
**
**
*/

func (c *authcontroller) MakeAuthHandlers(app *gin.RouterGroup) {
	app.POST("login", c.login)
	app.POST("register", c.register)
	app.POST("refresh", c.refreshToken)
	app.GET("logout", c.logout)
}