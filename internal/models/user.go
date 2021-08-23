package models

import "github.com/google/uuid"

type User struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	UUID     uuid.UUID `json:"uuid"`
	Role     string
}
