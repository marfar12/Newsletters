package model

import (
	"github.com/google/uuid"
)

type Newsletter struct {
	ID       uuid.UUID
	Name     string
	Desc     string
	EditorId uuid.UUID
}

type Editor struct {
	ID       uuid.UUID
	Email    string
	Password string
}

type Subscription struct {
	NewsletterId    uuid.UUID
	Email           string
	UnsubscribeCode string
}
