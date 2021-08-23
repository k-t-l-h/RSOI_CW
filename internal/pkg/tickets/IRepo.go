package tickets

import (
	"github.com/google/uuid"
	"rsoi-kp-k-t-l-h/internal/models"
)

type IRepo interface {
	GetTickets(uuid uuid.UUID) ([]models.Ticket, int)
	GetTicket(uuid uuid.UUID) (models.Ticket, int)
	DeleteTicket(uuid uuid.UUID) int
	CreateTicket(ticket models.Ticket) int
}
