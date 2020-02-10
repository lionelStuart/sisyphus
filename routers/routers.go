package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "sisyphus/routers/api/v1"
	"time"
)

var (
	blogCtrl    v1.BlogController
	articleCtrl v1.ArticleController
)

func InitRouters(engine *gin.Engine) error {
	//r := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	engine.GET("/ping", pong)
	apiv1 := engine.Group("/v1")
	//apiv1.Use()
	blog := apiv1.Group("blog")
	{
		blog.GET("/", blogCtrl.GetBlog)

	}

	articles := apiv1.Group("articles")
	{
		articles.GET("/:id", articleCtrl.GetArticle)
	}

	return nil
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}
