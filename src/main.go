package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"hvalfangst/golang-fasthttp-pgx/src/configuration"
	"hvalfangst/golang-fasthttp-pgx/src/db"
	CarDetailsHandler "hvalfangst/golang-fasthttp-pgx/src/handler"
	"log"
)

func main() {
	// Fetch JSON based on key 'db' for file 'configuration.json' to be used as Db
	conf, err := configuration.Get()
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	// Connect to the database based on Configuration derived from 'configuration.json'
	database, _ := db.CreateDB(conf.(configuration.Db))

	r := router.New()
	r.GET("/cars/{id}", CarDetailsHandler.GetCarById(database))

	addr := ":8080"
	fmt.Printf("Server is listening on %s...\n", addr)
	if err := fasthttp.ListenAndServe(addr, r.Handler); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
