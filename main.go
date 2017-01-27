package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Environment struct {
	Port             string `envconfig:"PORT" default:"3000"`
	DatabaseUser     string `envconfig:"DATABASE_USERNAME" default:"postgres"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"sql_tour_dev"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"5432"`
}

type Schema struct {
	create string
	drop   string
}

func (s Schema) Postgres() (string, string) {
	return s.create, s.drop
}

var schema = Schema{
	create: `
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	first_name text NOT NULL,
	last_name text NOT NULL,
	email text UNIQUE NOT NULL,
	inserted_at timestamp default now(),
	updated_at timestamp default now()
);

CREATE TABLE posts (
	id SERIAL PRIMARY KEY,
	title text,
	user_id integer REFERENCES users (id),
	inserted_at timestamp default now(),
	updated_at timestamp default now()
);
`,
	drop: `
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;
`,
}

type User struct {
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type Post struct {
	Title      string    `db:"title"`
	UserId     int64     `db:"user_id"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func main() {
	var env Environment

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

	create, drop := schema.Postgres()
	db.MustExec(drop)
	log.Println(drop)
	db.MustExec(create)
	log.Println(create)
}
