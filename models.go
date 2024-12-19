package main

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `gorm:"primarykey;column:id"`
	Name string    `gorm:"column:name"`
}

func (User) TableName() string {
	return "users"
}

type Order struct {
	ID     uuid.UUID `gorm:"primarykey;column:id"`
	UserID uuid.UUID `gorm:"column:user_id"`
	Name   string    `gorm:"column:name"`
	Age    *int      `gorm:"column:age"`
}

func (Order) TableName() string {
	return "orders"
}

type Bill struct {
	ID     uuid.UUID  `gorm:"primarykey;column:id"`
	UserID *uuid.UUID `gorm:"column:user_id"`
	Name   *string    `gorm:"column:name"`
}

func (Bill) TableName() string {
	return "bills"
}
