package config

type Environment struct {
	Port             string `envconfig:"PORT" default:"3000"`
	DatabaseUser     string `envconfig:"DATABASE_USERNAME" default:"postgres"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"sql_tour_dev"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"5432"`
}
