package main

import (
	"flag"
	"log"
	"skaro-im/skarolib-go"
)

var host = flag.String("host", "skaro.stalis-dev.xyz:8000", "http service address")

func main() {
	log.Println("Starting...")
	skaro.StartConnection(*host)
}
