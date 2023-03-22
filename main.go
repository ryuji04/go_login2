package main

import (
	"fmt"
	"master/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("template/*html")
	fmt.Println("first")
	router.GET("/login", controller.Login)
	router.POST("/confirm", controller.Confirm)
	router.GET("/top", controller.Top)
	router.GET("/", controller.Index)
	fmt.Println("最初のgetを通過")
	router.GET("/create", controller.Create)
	router.POST("/insert", controller.Insert)
	router.GET("/getall", controller.GetAll)
	fmt.Println("main detail前")
	router.GET("/detail/:id", controller.Detail)
	router.POST("/edit", controller.Edit)
	router.POST("/delete", controller.Delete)
	fmt.Println("8080前")
	router.Run(":8080")
	fmt.Println("8080後")
}
