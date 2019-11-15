package entity

import uuid "github.com/satori/go.uuid"

type Referee struct {
	refereeID uuid.UUID
	name      string
}
