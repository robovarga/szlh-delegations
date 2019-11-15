package parser

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/robovarga/szlh-delegations/internal/entity"
	"github.com/robovarga/szlh-delegations/internal/repository"

	"github.com/PuerkitoBio/goquery"
)

var reMatchDate = regexp.MustCompile(`(?m)([0-9]{2}\.[0-9]{2}\.[0-9]{4})`)

type Parser struct {
	referees *repository.RefereesRepository
	games    *repository.GamesRepository
}

func NewParser(referees *repository.RefereesRepository, games *repository.GamesRepository) *Parser {
	return &Parser{referees: referees, games: games}
}

func (p *Parser) Parse(body []byte) (games []*entity.Game) {

	log.Println("start parsing body")

	data := bytes.NewReader(body)

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		game := entity.NewGame()

		heading := tablehtml.Prev()
		gameDate := reMatchDate.FindString(heading.Text())

		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(i int, columnHtml *goquery.Selection) {
				if columnHtml.Index() == 0 {
					gameID, err := strconv.Atoi(columnHtml.Text())
					if err != nil {
						log.Println(err)
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
						log.Println("ERROR:", err)
					}

					gameDate, err := time.ParseInLocation("02.01.2006 15:04", gameDate+" "+strings.TrimSpace(infos[0]), timezone)
					if err != nil {
						log.Println("ERROR:", err)
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

		log.Println("Parsed Game:", game.ExternalID())

		games = append(games, game)
	})

	log.Println("finish parsing body")

	return games
}
