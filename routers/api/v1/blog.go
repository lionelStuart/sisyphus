package v1

import "github.com/gin-gonic/gin"

type BlogInfo struct {
	Title string `json:"title"`
}

type BlogController struct {
}

func (b *BlogController) GetBlog(context *gin.Context) {
	bi := &BlogInfo{
		Title: "test with sisyphus",
	}

	context.JSON(200, bi)
}
