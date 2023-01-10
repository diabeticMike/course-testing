package entity

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Teams []Team    `gorm:"many2many:users_teams;" json:"teams,omitempty"`
}
