package table

import "time"

type Meta struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	BookSN      string
	Author      string
	Description string
	//BookCover        string
	PublishDate  time.Time
	CreateAt     time.Time
	LastModifyAt time.Time
}
