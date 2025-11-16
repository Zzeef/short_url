package main

import (
	"short_url/internal/link"
	"short_url/internal/storage"
)

func main() {
	mongoStore := storage.NewMongo("mongodb://localhost:27017", "short_url")
	defer mongoStore.Close()

	linkRepo := link.NewRepository(mongoStore)
	linkService := link.NewService(linkRepo)
	linkHandler := link.NewHandler(linkService)
	linkRouter := link.NewRouter(linkHandler)

	linkRouter.Run()
}
