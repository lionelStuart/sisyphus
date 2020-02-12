package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sisyphus/models"
	"sisyphus/routers"
)

func NewServer(handler *gin.Engine) *http.Server {
	maxHeaderBytes := 1 << 20

	return &http.Server{
		Addr:           "127.0.0.1:8079",
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

// @title 测试
// @version 0.0.1
// @description  测试
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
	models.InitModel()
	routers.InitRouters(e)

	s := NewServer(e)

	s.ListenAndServe()

	// handlers :=routers.InitRouters()

}
