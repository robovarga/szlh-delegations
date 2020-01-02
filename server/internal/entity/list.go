package entity

import (
	"time"
)

type List struct {
	listID     int
	name       string
	listURL    string
	dateAdd    time.Time
	dateUpdate time.Time
}

func (l *List) DateUpdate() time.Time {
	return l.dateUpdate
}

func (l *List) SetDateUpdate(dateUpdate time.Time) {
	l.dateUpdate = dateUpdate
}

func (l *List) DateAdd() time.Time {
	return l.dateAdd
}

func (l *List) SetDateAdd(dateAdd time.Time) {
	l.dateAdd = dateAdd
}

func (l *List) ListURL() string {
	return l.listURL
}

func (l *List) SetListURL(listURL string) {
	l.listURL = listURL
}

func (l *List) Name() string {
	return l.name
}

func (l *List) SetName(name string) {
	l.name = name
}

func (l *List) ListID() int {
	return l.listID
}

func (l *List) SetListID(listID int) {
	l.listID = listID
}

func NewList(
	listID int,
	name string,
	listURL string,
) *List {
	return &List{
		listID:  listID,
		name:    name,
		listURL: listURL,
	}
}
