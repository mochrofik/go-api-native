package models

import "time"

type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(200)" json:"title"`
	AuthorID    uint      `gorm:"type:integer" json:"author_id"`
	Author      Author    `gorm:"foreignKey:AuthorID" json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type BookResponse struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	AuthorID    uint               `gorm:"type:integer" json:"-"`
	Author      AuthorBookResponse `gorm:"foreignKey:AuthorID" json:"author"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at"`
}
