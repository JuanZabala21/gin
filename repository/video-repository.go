package repository

import (
	"gin-project/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video entities.Video)
	Update(video entities.Video)
	Delete(video entities.Video)
	FindAll() []entities.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&entities.Video{}, &entities.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entities.Video) {
	db.connection.Save(&video)
}

func (db *database) Update(video entities.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entities.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entities.Video {
	var videos []entities.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
