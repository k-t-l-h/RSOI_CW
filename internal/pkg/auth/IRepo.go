package auth

import "RSOI_CW/internal/models"

type IRepo interface {
	GetUser(login string, password string) (models.User, int)
	GetUsers() ([]models.User, int)
	AddUser(user models.User) (models.User, int)
}
