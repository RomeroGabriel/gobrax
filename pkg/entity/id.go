package entity

import (
	"errors"

	"github.com/google/uuid"
)

type ID = uuid.UUID

var ErrIDIsRequired = errors.New("id is required")

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	if s == "" {
		return uuid.Nil, ErrIDIsRequired
	}
	id, err := uuid.Parse(s)
	return ID(id), err
}
