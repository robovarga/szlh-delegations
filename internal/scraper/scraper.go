package scraper

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Scraper struct{}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) Scrape() ([]byte, error) {
	client := http.Client{}
	resp, err := client.Get("http://www.hockeyslovakia.sk/sk/stats/delegation-lists/204")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
