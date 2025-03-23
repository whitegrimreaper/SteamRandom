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
var SteamDb *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	apiKey := os.Getenv("STEAM_API_KEY")
	//ctx :=  context.Background();
	steamID := 76561198057886745

	SteamDb = setupDB();

	client := steamapi.NewClient()
	client.SetKey(apiKey)
	gamesResp, err := client.GetOwnedGames(int64(steamID))
	if err != nil {
		return
	}

	fmt.Printf("Retrieved %d games\n", gamesResp.GameCount)
	fmt.Printf("Ree %v\n", SteamDb)

	gameList := gamesResp.Games

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(gameList), func(i, j int) {
		gameList[i], gameList[j] = gameList[j], gameList[i]
	})

	fmt.Printf("Random game: %v\n", gameList[0])

	for _, game := range gameList {
		exists, err := doesGameEntryExist(game.AppID)
		if err != nil {
			return
		}
		if !exists {
			fmt.Printf("Adding game to db\n")
			addGameEntry(game)
		}
	}
}

func checkGamesForNSFWContent(client steamapi.Client) {
	// This is a very expensive function to call, and will make an API call for every item in your steam library.
	// For me that is 1k plus calls per run, more than 1% of the 100k daily call limit
	// Use sparingly

	allGames := getAllGameEntries()

	for _, game := range allGames {
		// This is the issue line
		details, err := client.GetAppDetails(uint(game.AppID), COUNTRY_CODE, LANGUAGE_CODE, []string{})
		if err != nil {
			// Keep randomly getting errors here for specific games
			// Some are removed like MapleStory 2, some are private/dupe style games like many test servers
			fmt.Printf("Issue getting info for app: %d, name %s\n", game.AppID, game.Name)
			fmt.Printf("Error: %v\n", err.Error())
		} else if ids, ok := details.Data.ContentDescriptors.IDs.([]interface{}); ok {
			for _, v := range ids {
				//fmt.Printf("%v\n", reflect.TypeOf(v))
				// So the content descriptors are for violence and sexual content
				// 2 is for violence, which is fine on Twitch/YT in games
				// 5 is for any mature content, also not very helpful
				// 1 is general nudity or sexual content, 3 is for "adult only" sexual content, and 4 is "gratuitous" sexual content.
				// Personally I don't have many nsfw games, but I would generally say that 3 and 4 are tags for games you should not play on
				// regular livestreams. 1 is hit or miss
				if num, ok := v.(float64); ok && (num == 3 || num == 4){
					fmt.Printf("This game likely has sexual content and shouldn't be included: %s\n", game.Name)
					tagGameAsNSFW(game.AppID, true)
				}
			}
		}
	}	
}