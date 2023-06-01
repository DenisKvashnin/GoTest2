package order

import (
	"TestTask/internal/app/providers"
	"database/sql"
	"fmt"
	"log"
)

type Repository struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func New(
	host string,
	port string,
	user string,
	password string,
	dbName string) *Repository {

	return &Repository{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName,
	}
}

func (o Repository) Instance() *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		o.host, o.port, o.user, o.password, o.dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Ошибка при открытии соединения с базой данных: %s", err)
	}
	return db
}

func (o Repository) InitDb() {
	db := o.Instance()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error while closing db. %s", err)
		}
	}(db)

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS "order"(
		    order_id bigint PRIMARY KEY, 
		    status character varying, 
		    store_id bigint, 
		    date_create timestamp without time zone
		)
	`)
	if err != nil {
		panic("Ошибка подключения к БД.")
	}
}

func (o Repository) SaveAll(orders []providers.Order) {
	for _, ord := range orders {
		err := o.save(ord)
		if err != nil {
			log.Fatalf("Incorrect save: %s", err)
			return
		}
	}
}

func (o Repository) save(order providers.Order) error {
	db := o.Instance()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error while closing db. %s", err)
		}
	}(db)

	_, err := db.Exec(`
			INSERT INTO "order"(order_id, status, store_id, date_create) VALUES ($1, $2, $3, $4)
		`, order.OrderID, order.Status, order.StoreId, nil) //TODO: разобраться с датами
	if err != nil {
		log.Fatalf("Save error: %s", err)
		return err
	}

	return nil
}
