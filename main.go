package main

import (
	orderProvider "TestTask/internal/app/providers/order"
	orderRepo "TestTask/internal/app/repositories/order"
	"TestTask/internal/app/services/order"
	_ "github.com/lib/pq"
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
	orderRepository := orderRepo.New(host, port, user, password, dbName)
	orderRepository.InitDb()

	orderService := order.New(orderRepository, orderProvider.New(url))

	for {
		_, err := orderService.GetAndSaveOrder()

		if err != nil {
			log.Fatalf("Error while receiving order: %s", err)
		}

		time.Sleep(1 * time.Second)
	}
}
