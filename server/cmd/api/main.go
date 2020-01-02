package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/robovarga/szlh-delegations/internal/config"

	"github.com/robovarga/szlh-delegations/internal"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		dbDriver       = os.Getenv("DB_DRIVER")
		dbURI          = os.Getenv("DATABASE_URL")
		databaseConfig = config.NewPostgresConfig(dbDriver, dbURI)
	)
	log := config.NewLogger()

	log.Info("Loaded ENV Driver:", dbDriver)

	srv, err := internal.InitializeWeb(databaseConfig, log)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("initialized SRV.")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer cancel()
		srv.Serve(ctx)
	}()

	wg.Wait()
}
