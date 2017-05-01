package main

import (
	"log"
	"skaro-im/skarolib-go"
)

func main() {
	log.Println("Starting...")
	skaro.StartConnection("skaro.stalis-dev.xyz:8000")
}
