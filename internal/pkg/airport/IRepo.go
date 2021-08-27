package airport

import (
	"RSOI_CW/internal/models"
	"github.com/google/uuid"
)

type IRepo interface {
	SelectAirports() ([]models.Airport, int)
	SelectAirport(uuid uuid.UUID) (models.Airport, int)
}
