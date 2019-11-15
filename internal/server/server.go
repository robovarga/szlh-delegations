package server

import (
	"log"

	"github.com/robovarga/szlh-delegations/internal/parser"
	"github.com/robovarga/szlh-delegations/internal/repository"
	"github.com/robovarga/szlh-delegations/internal/scraper"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	scraper *scraper.Scraper
	parser  *parser.Parser
	games   *repository.GamesRepository
}

func NewServer(scraper *scraper.Scraper, parser *parser.Parser, gamesRepository *repository.GamesRepository) *Server {
	return &Server{scraper, parser, gamesRepository}
}

func (s *Server) Handle() error {

	log.Println("Start fetching")

	data, err := s.scraper.Scrape()
	if err != nil {
		return err
	}

	log.Println("Finish Fetching")

	games := s.parser.Parse(data)

	for _, game := range games {
		err = s.games.Insert(game)
		if err != nil {
			return err
		}
	}

	return nil
}
