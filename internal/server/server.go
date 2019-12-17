package server

import (
	"github.com/sirupsen/logrus"

	"github.com/robovarga/szlh-delegations/internal/parser"
	"github.com/robovarga/szlh-delegations/internal/repository"
	"github.com/robovarga/szlh-delegations/internal/scraper"

	_ "github.com/joho/godotenv/autoload"
)

const listID = 208

type Server struct {
	scraper *scraper.Scraper
	parser  *parser.Parser
	games   *repository.GamesRepository
	logger  *logrus.Logger
}

func NewServer(scraper *scraper.Scraper,
	parser *parser.Parser,
	gamesRepository *repository.GamesRepository,
	logger *logrus.Logger) *Server {

	return &Server{
		scraper,
		parser,
		gamesRepository,
		logger,
	}
}

func (s *Server) Handle() error {

	s.logger.Println("Start fetching")

	data, err := s.scraper.Scrape(listID)
	if err != nil {
		return err
	}

	s.logger.Println("Finish Fetching")

	games := s.parser.Parse(listID, data)

	for _, game := range games {
		err = s.games.Insert(game)
		if err != nil {
			return err
		}
	}

	return nil
}
