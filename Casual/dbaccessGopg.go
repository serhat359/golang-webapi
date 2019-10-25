package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SaveChapter(chapter *LhReadChapter) {
	var connectionString = getConnectionString()
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	db = db.Create(chapter)
	
	if db.Error != nil{
		panic(db.Error)
	}

	db.Close();
}
