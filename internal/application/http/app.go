package httpapp

import (
	"context"
	"log"
	"net/http"
	"time"

	"danek.com/telephone/config"
	"danek.com/telephone/internal/controller"
	"danek.com/telephone/internal/domain"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config      *config.Config
	Server      *http.Server
	UserHandler *controller.UserHandler
}

func NewApp(cfg *config.Config, uh *controller.UserHandler) *App {
	return &App{Config: cfg, UserHandler: uh}
}

func (app *App) Run() error {
	router := gin.Default()
	SetupRoutes(router, app.Config, app.UserHandler)
	adr := app.Config.HTTPServer.Port
	srv := &http.Server{
		Addr:    adr,
		Handler: router,
	}
	app.Server = srv

	if err := app.Server.ListenAndServe(); err != nil {
		return domain.FormatErr("Startin Server", err)
	}
	return nil
}

func SetupRoutes(router *gin.Engine, cfg *config.Config, uh *controller.UserHandler) {
	router.POST("/add-user", uh.Register)
	router.POST("/clear-user", uh.ClearUser)
	router.PUT("/update-user", uh.UpdateUser)
}

func (app *App) StopApp() error {
	log.Println("Http server stopping ")
	if app.Server == nil {
		log.Println("Server isn't running")
	}

	ctx, cancelapp := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelapp()

	if err := app.Server.Shutdown(ctx); err != nil {
		return domain.FormatErr("Stopping server", err)
	}
	log.Println("Server shoutdowned gracefully")
	return nil
}
