package models

import (
	"time"
)

type Course struct {
	ID        uint64    `gorm:"primary_id;auto_increment" json:"id"`
	Title     string    `gorm:"size:30;not null;unique" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Price     int64     `gorm:"not null" json:"price"`
	Discount  float64   `gorm:"not null" json:"discount"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (u *Course) Prepare() {

}
