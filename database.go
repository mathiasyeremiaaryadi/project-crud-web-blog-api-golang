package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SCHEMA"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // output ke console
		logger.Config{
			SlowThreshold: time.Second, // threshold untuk slow query
			LogLevel:      logger.Info, // Level: Silent, Error, Warn, Info
			Colorful:      true,        // warna
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return db, err
	}

	return db, err
}

func MigrateTables(databaseConnection *gorm.DB) {
	databaseConnection.AutoMigrate(
		&Post{},
		&Tag{},
	)
}

func InsertNewPost(post Post) (Post, error) {
	result := DBConnection.Create(&post)

	if result.RowsAffected != 1 || result.Error != nil {
		return Post{}, result.Error
	}

	return post, nil
}

func UpdatePost(postID int, post Post) (Post, error) {
	var existingPost Post
	err := DBConnection.Preload("Tags").First(&existingPost, postID).Error
	if err != nil {
		return Post{}, err
	}

	result := DBConnection.Model(&existingPost).Where("id = ?", postID).Updates(Post{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
	})
	if result.RowsAffected != 1 || result.Error != nil {
		return Post{}, result.Error
	}

	var newTags []Tag
	for _, tag := range post.Tags {
		tempTag := Tag{
			Name:   tag.Name,
			PostID: existingPost.ID,
		}
		newTags = append(newTags, tempTag)
	}

	err = DBConnection.Model(&existingPost).Association("Tags").Unscoped().Replace(newTags)
	if err != nil {
		return Post{}, err
	}

	var updatedPost Post
	err = DBConnection.Preload("Tags").First(&updatedPost, postID).Error
	if err != nil {
		return Post{}, err
	}

	return updatedPost, nil
}

func DeletePost(postID int) error {
	err := DBConnection.Where("post_id = ?", postID).Delete(&Tag{}).Error
	if err != nil {
		return err
	}

	err = DBConnection.Delete(&Post{}, postID).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPost(postID int) (Post, error) {
	var post Post
	err := DBConnection.Preload("Tags").First(&post, postID).Error
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func GetPosts(term string) ([]Post, error) {
	var (
		posts []Post
		err   error
	)
	if term != "" {
		err = DBConnection.Preload("Tags").Where("title LIKE ?", "%"+term+"%").Find(&posts).Error

	} else {
		err = DBConnection.Preload("Tags").Find(&posts).Error
	}

	if err != nil {
		return posts, err
	}

	return posts, nil
}
