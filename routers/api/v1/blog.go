package v1

import "github.com/gin-gonic/gin"

type BlogInfo struct {
	Title string `json:"title"`
}

type BlogController struct {
}

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /api/v1/record/{some_id} [get]
func (b *BlogController) GetBlog(context *gin.Context) {
	bi := &BlogInfo{
		Title: "test with sisyphus",
	}

	context.JSON(200, bi)
}
