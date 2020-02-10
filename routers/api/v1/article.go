package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sisyphus/models"
	"strconv"
)

type ArticleController struct {
}

func (a *ArticleController) GetArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
	if a, err := models.GetArticle(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, a)
	}

}

func (a *ArticleController) GetArticles(ctx gin.Context) {

}
