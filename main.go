package main

import (
	"fmt"

	"github.com/aasthagarg128/ZopSmart_GoFr_Project/stores/library"

	handlers "github.com/aasthagarg128/ZopSmart_GoFr_Project/handler"
	"gofr.dev/pkg/gofr"
)

func main() {
	fmt.Println("Hello World")
	// Create the application object
	app := gofr.New()
	fmt.Println("Staring the application")

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	store := library.New()
	h := handlers.New(store)

	// Specify the different routes supported by this service
	app.GET("/book/:id", h.Get)
	app.POST("/book", h.Create)
	app.DELETE("/book/:id", h.Delete)
	app.GET("/ping", h.Ping)
	app.Server.HTTP.Port = 8097

	fmt.Println("Starting the application at port 8097")
	app.Start()
}