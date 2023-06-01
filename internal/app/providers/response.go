package providers

type Response struct {
	Content []Order `json:"content"`
}

type Order struct {
	OrderID     int    `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int64  `json:"store_id"`
	DataCreated string `json:"date_created"`
}
