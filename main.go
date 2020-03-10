package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sisyphus/common/redis"
	"sisyphus/common/setting"
	"sisyphus/models"
	"sisyphus/routers"
)

func NewServer(handler *gin.Engine, conf *setting.Server) *http.Server {
	maxHeaderBytes := 1 << 20

	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", "0.0.0.0", conf.HttpPort),
		Handler:        handler,
		MaxHeaderBytes: maxHeaderBytes,
	}
}

//func buildContainer() *dig.Container{
//	container := dig.New()
//
//	// container.Provide(routers.InitRouters)
//	container.Provide(NewServer)
//	v1.Inject(container)
//
//	return container
//}

func init() {
	path := `conf/app.ini`
	setting.Setup(path)
	models.Setup()
	redis.SetUp()

}

// @title 测试 gin API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1/
func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
	gin.SetMode("debug")
	e := gin.New()
	// c := buildContainer()
	routers.InitRouters(e)

	s := NewServer(e, setting.GetServerConf())

	s.ListenAndServe()

	// handlers :=routers.InitRouters()

}
