package scraper

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const ListsURL = "http://www.hockeyslovakia.sk/sk/stats/delegation-lists/"

type Scraper struct {
	logger *logrus.Logger
}

func NewScraper(logger *logrus.Logger) *Scraper {
	return &Scraper{
		logger: logger,
	}
}

func (s *Scraper) Scrape(listID int) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Get(ListsURL + strconv.Itoa(listID))
	if err != nil {
		s.logger.Fatal(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
