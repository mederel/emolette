package main

import (
	"emolette/emolette"
	"log"
)

func main() {
	log.Println("Starting...")
	emolette.GenerateLenGoFiles()
	emolette.Play()
}
