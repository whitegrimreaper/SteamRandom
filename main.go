package main

import (
	//"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/Jleagle/steam-go/steamapi"
	"gorm.io/gorm"
	"math/rand"
	//"database/sql/driver"
	//"gorm.io/driver/sqlite"
)

type Game struct {
	ID       uint   `gorm:"primaryKey"`
	AppID    uint64 `gorm:"uniqueIndex"`
	Name     string
	Playtime int
}

var SteamDB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	apiKey := os.Getenv("STEAM_API_KEY")
	//ctx :=  context.Background();
	steamID := 76561198057886745

	steamDb, err := setupDB();
	if err != nil {
		panic("failed to setup db")
	}

	client := steamapi.NewClient()
	client.SetKey(apiKey)
	gamesResp, err := client.GetOwnedGames(int64(steamID))
	if err != nil {
		return
	}

	fmt.Printf("Retrieved %d games\n", gamesResp.GameCount)
	fmt.Printf("Ree %v\n", steamDb)

	gameList := gamesResp.Games

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gameList), func(i, j int) {
		gameList[i], gameList[j] = gameList[j], gameList[i]
	})

	fmt.Printf("Random game: %v\n", gameList[0])
}