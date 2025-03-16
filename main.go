package main

import (
	//"context"
	"fmt"
	"os"
	

	"github.com/joho/godotenv"
	"github.com/Jleagle/steam-go/steamapi"
)

type Game struct {
	ID       uint   `gorm:"primaryKey"`
	AppID    uint64 `gorm:"uniqueIndex"`
	Name     string
	Playtime uint64
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	apiKey := os.Getenv("STEAM_API_KEY")
	//ctx :=  context.Background();
	steamID := 76561198057886745

	client := steamapi.NewClient()
	client.SetKey(apiKey)
	games, err := client.GetOwnedGames(int64(steamID))
	if err != nil {
		return
	}

	fmt.Printf("Retrieved %d games\n", games.GameCount)
}