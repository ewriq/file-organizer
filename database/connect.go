package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"file-organizer/model"
)

var db *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("./database/sql/sql.db"), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı!")
	}

	err = database.AutoMigrate(&model.File{})
	if err != nil {
        panic("Migration başarısız!")
    }

	db = database
	fmt.Println("Veritabanı bağlantısı başarılı.")
}