package controller

import (
	"fmt"
	"master/model"
	_ "master/model"
	"master/server"
	_ "master/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type User struct{
// 	Id int
// 	Name string
// 	Departure string
// }

// func NewUser()User{
// 	return User{}
// }

// func setRouter() gin.Default {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("template/*html")
// 	return router
// }

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func Confirm(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	user := server.NewUser()
	if !user.Confirm(id, name) {
		panic("ログイン失敗")
	}

	c.Redirect(http.StatusMovedPermanently, "/top")
}

func Top(c *gin.Context) {
	c.HTML(200, "top.html", gin.H{})
}

func Create(c *gin.Context) {
	// router := setRouter()
	// router.POST("/create", func(c *gin.Context) {
	// name := c.PostForm("name")
	// departure := c.PostForm("departure")
	// user := server.NewUser()
	// user.Create(name, departure)
	c.HTML(200, "create.html", gin.H{})
}

func Insert(c *gin.Context) {
	name := c.PostForm("name")
	departure := c.PostForm("departure")
	user := server.NewUser()
	user.Create(name, departure)
	fmt.Println("insert:", name, departure)
	c.Redirect(http.StatusMovedPermanently, "/getall")
}

func GetAll(c *gin.Context) {
	// router := setRouter()
	// router.GET("/getAll", func(c *gin.Context) {
	user := server.NewUser()
	users := user.GetAll()
	fmt.Println("getall:", users)
	c.HTML(http.StatusOK, "user_all.html", gin.H{"users": users})
}

func Detail(c *gin.Context) {
	fmt.Println("cont detail")
	u := server.NewUser()
	id := c.Param("id")
	var user model.User = u.FindById(id)
	fmt.Println("detail user:", user)
	c.HTML(http.StatusOK, "detail.html", gin.H{"user": user})
}

func Edit(c *gin.Context) {
	user := server.NewUser()
	id := c.PostForm("id")
	name := c.PostForm("name")
	departure := c.PostForm("departure")
	fmt.Println("cont Edit:", id, name, departure)
	user.Edit(id, name, departure)
	c.Redirect(http.StatusMovedPermanently, "/detail/"+id)
}

func Delete(c *gin.Context) {
	user := server.NewUser()
	id := c.PostForm("id")
	fmt.Println("cont Delete:", id)
	user.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/getall")
}
