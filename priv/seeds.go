package main

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/joshblack/sql-tour/config"
	. "github.com/joshblack/sql-tour/model"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
	"net/url"
)

func main() {
	var env config.Environment

	err := envconfig.Process("service", &env)
	if err != nil {
		log.Fatalf("Error processing env: %v", err)
	}

	log.Println("Setting up the database connection...")
	db, err := setup(env)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("Beginning seeding...")

	tx, err := db.Beginx()
	if err != nil {
		log.Fatalln(err)
	}

	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	user := SeedUser()
	log.Println(user)
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func setup(env config.Environment) (*sqlx.DB, error) {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(env.DatabaseUser, env.DatabasePassword),
		Host:   env.DatabaseHost + ":" + env.DatabasePort,
		Path:   env.DatabaseName,
	}

	v := url.Values{}
	v.Set("sslmode", "disable")

	u.RawQuery = v.Encode()
	dataSourceName := u.String()

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
