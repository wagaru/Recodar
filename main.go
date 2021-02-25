package main

import (
	"fmt"
	"log"

	"github.com/wagaru/Recodar-backend/internal/config"
	"github.com/wagaru/Recodar-backend/internal/delivery/http"
	"github.com/wagaru/Recodar-backend/internal/repository"
	"github.com/wagaru/Recodar-backend/internal/usecase"
)

func main() {
	log.Printf("Init...")

	// Load config
	config, err := config.LoadConfig("./conf", "app", "env")
	if err != nil {
		fmt.Printf("Load config failed: %v", err)
	}
	log.Println("Load config completed.")

	// Init MongoDB
	repo, err := repository.NewMongoRepo(config)
	if err != nil {
		fmt.Printf("Init MongoDB failed: %v", err)
	}
	log.Println("Init Mongo DB completed.")

	// Init usecase
	usecase := usecase.NewUsecase(repo)
	log.Println("Init usecase completed.")

	// Init delivery
	delivery := http.NewHttpDelivery(usecase)
	log.Println("Init delivery completed.")

	delivery.Run(config.ServerPort)
}
