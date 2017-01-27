package main

import (
	"github.com/joshblack/sql-tour/config"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	var env config.Environment

	err := envconfig.Process("service", &env)
	if err != nil {
		log.Fatalf("Error processing env: %v", err)
	}

	db, err := CreateStore(StoreConfig{
		User:     env.DatabaseUser,
		Password: env.DatabasePassword,
		Host:     env.DatabaseHost + ":" + env.DatabasePort,
		Database: env.DatabaseName,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}
