package flights

import (
	"RSOI_CW/internal/models"
	"github.com/google/uuid"
)

type IRepo interface {
	CreateFlight(flight models.Flight) int
	ReadFlight(FlightUUID uuid.UUID) (models.Flight, int)
	ReadFlights() ([]models.Flight, int)
	UpdateFlight(FlightUUID uuid.UUID, flight models.Flight) int
}
