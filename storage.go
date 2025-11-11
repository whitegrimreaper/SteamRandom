package main

import (
	"fmt"
	"errors"

	"gorm.io/gorm"
	//"database/sql/driver"
	"gorm.io/driver/sqlite"
)

// For future reference
// DB objects will have Entry appended
// so Game is from the framework, and GameEntry is for the DB
// TODO: could save all this info since storage space is not an issue at this level

// This table will hold all games I own, regardless of if they are played or not
type GameEntry struct {
	ID              uint   `gorm:"primaryKey"`
	AppID           uint64 `gorm:"uniqueIndex"`
	Name            string
	Playtime        int
	IsAdultGame     bool `gorm:"default:false"`
	HasPlayedYet    bool `gorm:"default:false"`
	HasListingIssue bool `gorm:"default:false"`
}

// Copy of the Game struct from OwnedGames for ease of use
type Game struct {
	AppID                    int    `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeForever          int    `json:"playtime_forever"`
	PlaytimeWindows          int    `json:"playtime_windows_forever"`
	PlaytimeMac              int    `json:"playtime_mac_forever"`
	PlaytimeLinux            int    `json:"playtime_linux_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
}

func setupDB() (db *gorm.DB){
	db, err := gorm.Open(sqlite.Open("dbs/games_test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&GameEntry{})
	return db
}

func doesGameEntryExist(targetGame int)(exists bool, err error) {
	var game GameEntry
	err = SteamDb.First(&game, "app_id = ?", targetGame).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	fmt.Printf("exists!\n")
	return true, nil
}

func tagGameAsNSFW(targetGame uint64, isNSFW bool)(err error) {
	var game GameEntry
	err = SteamDb.First(&game, "app_id = ?", targetGame).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	err = SteamDb.Model(&game).Update("IsAdultGame", isNSFW).Error
	if err != nil {
		fmt.Printf("Error in Update %+v\n", err)
		return err
	}
	return nil
}

func tagGameAsUnlisted(targetGame uint64, isUnlisted bool)(err error) {
	var game GameEntry
	err = SteamDb.First(&game, "app_id = ?", targetGame).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	err = SteamDb.Model(&game).Update("HasListingIssue", isUnlisted).Error
	if err != nil {
		fmt.Printf("Error in Update %+v\n", err)
		return err
	}
	return nil
}

func addGameEntry(game Game) {
	var gameEntry GameEntry
	
	err := SteamDb.Where(GameEntry{AppID: uint64(game.AppID), Name: game.Name}).FirstOrCreate(&gameEntry).Error
	if err != nil {
		fmt.Printf("Error in init %+v\n", err)
	}
}

func getAllGameEntries()([]GameEntry) {
	var gameEntries []GameEntry
	result := SteamDb.Find(&gameEntries)
	if result.Error != nil {
		fmt.Printf("HUH\n")
	}
	return gameEntries
}