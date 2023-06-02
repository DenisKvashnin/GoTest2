package order

import (
	"TestTask/internal/app/providers"
	OrderService "TestTask/internal/app/services/order"
	"encoding/json"
	"log"
	"net/http"
)

type OrderApi struct {
	orderService *OrderService.Service
}

func New(orderService *OrderService.Service) *OrderApi {
	return &OrderApi{
		orderService: orderService,
	}
}

func (o OrderApi) GetOrder(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IDs []int `json:"ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ids := request.IDs
	orders, err := o.orderService.GetOrderByIDs(ids)
	if err != nil {
		log.Println(err)
		return
	}

	response := providers.Response{Content: orders}
	json, err := json.Marshal(response)
	if err != nil {
		log.Println("Error while encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(json); err != nil {
		log.Println(err)
	}
}
