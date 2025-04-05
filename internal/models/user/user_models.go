package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *User) TableName() string {
	return "users"
}
