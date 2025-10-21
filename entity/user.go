package entity

import "time"

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"type:text"`
	Email     string `json:"email" gorm:"type:text;unique;not null"`
	Password  string `json:"password" gorm:"type:text;not null"`
	Username  string `json:"username" gorm:"type:text"`
	CreatedAt time.Time
}
