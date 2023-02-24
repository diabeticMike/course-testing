package entity

import "github.com/google/uuid"

type Employee struct {
	ID              uuid.UUID
	Name            string
	Department      string
	EmployeeContact EmployeeContact `gorm:"foreignKey:EmployeeID;references:ID"`
}

func (Employee) TableName() string {
	return "employee"
}
