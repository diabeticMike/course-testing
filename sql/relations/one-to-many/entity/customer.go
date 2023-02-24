package entity

import "github.com/google/uuid"

type Customer struct {
	ID     uuid.UUID
	Name   string
	Orders []Order `gorm:"foreignKey:CustomerID;references:ID"`
}

func (Customer) TableName() string {
	return "customer"
}
