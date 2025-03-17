package main

import (
	"fmt"
	"errors"

	"gorm.io/gorm"
	//"database/sql/driver"
	"gorm.io/driver/sqlite"
)

// This table will hold all games I own, regardless of if they are played or not
type GameEntry struct {
	ID       uint   `gorm:"primaryKey"`
	AppID    uint64 `gorm:"uniqueIndex"`
	Name     string
	Playtime int
}

func setupDB() (db *gorm.DB, err error){
	db, err = gorm.Open(sqlite.Open("dbs/games_test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	err = db.AutoMigrate(&GameEntry{})
	if err != nil {
		return nil, err
	}

	return db, err
}

func doesGameEntryExist(targetGame int)(exists bool, err error) {
	var game GameEntry
	err = steamDb.First(&game, "id = ?", targetGame).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Game Entry doesn't exist!")
			return false, nil
		}
		return false, err
	}
	fmt.Printf("exists!\n")
	return true, nil
}

func addGameEntry() {
	return
}