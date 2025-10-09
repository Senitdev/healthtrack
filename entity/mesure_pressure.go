package entity

import "time"

type PressureMeasurement struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id    int       `json:"userId" gorm:"not nul;index"`
	Systole    int       `json:"systole" gorm:"type:int"`
	Diastole   int       `json:"diastole" gorm:"type:int"`
	Pulsion    int       `json:"pulse" gorm:"type:int"`
	Context    string    `json:"context" gorm:"type:string"`
	Taken_at   time.Time `json:"taken_at"`
	Notes      string    `json:"notes,omitempty" gorm:"type:string"`
	Created_At time.Time `json:"createdAt"`
	Updated_At time.Time `json:"updatedAt"`
	User       User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
