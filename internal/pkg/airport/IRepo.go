package airport

import (
	"github.com/google/uuid"
	"rsoi-kp-k-t-l-h/internal/models"
)

type IRepo interface {
	SelectAirports() ([]models.Airport, int)
	SelectAirport(uuid uuid.UUID) (models.Airport, int)
}
