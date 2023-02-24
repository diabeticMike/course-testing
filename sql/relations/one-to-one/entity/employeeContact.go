package entity

import "github.com/google/uuid"

type EmployeeContact struct {
	ID            uuid.UUID
	ContactNumber int
	EmployeeID    uuid.UUID
}

func (EmployeeContact) TableName() string {
	return "employee_contact"
}
