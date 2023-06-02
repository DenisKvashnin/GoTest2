package main

import (
	OrderApi "TestTask/internal/app/api/order"
	OrderProvider "TestTask/internal/app/providers/order"
	Database "TestTask/internal/app/repositories/base"
	OrderRepository "TestTask/internal/app/repositories/order"
	OrderService "TestTask/internal/app/services/order"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/labstack/echo/v4"
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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			_, err := orderService.GetAndSaveOrder()

			if err != nil {
				log.Println("Error while receiving order: ", err)
			}
		}
	}()

	router := mux.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		wg.Add(1)
		log.Fatal(server.ListenAndServe())
	}()
	orderApi := OrderApi.New(orderService)

	router.HandleFunc("/api/v1/order", orderApi.GetOrder).Methods("POST")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	wg.Wait()

}
