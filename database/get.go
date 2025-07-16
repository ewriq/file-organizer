package database

import (
	"file-organizer/model"
)

func GetFilesByExtension() ([]model.File, error) {
	var files []model.File
	err := db.Find(&files).Error
	return files, err
}
