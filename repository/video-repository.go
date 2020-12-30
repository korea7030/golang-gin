package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // import 필요
	"gitlab.com/pragmaticreviews/gin-poc/entity"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	// sqlite open
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	// Video, Person entity Migrate
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

// Close DB
func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

// Save
func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

// Update
func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

// Delete
func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

// Select
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	// get foreign key info
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
