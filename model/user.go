package model

import (
	"github.com/icrowley/fake"
	"time"
)

type User struct {
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func SeedUser() User {
	user := User{
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
		Email:     fake.EmailAddress(),
	}
	return user
}
