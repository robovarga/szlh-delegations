package entity

import (
	"time"
)

type Referee struct {
	refereeID           int
	name                string
	dateAdd, dateUpdate time.Time
}

func NewReferee(refereeID int,
	name string,
	dateAdd time.Time,
	dateUpdate time.Time,
) *Referee {
	return &Referee{
		refereeID:  refereeID,
		name:       name,
		dateAdd:    dateAdd,
		dateUpdate: dateUpdate,
	}
}

func (r *Referee) DateUpdate() time.Time {
	return r.dateUpdate
}

func (r *Referee) DateAdd() time.Time {
	return r.dateAdd
}

func (r *Referee) Name() string {
	return r.name
}

func (r *Referee) ID() int {
	return r.refereeID
}
