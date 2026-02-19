package main

import (
	"log"
	"urlmasker/internal/app"
)

const Prefix = "http://"

func main() {
	//link := "Hello, its my page //string//: http://localhost123.com See you http://localhost123.com "
	producer := &app.FileProducer{Filename: "input.txt"}
	presenter := &app.FilePresenter{Filename: "output.txt"}
	service := app.NewService(producer, presenter)
	if err := service.Run(); err != nil {
		log.Fatalf("Ошибка в работе сервиса: %v", err)
	}
}
