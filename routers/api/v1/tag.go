package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"sisyphus/common/app"
	"sisyphus/common/ecode"
	tagService "sisyphus/service/tag"
)

type TagController struct {
}

// @Summary Get multiple article tags
// @Tags TagController
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags [get]
func (t *TagController) GetTags(ctx *gin.Context) {
	ginX := app.GinX{C: ctx}
	name := ctx.Query("name")
	state := -1

	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagSvc := tagService.Tag{
		Name:     name,
		State:    state,
		PageNum:  com.StrTo(ctx.Query("page")).MustInt(),
		PageSize: 4, // TODO
	}

	tags, err := tagSvc.GetAll()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	count, err := tagSvc.Count()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, gin.H{
		"lists": tags,
		"total": count,
	})

}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add article tag
// @Tags TagController
// @Produce  json
// @Param form body v1.AddTagForm true "AddTagForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags [post]
func (t *TagController) AddTag(ctx *gin.Context) {
	var (
		ginX = app.GinX{C: ctx}
		form AddTagForm
	)

	httpCode, errCode := app.BindAndValid(ctx, &form)
	if errCode != ecode.SUCCESS {
		ginX.JSON(httpCode, errCode, nil)
		return
	}

	tagSvc := tagService.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exist, err := tagSvc.ExistByName()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if exist {
		ginX.JSON(http.StatusOK, ecode.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagSvc.Add()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, nil)

}

type EditTagForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update article tag
// @Tags TagController
// @Produce  json
// @Param id path int true "ID"
// @Param form body v1.EditTagForm true "EditTagForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags/{id} [put]
func (t *TagController) EditTag(ctx *gin.Context) {
	var (
		ginX = app.GinX{C: ctx}
		form = EditTagForm{ID: com.StrTo(ctx.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(ctx, &form)
	if errCode != ecode.SUCCESS {
		ginX.JSON(httpCode, errCode, nil)
		return
	}

	tagSvc := tagService.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := tagSvc.ExistByID()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		ginX.JSON(http.StatusOK, ecode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagSvc.Edit()
	if err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, nil)
}

// @Summary Delete article tag
// @Tags TagController
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tags/{id} [delete]
func (t *TagController) DeleteTag(ctx *gin.Context) {
	var (
		ginX  = app.GinX{C: ctx}
		valid = validation.Validation{}
		id    = com.StrTo(ctx.Param("id")).MustInt()
	)
	valid.Min(id, 1, "id").Message("ID requires large than 0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		ginX.JSON(http.StatusBadRequest, ecode.INVALID_PARAMS, nil)
		return
	}

	tagSvc := tagService.Tag{ID: id}
	exists, err := tagSvc.ExistByID()
	if err != nil {
		ginX.JSON(http.StatusBadRequest, ecode.ERROR_EXIST_TAG_FAIL, nil)
		return
	} else if !exists {
		ginX.JSON(http.StatusOK, ecode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := tagSvc.Delete(); err != nil {
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, nil)
}
