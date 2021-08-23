package models

import "github.com/google/uuid"

type Ticket struct {
	TicketUUID uuid.UUID `json:"ticket_uuid"`
	FlightUUID uuid.UUID `json:"flight_uuid"`
	UserUUID   uuid.UUID `json:"user_uuid"`
}
