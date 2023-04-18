package main

import (
	"os"

	"boilerplate/server"

	"boilerplate/utils"

	"github.com/joho/godotenv"
)

func main() {
	app := server.New()
	err := godotenv.Load()
	if err != nil {
		utils.InfoLog.Println("Error loading .env file")
	}
	utils.InfoLog.Println("starting the server on port", os.Getenv("API_PORT"))
	if err := app.Listen(os.Getenv("API_PORT")); err != nil {
		utils.ErrorLog.Panic(err.Error())
	}
}
