package main

import (
	"ZopSmart_GoFr_Project/handler"
	"ZopSmart_GoFr_Project/stores/library"
	"gofr.dev/pkg/gofr"
)

func main() {
	// Create the application object
	app := gofr.New()

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	store := library.New()
	h := handlers.New(store)

	// Specify the different routes supported by this service
	app.GET("/book/:id", h.Get)
	app.POST("/book", h.Create)
	app.DELETE("/book/:id", h.Delete)
	app.Server.HTTP.Port = 8097

	app.Start()
}

