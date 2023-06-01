package order

import (
	"TestTask/internal/app/providers"
	Base "TestTask/internal/app/repositories/base"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	database *Base.Database
}

func New(database *Base.Database) *Repository {
	return &Repository{
		database: database,
	}
}

func (o Repository) SaveAll(orders []providers.Order) {
	for _, order := range orders {
		err := o.save(order)
		if err != nil {
			log.Println("Incorrect save:", err)
			return
		}
	}
}

func (o Repository) save(order providers.Order) error {
	_, err := o.database.Db.Exec(`
			INSERT INTO "order"(order_id, status, store_id, date_create) VALUES ($1, $2, $3, $4)
		`, order.OrderID, order.Status, order.StoreId, nil) //TODO: разобраться с датами
	if err != nil {
		log.Println("Save error:", err)
		return err
	}

	return nil
}
