package main

import (
	//"context"
	"fmt"
	"os"
	"time"
	//"slices"
	//"reflect"

	"github.com/joho/godotenv"
	"github.com/Jleagle/steam-go/steamapi"
	"gorm.io/gorm"
	"math/rand"
	//"database/sql/driver"
	//"gorm.io/driver/sqlite"
)
var steamDb *gorm.DB

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

	/*for _, game := range gameList {
		if(game.Name == "Subverse") {
			details, err := client.GetAppDetails(uint(game.AppID), COUNTRY_CODE, LANGUAGE_CODE, []string{})
			if err != nil {
				return
			}

			fmt.Printf("%v\n", reflect.TypeOf(details.Data.ContentDescriptors.IDs))
			if ids, ok := details.Data.ContentDescriptors.IDs.([]interface{}); ok {
				for _, v := range ids {
					//fmt.Printf("%v\n", reflect.TypeOf(v))
					if num, ok := v.(float64); ok && num == 4 {
						fmt.Printf("This game likely has sexual content and shouldn't be included: %s\n", game.Name)
					}
				}
			}
			
		//if(slices.Contains(details.Data.ContentDescriptors.IDs, "4")) {
		//	fmt.Printf("Game details: %v\n", details.Data.ContentDescriptors.IDs)
		//}
		}
		

		
		time.Sleep(100 * time.Millisecond)
	}*/
}