package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sisyphus/common/setting"
	_ "sisyphus/docs"
	v1 "sisyphus/routers/api/v1"
	"time"
)

var (
	blogCtrl    v1.BlogController
	articleCtrl v1.ArticleController
	tagCtrl     v1.TagController
	authCtrl    v1.AuthController
	fileCtrl    v1.FileController
	userCtrl    v1.UserController
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

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.GET("/ping", pong)

	apiv1 := engine.Group("/v1")

	auth := apiv1.Group("/auth")
	{
		auth.GET("/", authCtrl.GetAuth)
	}

	file := engine.Group("/file")
	{
		file.POST("/upload", fileCtrl.Upload)
		//file.StaticFS("/static",http.Dir(setting.GetAppConf().StaticRootPath))
		file.Static("/static", setting.GetAppConf().StaticRootPath)
	}

	//apiv1.Use()
	blog := apiv1.Group("blog")
	{
		blog.GET("/", blogCtrl.GetBlog)

	}

	articles := apiv1.Group("articles")
	{
		articles.GET("/", articleCtrl.GetArticles)
		articles.POST("/", articleCtrl.AddArticles)
		articles.GET("/:id", articleCtrl.GetArticle)
		articles.PUT("/:id", articleCtrl.EditArticle)
		articles.DELETE("/:id", articleCtrl.DeleteArticle)
	}

	tags := apiv1.Group("tags")
	{
		tags.GET("/", tagCtrl.GetTags)
		tags.POST("/", tagCtrl.AddTag)
		tags.PUT("/:id", tagCtrl.EditTag)
		tags.DELETE("/:id", tagCtrl.DeleteTag)
	}

	users := apiv1.Group("users")
	{
		users.GET("/:id", userCtrl.GetUserProfile)
	}
	return nil
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}
