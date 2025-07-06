package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"danek.com/telephone/config"
	"danek.com/telephone/internal/application"
	"danek.com/telephone/internal/domain"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		domain.FormatErr("Loading config", err)
		return
	}
	log.Println("Loaded config file")

	app, err := application.New(cfg)
	if err != nil {
		domain.FormatErr("Loading app", err)
		return
	}
	log.Println("App started")
	go app.HTTPSrv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sign := <-stop
	log.Println("stopping application", slog.String("signal", sign.String()))

	app.HTTPSrv.StopApp()

	log.Println("application stopped")
}
