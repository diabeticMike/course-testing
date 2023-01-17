package entity

import "github.com/google/uuid"

type Team struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Users []User    `gorm:"many2many:users_teams;" json:"users,omitempty"`
}
