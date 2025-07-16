package database

import (

	"time"
	"file-organizer/model"
)

func Add(name, filetype string, size int64, path string) error {
	file := model.File{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Name: name,
		Type: filetype,
		Size: size,
		Path: path,
	}
	return db.Create(&file).Error
}