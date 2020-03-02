package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"sisyphus/common/app"
	"sisyphus/common/ecode"
	articleService "sisyphus/service/article"
	tagService "sisyphus/service/tag"
	"strconv"
)

type ArticleController struct {
}

// @Summary Get a single article By id
// @Tags ArticleController
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /articles/{id} [get]
func (a *ArticleController) GetArticle(ctx *gin.Context) {
	ginX := app.GinX{C: ctx}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ginX.JSON(http.StatusBadRequest, ecode.INVALID_PARAMS, nil)
		return
	}

	articleSvc := articleService.Article{ID: id}

	if exists, err := articleSvc.ExistByID(); err != nil {
		ginX.JSON(http.StatusBadRequest, ecode.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	} else if !exists {
		ginX.JSON(http.StatusOK, ecode.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleSvc.Get()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	ginX.JSON(http.StatusOK, ecode.SUCCESS, article)
}

// @Summary Get multiple articles
// @Tags ArticleController
// @Produce  json
// @Param tag_id query int false "TagID"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /articles [get]
func (a *ArticleController) GetArticles(ctx *gin.Context) {
	var (
		ginX  = app.GinX{C: ctx}
		valid = validation.Validation{}
		state = -1
		tagId = -1
	)

	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}
	if arg := ctx.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		ginX.JSON(http.StatusBadRequest, ecode.INVALID_PARAMS, nil)
	}

	articleSvc := articleService.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  com.StrTo(ctx.Query("page")).MustInt(),
		PageSize: 4, //TODO
	}

	total, err := articleSvc.Count()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := articleSvc.GetAll()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_GET_ARTICLE_FAIL, nil)
	}

	data := map[string]interface{}{
		"lists": articles,
		"total": total,
	}
	ginX.JSON(http.StatusOK, ecode.SUCCESS, data)
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

// @Summary Add article
// @Tags ArticleController
// @Accept json
// @Produce  json
// @Param   form body v1.AddArticleForm true "Add articles form"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /articles [post]
func (a *ArticleController) AddArticles(ctx *gin.Context) {
	ginX := app.GinX{C: ctx}

	var form AddArticleForm
	httpCode, errCode := app.BindAndValid(ctx, &form)
	if errCode != ecode.SUCCESS {
		ginX.JSON(httpCode, errCode, nil)
		return
	}

	tagSvc := tagService.Tag{ID: form.TagID}
	exists, err := tagSvc.ExistByID()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if !exists {
		ginX.JSON(http.StatusOK, ecode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleSvc := articleService.Article{
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		State:         form.State,
	}

	if err := articleSvc.Add(); err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}
	ginX.JSON(http.StatusOK, ecode.SUCCESS, nil)
}

type EditArticleForm struct {
	ID            int    `form:"id" valid:"Required;Min(1)"`
	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
	ModifiedBy    string `form:"modified_by" valid:"Required;MaxSize(100)"`
	CoverImageUrl string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update article
// @Tags ArticleController
// @Produce  json
// @Param id path int true "ID"
// @Param from body v1.EditArticleForm true "EditArticleForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /articles/{id} [put]
func (a *ArticleController) EditArticle(ctx *gin.Context) {
	ginX := app.GinX{C: ctx}
	form := EditArticleForm{ID: com.StrTo(ctx.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(ctx, &form)
	if errCode != ecode.SUCCESS {
		ginX.JSON(httpCode, errCode, nil)
		return
	}

	articleSvc := articleService.Article{
		ID:            form.ID,
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		CoverImageUrl: form.CoverImageUrl,
		ModifiedBy:    form.ModifiedBy,
		State:         form.State,
	}

	exists, err := articleSvc.ExistByID()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagSvc := tagService.Tag{ID: form.TagID}
	exists, err = tagSvc.ExistByID()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		ginX.JSON(http.StatusOK, ecode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleSvc.Edit()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, nil)
}

// @Summary Delete article
// @Tags ArticleController
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /articles/{id} [delete]
func (a *ArticleController) DeleteArticle(ctx *gin.Context) {
	ginX := app.GinX{C: ctx}

	valid := validation.Validation{}
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid.Min(id, 1, "id").Message("ID large then 0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		ginX.JSON(http.StatusOK, ecode.INVALID_PARAMS, nil)
		return
	}

	articleSvc := articleService.Article{ID: id}
	exist, err := articleSvc.ExistByID()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	} else if !exist {
		ginX.JSON(http.StatusOK, ecode.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleSvc.Delete()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, nil)
}
