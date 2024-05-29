package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/siampanya/capcha-with-gemini/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	app.App(ctx)

}
