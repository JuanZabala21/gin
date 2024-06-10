package entities

import "time"

type Person struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:"name" binding:"required" gorm:"type:varchar(32)"`
	LastName string `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Age      int8   `json:"age" binding:"gte=1,lte=30"`
	Email    string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=10" gorm:"type:varchar(100)" validate:"is-cool"`
	Description string    `json:"description" binding:"required" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
