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
	scraper   *scraper.Scraper
	parser    *parser.Parser
	gamesRepo *repository.GamesRepository
	listsRepo *repository.ListRepository
	logger    *logrus.Logger
}

func NewServer(scraper *scraper.Scraper,
	parser *parser.Parser,
	gamesRepository *repository.GamesRepository,
	listRepository *repository.ListRepository,
	logger *logrus.Logger) *Server {

	return &Server{
		scraper:   scraper,
		parser:    parser,
		gamesRepo: gamesRepository,
		listsRepo: listRepository,
		logger:    logger,
	}
}

func (s *Server) Handle() error {

	lists, err := s.listsRepo.GetLists()
	if err != nil {
		return err
	}

	s.logger.Println("Start fetching")

	data, err := s.scraper.Scrape(lists[0])
	if err != nil {
		return err
	}

	s.logger.Println("Finish Fetching")

	games := s.parser.Parse(lists[0], data)

	for _, game := range games {
		gameID, gameUUID, err := s.gamesRepo.CheckGame(game.ExternalID(), game.List().ListID())
		if err != nil {
			return err
		}

		if gameID > 0 {
			game.SetID(gameID)
			game.SetUUID(gameUUID)
			err = s.gamesRepo.Update(game)
		} else {
			err = s.gamesRepo.Insert(game)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
