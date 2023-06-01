package base

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	Db       *sql.DB
}

func New(
	host string,
	port string,
	user string,
	password string,
	dbName string) *Database {

	db := Database{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName,
		Db:       nil,
	}
	err := db.initDb()
	if err != nil {
		panic(err)
	}

	return &db
}

func (o *Database) initDb() error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		o.host, o.port, o.user, o.password, o.dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}
	if err != nil {
		panic(err)
	}

	o.Db = db

	_, err = o.Db.Exec(`
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

	return nil
}

func (o Database) CloseDb() error {
	err := o.Db.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
