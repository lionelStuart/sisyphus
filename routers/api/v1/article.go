package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sisyphus/common/app"
	"sisyphus/common/ecode"
	"sisyphus/models"
	"strconv"
)

type ArticleController struct {
}

// @Summary Get a single article By id
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} models.Article
// @Failure 500
// @Router /articles/{id} [get]
func (a *ArticleController) GetArticle(ctx *gin.Context) {
	fmt.Println("do get")
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

func (a *ArticleController) GetArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get articles")

}

type AddArticleForm struct {
	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy     string `form:"created_by" valid:"Required;MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Get a single article By id
// @Produce  json
// @Param tag_id body int true "TagID"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param created_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} string"ok"
// @Failure 500 {string} string"fail"
// @Router /articles [post]
func (a *ArticleController) AddArticles(ctx *gin.Context) {
	fmt.Println("testset")
	var form AddArticleForm
	httpCode, errCode := app.BindAndValid(ctx, &form)
	if errCode != ecode.SUCCESS {
		ctx.JSON(httpCode, gin.H{
			"err": errCode,
		})
	}

	article := gin.H{
		"tag_id":          form.TagID,
		"title":           form.Title,
		"desc":            form.Desc,
		"content":         form.Content,
		"created_by":      form.CreatedBy,
		"cover_image_url": form.CoverImageUrl,
		"state":           form.State,
	}

	if err := models.AddArticle(article); err != nil {
		ctx.JSON(httpCode, gin.H{
			"err": ecode.ERROR_ADD_ARTICLE_FAIL,
		})
	}

	ctx.JSON(http.StatusOK, ecode.SUCCESS)
}

func (a *ArticleController) EditArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "edit articles")
}

func (a *ArticleController) DeleteArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "delete articles")
}
