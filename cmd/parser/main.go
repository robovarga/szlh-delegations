package main

import (
	"log"
	"os"

	"github.com/robovarga/szlh-delegations/internal"
	"github.com/robovarga/szlh-delegations/internal/config"
)

// _ "github.com/joho/godotenv/autoload"

func main() {

	var (
		dbDriver = os.Getenv("DB_DRIVER")
		dbURI    = os.Getenv("DATABASE_URL")
		dbConfig = config.NewPostgresConfig(dbDriver, dbURI)
	)

	log.Println("started app")

	server, err := internal.InitializeApp(dbConfig)
	if err != nil {
		panic(err)
	}

	err = server.Handle()
	if err != nil {
		panic(err)
	}
}
