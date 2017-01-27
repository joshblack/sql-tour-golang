package main

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/url"
)

type StoreConfig struct {
	User     string
	Password string
	Host     string
	Database string
}

func CreateStore(s StoreConfig) (*sqlx.DB, error) {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(s.User, s.Password),
		Host:   s.Host,
		Path:   s.Database,
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
