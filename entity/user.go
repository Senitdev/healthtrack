package entity

import "time"

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"type:text"`
	Email     string `json:"email" gorm:"type:text;unique;not nul"`
	Password  string `json:"-" gorm:"type:text"`
	CreatedAt time.Time
}
