package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/izmno/hosting-manager/src/core"
)

type AccountController struct {
	core.Controller
}

func (ctrl *AccountController) RegisterRoutes(router gin.IRouter) {
	router.GET("/", ctrl.EditAccount)
}

func (ctrl *AccountController) EditAccount(c *gin.Context) {
	ctrl.Render(c, "account.html", gin.H{
		"title": "Account",
	})
}
