package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getGormDB() *gorm.DB {
	var connectionString = getConnectionString()
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}

func SaveChapter(chapter *LhReadChapter) {
	db := getGormDB()
	defer db.Close()

	db = db.Create(chapter)
	if db.Error != nil{
		panic(db.Error)
	}
}

func SaveScore(score *LhScore){
	db := getGormDB()
	defer db.Close()

	db = db.Create(score)
	if db.Error != nil{
		panic(db.Error)
	}
}
