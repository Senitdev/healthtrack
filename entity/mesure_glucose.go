package entity

import "time"

type GlucoseMeasurement struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id   int       `json:"user_id" gorm:"not null;index"`
	Username  string    `json:"username" gorm:"not null"`
	Value     float64   `json:"value" gorm:"type:numeric"`
	Unite     string    `json:"unite" gorm:"type:text"`
	Context   string    `json:"context" gorm:"type:text"`
	Take_at   time.Time `json:"take_at"`
	Notes     string    `json:"notes,omitempty" gorm:"type:text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
