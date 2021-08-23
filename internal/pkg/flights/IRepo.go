package flights

import (
	"github.com/google/uuid"
	"rsoi-kp-k-t-l-h/internal/models"
)

type IRepo interface {
	CreateFlight(flight models.Flight) int
	ReadFlight(FlightUUID uuid.UUID) (models.Flight, int)
	ReadFlights() ([]models.Flight, int)
	UpdateFlight(FlightUUID uuid.UUID, flight models.Flight) int
}
