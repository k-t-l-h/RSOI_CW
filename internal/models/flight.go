package models

import (
	"github.com/google/uuid"
	"time"
)

type Flight struct {
	ID          uuid.UUID `json:"id"`
	From        uuid.UUID `json:"from"`
	FromCity    string    `json:"from_city,omitempty"`
	To          uuid.UUID `json:"to"`
	ToCity      string    `json:"to_city,omitempty"`
	Date        time.Time `json:"date"`
	Tickets     int       `json:"tickets"`
	TicketsSold int       `json:"tickets_sold"`
}
