package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/izmno/hosting-manager/src/core"
)

type HomeController struct {
	core.Controller
}

func (ctrl *HomeController) RegisterRoutes(router gin.IRouter) {
	router.GET("/", ctrl.Home)
}

func (ctrl *HomeController) Home(c *gin.Context) {
	ctrl.Render(c, "home.html", gin.H{
		"title": "Home",
	})
}
