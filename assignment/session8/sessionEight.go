package session8

import (
	"fmt"
	"log"

	"github.com/claravelita/training-golang-mnc/assignment/session8/api"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func MainAssigmentSession8() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Println("application has started...")
	newEcho := echo.New()
	serverAPI := api.NewServer(newEcho)

	serverAPI.InitializeServer()
}
