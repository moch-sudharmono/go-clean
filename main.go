package main

import (
	"fmt"
	"os"

	config "github.com/joho/godotenv"
)

func main() {
	err := config.Load(".env")
	if err != nil {
		fmt.Println(".env not loaded properly")
		os.Exit(2)
	}
}
