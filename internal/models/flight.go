package models

import (
	"github.com/google/uuid"
	"time"
)

type Flight struct {
	From uuid.UUID `json:"from"`
	To   uuid.UUID `json:"to"`
	Date time.Time `json:"date"`
}
