package model

import (
	uuid "github.com/satori/go.uuid"
)

func CreateSession() uuid.UUID {
	sID := uuid.NewV4()
	return sID
}
