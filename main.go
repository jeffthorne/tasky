package main

import (
	"net/http"
	controller "github.com/jeffthorne/tasky/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func main() {
	godotenv.Overload()
	
	router := gin.Default()
	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", index)
	router.GET("/todos/:userid", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodo)
	router.POST("/todo/:userid", controller.AddTodo)
	router.DELETE("/todo/:userid/:id", controller.DeleteTodo)
	router.DELETE("/todos/:userid", controller.ClearAll)
	router.PUT("/todo", controller.UpdateTodo)


	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/todo", controller.Todo)

	router.Run(":8080" )

}
