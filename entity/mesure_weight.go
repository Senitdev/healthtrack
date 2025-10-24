package entity

import "time"

type WeightMeasurement struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id   int       `json:"user_id" gorm:"not null;index"`
	Username  string    `json:"username" gorm:"not null"`
	Weight    float64   `json:"weight"`           // en kg
	Height    float64   `json:"height,omitempty"` // en cm
	BMI       float64   `json:"bmi,omitempty"`
	Unit      string    `json:"unit" gorm:"default:kg"`
	Context   string    `json:"context,omitempty"`
	TakenAt   time.Time `json:"taken_at"`
	Notes     string    `json:"notes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
