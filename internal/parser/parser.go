package parser

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/robovarga/szlh-delegations/internal/scraper"

	"github.com/sirupsen/logrus"

	"github.com/robovarga/szlh-delegations/internal/entity"
	"github.com/robovarga/szlh-delegations/internal/repository"

	"github.com/PuerkitoBio/goquery"
)

var reMatchDate = regexp.MustCompile(`(?m)([0-9]{2}\.[0-9]{2}\.[0-9]{4})`)

type Parser struct {
	referees *repository.RefereesRepository
	games    *repository.GamesRepository
	lists    *repository.ListRepository
	logger   *logrus.Logger
}

func NewParser(referees *repository.RefereesRepository,
	games *repository.GamesRepository,
	lists *repository.ListRepository,
	logger *logrus.Logger) *Parser {

	return &Parser{
		referees: referees,
		games:    games,
		lists:    lists,
		logger:   logger,
	}
}

func (p *Parser) Parse(listID int, body []byte) (games []*entity.Game) {

	p.logger.Debug("start parsing body")

	data := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		p.logger.Error(err)
	}

	list, err := p.lists.FindByID(listID)
	if err != nil {
		p.logger.Error(err)
		return
	}
	if list == nil {
		list = entity.NewList(listID, "Jolaus", scraper.ListsURL+strconv.Itoa(listID))
	}

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		game := entity.NewGame()
		game.SetList(list)

		heading := tablehtml.Prev()
		gameDate := reMatchDate.FindString(heading.Text())

		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(i int, columnHtml *goquery.Selection) {
				if columnHtml.Index() == 0 {
					gameID, err := strconv.Atoi(columnHtml.Text())
					if err != nil {
						p.logger.Error(err)
					}

					game.SetExternalID(gameID)
				}
				if columnHtml.Index() == 1 {
					teams := strings.Split(columnHtml.Text(), "vs.")

					game.SetHome(strings.TrimSpace(teams[0]))
					game.SetAway(strings.TrimSpace(teams[1]))
				}
				if columnHtml.Index() == 2 {
					infos := strings.Split(columnHtml.Text(), "-")

					game.SetVenue(strings.TrimSpace(infos[1]))

					timezone, err := time.LoadLocation("Europe/Warsaw")
					if err != nil {
						p.logger.Error("ERROR:", err)
					}

					gameDate, err := time.ParseInLocation("02.01.2006 15:04", gameDate+" "+strings.TrimSpace(infos[0]), timezone)
					if err != nil {
						p.logger.Error("ERROR:", err)
					}

					game.SetDate(gameDate)
				}
				if columnHtml.Index() == 3 {

					columnHtml.Find("span").Each(func(i int, refereeSpan *goquery.Selection) {

						// fmt.Println(refereeSpan.Text())

					})

				}
			})

		})

		p.logger.Debug("Parsed Game:", game.ExternalID())

		games = append(games, game)
	})

	p.logger.Debug("finish parsing body")

	return games
}
