package models

import (
	"time"
)

type User struct {
	Id           string    `gorm:"not null"`
	Username     string    `gorm:"unique;not null;min:5;max:20"`
	Email        string    `gorm:"unique;not null"`
	Password     string    `gorm:"not null;min:8;max:16"`
	IsDelete     bool      `gorm:"not null;default:false"`
	CreateAt     time.Time `gorm:"autoCreateTime"`
	UpdateAt     time.Time `gorm:"autoUpdateTime"`
	CurrentMoney int       `gorm:"not null"`
	PathImage    string
}
