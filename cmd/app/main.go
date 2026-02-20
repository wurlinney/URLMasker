package main

import (
	"log"
	"urlmasker/internal/app"
)

const Prefix = "http://"

func main() {
	producer := app.NewProducer("input.txt")
	presenter := app.NewPresenter("output.txt")
	service := app.NewService(producer, presenter)
	if err := service.Run(); err != nil {
		log.Fatalf("Ошибка в работе сервиса: %v", err)
	}
}
