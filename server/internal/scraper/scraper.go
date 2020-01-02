package scraper

import (
	"io/ioutil"
	"net/http"

	"github.com/robovarga/szlh-delegations/internal/entity"

	"github.com/sirupsen/logrus"
)

const ListsURL = "http://www.hockeyslovakia.sk/"

type Scraper struct {
	logger *logrus.Logger
}

func NewScraper(logger *logrus.Logger) *Scraper {
	return &Scraper{
		logger: logger,
	}
}

func (s *Scraper) Scrape(list *entity.List) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Get(ListsURL + list.ListURL())
	if err != nil {
		s.logger.Fatal(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
