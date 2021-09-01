package database

import (
	"fmt"
)

//Config to maintain DB configuration properties
type Config struct {
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.gxxa4.mongodb.net/%s?retryWrites=true&w=majority", config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME)
	return connectionString
}
