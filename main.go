package main

import (
	"chibihokage/simple-api/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user/add", user.Add)
	r.GET("/user/:phone", user.Get)
	r.GET("/users", user.List)
	r.Run() // listen and serve on 0.0.0.0:8080
}
