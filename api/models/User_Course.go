package models

import "time"

type UserCourse struct {
	ID        uint64    `gorm:"primary_id;auto_increment" json:"id"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	CourseID  int64     `gorm:"not null" json:"course_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}
