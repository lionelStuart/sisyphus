package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"sisyphus/common/app"
	"sisyphus/common/ecode"
	"sisyphus/common/utils"
	authService "sisyphus/service/auth"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

type AuthController struct {
}

// @Summary GetAuth
// @Tags Auth
// @Produce  json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func (c *AuthController) GetAuth(ctx *gin.Context) {
	var (
		ginX  = app.GinX{C: ctx}
		valid = validation.Validation{}
	)

	username := ctx.Query("username")
	password := ctx.Query("password")

	a := auth{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)
	if !ok {
		app.MarkErrors(valid.Errors)
		ginX.JSON(http.StatusBadRequest, ecode.INVALID_PARAMS, nil)
		return
	}

	authSvc := authService.Auth{Username: username, Password: password}
	exist, err := authSvc.Check()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	} else if !exist {
		ginX.JSON(http.StatusUnauthorized, ecode.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenToken(username, password)
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_AUTH_TOKEN, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, gin.H{
		"token": token,
	})

}
