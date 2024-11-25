package main

import (
	"log"

	"github.com/yaminthered/shuttle-mate/api"
)

func main() {
	router := api.Start()

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
