package main

import (
	"context"
	"fmt"
	

	"github.com/joho/godotenv"
	"github.com/Jleagle/steam-go/steamapi"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {
	apiKey := os.Getenv("STEAM_API_KEY")
}