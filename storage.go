package main

import (
	"gorm.io/gorm"
	//"database/sql/driver"
	"gorm.io/driver/sqlite"
)

func setupDB() (db *gorm.DB, err error){
	db, err = gorm.Open(sqlite.Open("dbs/games_test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	err = db.AutoMigrate(&Game{})
	if err != nil {
		return nil, err
	}

	return db, err
}