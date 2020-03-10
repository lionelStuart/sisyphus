package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"net/http"
	"path"
	"sisyphus/common/app"
	"sisyphus/common/ecode"
	fileUtils "sisyphus/common/files"
)

type FileController struct {
}

// @Summary UploadFile
// @Tags FileSys
// @Produce  json
// @Param image formData file  true "image"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /upload [post]
func (c *FileController) Upload(ctx *gin.Context) {
	ginX := app.GinX{C: ctx}
	file, image, err := ctx.Request.FormFile("image")
	if err != nil {
		log.Warn(err)
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR, nil)
		return
	}
	if image == nil {
		ginX.JSON(http.StatusBadRequest, ecode.INVALID_PARAMS, nil)
		return
	}

	imageName := fileUtils.GenImageName(image.Filename)
	fullPath := fileUtils.GetImageFullPath()
	savePath := fileUtils.GetImagePath()
	src := path.Join(fullPath, imageName)

	if !fileUtils.CheckImageExt(imageName) || !fileUtils.CheckImageSize(file) {
		ginX.JSON(http.StatusBadRequest, ecode.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	err = fileUtils.CheckImage(fullPath)
	if err != nil {
		log.Warn(err)
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := ctx.SaveUploadedFile(image, src); err != nil {
		log.Warn(err)
		ginX.JSON(http.StatusInternalServerError, ecode.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	ginX.JSON(http.StatusOK, ecode.SUCCESS, gin.H{
		"image_url":      fileUtils.GetImageFullUrl(imageName),
		"image_save_url": path.Join(savePath, imageName),
	})
}

func (c *FileController) Download(ctx *gin.Context) {

}
