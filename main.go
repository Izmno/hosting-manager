package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izmno/hosting-manager/src/controller"
	"github.com/izmno/hosting-manager/src/core"
	"github.com/koding/multiconfig"
)

var db = make(map[string]string)

func setupRouter(c *AppConfig) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.
		Use(gin.Logger()).
		Use(gin.Recovery()).
		StaticFile("/favicon.ico", "./public/favicon.ico").
		StaticFS("/assets", http.Dir("public"))

	tm := core.NewTemplateManager("src/templates")
	if err := tm.Load(); err != nil {
		panic(err)
	}

	core.SetTemplateManager(tm)

	cg := core.NewControllerGroup().
		AddController(
			new(controller.HomeController),
			core.WithRoutePrefix("/"),
		).
		AddController(
			new(controller.SupportController),
			core.WithRoutePrefix("/support"),
		).
		AddController(
			new(controller.AccountController),
			core.WithRoutePrefix("/account"),
		).
		AddController(
			new(controller.WebsiteController),
			core.WithRoutePrefix("/websites"),
		)

	cg.RegisterRoutes(r)

	return r
}

type AppConfig struct {
	ConfigFile string `flagUsage:"Additional config file to load"`
	Debug      bool   `flagUsage:"Enable debug mode" default:"false"`
	Port       int    `flagUsage:"Server should listen to this port" default:"8080"`
}

func main() {
	c := new(AppConfig)

	m := multiconfig.MultiLoader(
		&multiconfig.TagLoader{},
		&multiconfig.EnvironmentLoader{},
		&multiconfig.FlagLoader{},
	)

	if err := m.Load(c); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			// Help requested, should have been printed by now.
			return
		}

		panic(err)
	}

	if c.ConfigFile != "" {
		m := multiconfig.NewWithPath(c.ConfigFile)
		if err := m.Load(c); err != nil {
			panic(err)
		}
	}

	r := setupRouter(c)
	r.Run(fmt.Sprintf(":%d", c.Port))
}
