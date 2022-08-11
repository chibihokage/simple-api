package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {
	var req Profile
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	users = append(users, req)

	c.Status(http.StatusCreated)
}

func Get(c *gin.Context) {
	var user Profile
	phone := c.Param("phone")

	for _, u := range users {
		if u.Phone == phone {
			user = u
			break
		}
	}

	c.JSON(200, user)
}

func List(c *gin.Context) {
	c.JSON(200, users)
}
