package main

import "time"

type Post struct {
	ID        int `gorm:"primaryKey,autoIncrement"`
	Title     string
	Content   string
	Category  string
	Tags      []Tag
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

type Tag struct {
	ID        int `gorm:"primaryKey,autoIncrement"`
	Name      string
	PostID    int
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}
