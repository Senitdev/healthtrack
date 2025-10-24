package entity

import "time"

type PressureMeasurement struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id    int       `json:"user_id" gorm:"not nul;index"`
	Username   string    `json:"username" gorm:"not null"`
	Systole    int       `json:"systole" gorm:"type:int;not null"`
	Diastole   int       `json:"diastole" gorm:"type:int;not null"`
	Pulsion    int       `json:"pulse" gorm:"type:int;not null"`
	Context    string    `json:"context" gorm:"type:string;not null"`
	Taken_at   time.Time `json:"taken_at"`
	Notes      string    `json:"notes,omitempty" gorm:"type:string"`
	Created_At time.Time `json:"createdAt"`
	Updated_At time.Time `json:"updatedAt"`
	User       User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
