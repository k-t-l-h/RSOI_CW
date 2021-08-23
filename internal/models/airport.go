package models

import "github.com/google/uuid"

type Airport struct {
	AirportUUID uuid.UUID `json:"airport_uuid"`
	Name        string    `json:"name"`
}
