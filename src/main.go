package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/bitokss/bitok-user-service/src/app"
	"github.com/bitokss/bitok-user-service/src/repo/postgres/v1"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	repo.DB = repo.PostgresInit()
}

func main() {
	port := flag.Int("port", 8080, "service port")
	flag.Parse()
	app.StartApp(fmt.Sprintf(":%d", port))
}
