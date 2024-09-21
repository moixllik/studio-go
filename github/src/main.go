package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	github_token := os.Getenv("GITHUB")
	fmt.Printf("GITHUB: %s\n\n", github_token)
	
	app := gin.Default()
	app.Static("/", "./dist")
	app.Run(":8080")
}
