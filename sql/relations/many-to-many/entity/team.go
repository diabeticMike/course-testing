package entity

import "github.com/google/uuid"

type Team struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Users []User    `gorm:"many2many:user_team;" json:"users,omitempty"`
}

func (Team) TableName() string {
	return "team"
}
