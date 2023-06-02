package order

import (
	"TestTask/internal/app/providers"
	Base "TestTask/internal/app/repositories/base"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
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

func (o Repository) GetOrderByIDs(IDs []int) ([]providers.Order, error) {
	idParams := make([]interface{}, len(IDs))
	for i, id := range IDs {
		idParams[i] = id
	}

	idPlaceholders := make([]string, len(IDs))
	for i := range idPlaceholders {
		idPlaceholders[i] = "$" + strconv.Itoa(i+1)
	}
	idCondition := fmt.Sprintf("WHERE order_id IN (%s)", strings.Join(idPlaceholders, ","))

	query := fmt.Sprintf(`
		SELECT order_id, status, store_id, date_create
		FROM "order"
		%s
	`, idCondition)

	rows, err := o.database.Db.Query(query, idParams...)
	if err != nil {
		log.Println("Error while querying orders:", err)
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	orders := make([]providers.Order, 0)
	for rows.Next() {
		var order providers.Order
		err := rows.Scan(&order.OrderID, &order.Status, &order.StoreId, &order.DataCreated)
		orders = append(orders, order)
		if err != nil {
			log.Println("Error while scanning order:", err)
			continue
		}
	}
	if err := rows.Err(); err != nil {
		log.Println("Error while iterating over orders:", err)
		return nil, err
	}

	return orders, nil
}
