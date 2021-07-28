package models

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/config"
)

// var db *gorm.DB
var db *sql.DB

type Image struct {
	gorm.Model
	ImageID     int64  `gorm:"primary_key;column:id;type:int(11) unsigned auto_increment;not null;" json:"id"`
	FileName    string `json:"filename"`
	Title       string `json:"title"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&Image{})
}

func (i Image) UploadImage() {
	queryString := fmt.Sprintf("INSERT INTO Images " +
		"(FileName, Title, Tags, Description, created_at, updated_at, deleted_at) " +
		"VALUES (?, ?, ?, ?,  CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), NULL)")

	_, err := db.Query(queryString, i.FileName, i.Title, i.Tags, i.Description)
	if err != nil {
		panic(err.Error())
	}
}

func GetAllImages() []Image {
	var Images []Image

	results, err := db.Query("SELECT * from images")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var image Image
		err = results.Scan(&image.ImageID, &image.FileName, &image.Title, &image.Tags, &image.Description, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			panic(err.Error())
		}

		Images = append(Images, image)

	}

	return Images
}

func GetImageById(Id int64) (*Image, *sql.DB) {
	// var image Image
	// db := db.Where("idimages = ?", Id).Find(&image)

	var image Image

	queryString := fmt.Sprintf("SELECT * from images where ID=?")

	results, err := db.Query(queryString, Id)

	for results.Next() {
		err = results.Scan(&image.ImageID, &image.FileName, &image.Title, &image.Tags, &image.Description, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			panic(err.Error())
		}
	}

	return &image, db
}

func DeleteImage(Id int64) {
	// db.Where("idimages = ?", Id).Delete(image)
	// fmt.Println(Id)
	queryString := fmt.Sprintf("DELETE from images where ID=?")

	_, err := db.Query(queryString, Id)
	if err != nil {
		panic(err.Error())
	}
	// return image
}

func GetRecentId() (id int64) {
	var image Image
	results, err := db.Query("SELECT * FROM images ORDER BY ID DESC LIMIT 1;")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&image.ImageID, &image.FileName, &image.Title, &image.Tags, &image.Description, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt)
		if err != nil {
			panic(err.Error())
		}
	}

	return image.ImageID
}
