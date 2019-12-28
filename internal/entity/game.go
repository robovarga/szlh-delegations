package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Game struct {
	gameID     int
	gameUUID   uuid.UUID
	externalID int
	listID     *List
	home       string
	away       string
	venue      string
	date       time.Time
	referees   []*Referee
}

func NewGame(gameID int,
	gameUUID uuid.UUID,
	externalID int,
	listID *List,
	home, away, venue string,
	date time.Time,
	// referees []*Referee,
) *Game {

	return &Game{
		gameID:     gameID,
		gameUUID:   gameUUID,
		externalID: externalID,
		listID:     listID,
		home:       home,
		away:       away,
		venue:      venue,
		date:       date,
		// referees:   referees,
	}
}

func GenerateGame() *Game {
	return &Game{
		gameUUID: uuid.NewV1(),
	}
}

func (g *Game) UUID() uuid.UUID {
	return g.gameUUID
}

func (g *Game) SetUUID(gameUUID uuid.UUID) {
	g.gameUUID = gameUUID
}

func (g *Game) ID() int {
	return g.gameID
}

func (g *Game) SetID(gameID int) {
	g.gameID = gameID
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

func (g *Game) AddReferee(referee *Referee) {
	g.referees = append(g.referees, referee)
}

func (g *Game) Referees() []*Referee {
	return g.referees
}
