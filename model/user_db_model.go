package model

import (
	"time"

	"gorm.io/gorm"
)

type UserDB struct {
	Id        int `gorm:"primary_key; auto_increment; not_null"`
	Name      string
	Email     string
	Alamat    string
	Password  string
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (e *UserDB) TableName() string {
	return "user"
}
