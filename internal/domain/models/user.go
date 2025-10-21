package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Bio       string    `json:"bio"`
	AvatarUrl string    `json:"avatar_url"`
	Website   string    `json:"website"`
	Location  string    `json:"location"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Telephone string    `json:"telephone"`
}
