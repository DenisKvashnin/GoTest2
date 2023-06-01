package main

import (
	OrderProvider "TestTask/internal/app/providers/order"
	Database "TestTask/internal/app/repositories/base"
	OrderRepository "TestTask/internal/app/repositories/order"
	OrderService "TestTask/internal/app/services/order"
	"log"
	"time"
)

func main() {

	var (
		host     = "localhost"
		port     = "5432"
		user     = "myUser"
		password = "myPassword"
		dbName   = "myDb"
		url      = "http://localhost:8081/"
	)

	database := Database.New(host, port, user, password, dbName)
	defer func() {
		err := database.CloseDb()
		if err != nil {
			log.Println(err)
		}
	}()

	orderProvider := OrderProvider.New(url)
	orderRepository := OrderRepository.New(database)
	orderService := OrderService.New(orderRepository, orderProvider)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		_, err := orderService.GetAndSaveOrder()

		if err != nil {
			log.Fatalf("Error while receiving order: %s", err)
		}
	}
}
