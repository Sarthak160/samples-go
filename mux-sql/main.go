package main

import (

	"log"
)

func main() {
	a := &App{}
	err := a.Initialize(
		"postgres",
		"password",
		"postgres")

	if err != nil {
		log.Fatal("Failed to initialize app", err)
	}



	log.Printf("😃 Connected to 8010 port !!")

	a.Run(":8010")
}
