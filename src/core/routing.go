package core

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	RegisterRoutes(router gin.IRouter)
}

type ControllerGroup struct {
	controllers []*ControllerConfig
}

func NewControllerGroup(controllers ...*ControllerConfig) *ControllerGroup {
	return &ControllerGroup{
		controllers: controllers,
	}
}

func (cg *ControllerGroup) AddController(ctrl ControllerInterface, options ...ControllerOption) *ControllerGroup {
	return cg.AddControllerConfig(NewControllerConfig(ctrl, options...))
}

func (cg *ControllerGroup) AddControllerConfig(config *ControllerConfig) *ControllerGroup {
	if cg.controllers == nil {
		cg.controllers = make([]*ControllerConfig, 0)
	}

	cg.controllers = append(cg.controllers, config)
	return cg
}

func (cg *ControllerGroup) RegisterRoutes(router gin.IRouter) {
	for _, controller := range cg.controllers {
		controller.controller.RegisterRoutes(router.Group(controller.routePrefix, controller.middleWare...))
	}
}

type ControllerConfig struct {
	controller ControllerInterface

	routePrefix string
	middleWare  []gin.HandlerFunc
}

func NewControllerConfig(controller ControllerInterface, options ...ControllerOption) *ControllerConfig {
	config := &ControllerConfig{
		controller: controller,
	}

	for _, option := range options {
		option(config)
	}

	return config
}

type ControllerOption func(*ControllerConfig)

func WithRoutePrefix(prefix string) ControllerOption {
	return func(c *ControllerConfig) {
		c.routePrefix = prefix
	}
}

func WithMiddleWare(middleWare ...gin.HandlerFunc) ControllerOption {
	return func(c *ControllerConfig) {
		c.middleWare = middleWare
	}
}
