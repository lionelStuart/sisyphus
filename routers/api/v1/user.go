package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sisyphus/common/app"
	"sisyphus/common/ecode"
	"sisyphus/service/user-svc/handler"
)

type UserController struct {
}

// @Summary Get user-svc profile by id
// @Tags UserController
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/{id} [get]
func (c *UserController) GetUserProfile(ctx *gin.Context) {
	var (
		ginX = app.GinX{C: ctx}
		id   string
	)
	id = ctx.Param("id")
	if len(id) == 0 {
		ginX.JSON(http.StatusBadRequest, ecode.INVALID_PARAMS, nil)
		return
	}

	userSvc := handler.User{ID: id}
	profile, err := userSvc.GetProfile()

	if err != nil {
		ginX.JSON(http.StatusBadRequest, ecode.ERROR_NOT_EXIST_USER, nil)
		return
	}
	ginX.JSON(http.StatusOK, ecode.SUCCESS, profile)
}
