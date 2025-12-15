package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/rafaelcamelo31/graduate-go-course/projects/auction/configuration/database/mongodb"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
