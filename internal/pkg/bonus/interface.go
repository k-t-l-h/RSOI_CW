package bonus

import "github.com/google/uuid"

type IRepo interface {
	GetBonus(UserUUID uuid.UUID) (int, int)
	SetBonus(UserUUID uuid.UUID, Extra int) (int, int)
	CreateBonus(UserUUID uuid.UUID) int
}
