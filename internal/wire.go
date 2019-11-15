//+build wireinject

package internal

import (
	"github.com/go-chi/chi"
	"github.com/google/wire"

	"github.com/robovarga/szlh-delegations/internal/config"
	"github.com/robovarga/szlh-delegations/internal/parser"
	"github.com/robovarga/szlh-delegations/internal/repository"
	"github.com/robovarga/szlh-delegations/internal/scraper"
	"github.com/robovarga/szlh-delegations/internal/server"
)

func InitializeApp(dbConfig *config.DatabaseConfig) (*server.Server, error) {
	panic(wire.Build(
		repository.NewDBConnection,
		repository.NewRefereesRepository,
		repository.NewGamesRepository,
		parser.NewParser,
		scraper.NewScraper,
		server.NewServer,
	))
}

func InitializeWeb(dbConfig *config.DatabaseConfig) (*server.WebServer, error) {
	panic(wire.Build(
		chi.NewRouter,
		repository.NewDBConnection,
		repository.NewGamesRepository,
		server.NewHealthCheckHandler,
		server.NewWebServer,
	))
}
