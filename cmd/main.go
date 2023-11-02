package main

import (
	_ "github.com/lib/pq"
	web "gitlab.com/web1"
	"gitlab.com/web1/internal/config"
	"gitlab.com/web1/internal/handler"
	"gitlab.com/web1/internal/logger"
	"gitlab.com/web1/internal/repositories"
	"gitlab.com/web1/internal/services"
	"log"
)

func main() {
	lg, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer lg.Sync()

	cfg, err := config.LoadConfigFromYaml()
	if err != nil {
		lg.Fatal(err.Error())
	}

	db, err := repositories.NewPostgresDB(cfg)
	if err != nil {
		lg.Fatal(err.Error())
	}
	defer db.Close()

	repo := repositories.NewRepository(db)
	server := services.NewService(repo)
	handlers := handler.NewHandler(server, lg)

	srv := new(web.Server)
	err = srv.Run(cfg.Port, handlers.InitRoutes())
	if err != nil {
		lg.Fatal(err.Error())
	}

}
