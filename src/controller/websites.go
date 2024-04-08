package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/izmno/hosting-manager/src/core"
	"github.com/izmno/hosting-manager/src/models"
)

type WebsiteController struct {
	core.Controller
}

func (ctrl *WebsiteController) RegisterRoutes(router gin.IRouter) {
	router.GET("/", ctrl.Websites)
}

func (ctrl *WebsiteController) Websites(c *gin.Context) {
	websites := []*models.Website{
		{Name: "John Doe's Webpage", Status: "Live"},
		{Name: "JohnDoesIt.example.com", Status: "In maintenance"},
	}

	ctrl.Render(c, "websites.html", gin.H{
		"title":    "Websites",
		"websites": websites,
	})
}
