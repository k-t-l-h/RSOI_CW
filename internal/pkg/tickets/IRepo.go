package tickets

import (
	"RSOI_CW/internal/models"
	"github.com/google/uuid"
)

type IRepo interface {
	GetAllTickets() ([]models.Ticket, int)
	GetTickets(uuid uuid.UUID) ([]models.Ticket, int)
	GetTicket(uuid uuid.UUID) (models.Ticket, int)
	DeleteTicket(uuid uuid.UUID) int
	CreateTicket(ticket models.Ticket) int
}
