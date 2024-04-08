package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/izmno/hosting-manager/src/core"
)

type SupportController struct {
	core.Controller
}

func (ctrl *SupportController) RegisterRoutes(router gin.IRouter) {
	router.GET("/", ctrl.Support)
}

func (ctrl *SupportController) Support(c *gin.Context) {
	ctrl.Render(c, "support.html", gin.H{
		"title": "Support",
	})
}
