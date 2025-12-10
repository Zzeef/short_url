package main

import (
	"short_url/internal/config"
	"short_url/internal/link"
	"short_url/internal/storage"
)

func main() {
	cfg := config.LoadConfig()
	mongoStore := storage.NewMongo(cfg.DB.Host, cfg.DB.Name)
	defer mongoStore.Close()

	linkRepo := link.NewRepository(mongoStore)
	linkService := link.NewService(linkRepo)
	linkHandler := link.NewHandler(linkService)
	linkRouter := link.NewRouter(linkHandler)

	linkRouter.Run(cfg.Port)
}
