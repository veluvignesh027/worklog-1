package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/veluvignesh027/worklog/Database"
)

func InitHandelers(router *gin.Engine) error {
	router.GET("/", indexPageHandler)
	router.GET("/auth-user", authUserHandler)
	router.GET("/home", homePageHandler)

	return nil
}

func indexPageHandler(c *gin.Context) {
	c.File("Templates/GET_IndexPage.html")
}

func homePageHandler(c *gin.Context) {
	c.File("Templates/templates/GET_HomePage.html")
}

func authUserHandler(c *gin.Context) {
	fname := c.DefaultQuery("username", "unknown")
	pass := c.DefaultQuery("password", "none")

	for _, user := range database.Users {
		if user.Name == fname {
			if user.Pass == pass {
				c.String(http.StatusAccepted, "Sucess!")
			} else {
				c.String(http.StatusBadGateway, "Password wrong!")
			}
		} else {
			c.String(http.StatusBadGateway, "User not exists")
		}
	}

	c.String(http.StatusBadRequest, "Bad request")

}
