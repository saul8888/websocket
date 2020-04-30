package main

import (
	"websocket/client"

	"github.com/labstack/echo"
)

func main() {
	route := echo.New()
	// Create a simple file server
	route.Static("/", "./server")

	// Configure websocket route
	route.GET("/ws", client.ClientRead)

	// Start listening for incoming chat messages
	go client.ClientWrite()

	// Start the server on localhost port 8000 and log any errors
	route.Logger.Fatal(route.Start(":8000"))
}
