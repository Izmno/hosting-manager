package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var _core_TemlateManager *TemplateManager

func SetTemplateManager(tm *TemplateManager) {
	_core_TemlateManager = tm
}

func GetTemplateManager() *TemplateManager {
	if _core_TemlateManager == nil {
		panic("TemplateManager is not set")
	}
	return _core_TemlateManager
}

type Controller struct{}

func (c *Controller) Render(ctx *gin.Context, template string, data any) {
	c.RenderStatus(ctx, http.StatusOK, template, data)
}

func (c *Controller) RenderStatus(ctx *gin.Context, code int, template string, data any) {
	r := render.HTML{
		Template: GetTemplateManager().Get(template),
		Name:     template,
		Data:     data,
	}

	ctx.Render(code, r)
}
