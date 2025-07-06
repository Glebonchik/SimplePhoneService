package application

import (
	"danek.com/telephone/config"
	httpapp "danek.com/telephone/internal/application/http"
	"danek.com/telephone/internal/controller"
	"danek.com/telephone/internal/db"
	"danek.com/telephone/internal/domain"
	"danek.com/telephone/internal/usecase"
)

type Application struct {
	HTTPSrv *httpapp.App
}

func New(cfg *config.Config) (*Application, error) {
	post, err := db.NewDBConn(cfg)
	if err != nil {
		return nil, domain.FormatErr("New DB conn", err)
	}

	gormRepo := db.NewRepostitory(post)
	userUseCase := usecase.NewUserUseCase(gormRepo)
	userHandler := controller.NewUserHandler(userUseCase)

	httpApp := httpapp.NewApp(cfg, userHandler)
	return &Application{HTTPSrv: httpApp}, nil
}
