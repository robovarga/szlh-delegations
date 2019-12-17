package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Game struct {
	gameID     uuid.UUID
	externalID int
	listID     *List
	home       string
	away       string
	venue      string
	date       time.Time
	referees   []*Referee
}

func NewGame() *Game {
	return &Game{
		gameID: uuid.NewV1(),
	}
}

func (g *Game) ID() uuid.UUID {
	return g.gameID
}

func (g *Game) Date() time.Time {
	return g.date
}

func (g *Game) SetDate(date time.Time) {
	g.date = date
}

func (g *Game) Venue() string {
	return g.venue
}

func (g *Game) SetVenue(venue string) {
	g.venue = venue
}

func (g *Game) Away() string {
	return g.away
}

func (g *Game) SetAway(away string) {
	g.away = away
}

func (g *Game) Home() string {
	return g.home
}

func (g *Game) SetHome(home string) {
	g.home = home
}

func (g *Game) ExternalID() int {
	return g.externalID
}

func (g *Game) SetExternalID(externalID int) {
	g.externalID = externalID
}

func (g *Game) List() *List {
	return g.listID
}

func (g *Game) SetList(listID *List) {
	g.listID = listID
}
