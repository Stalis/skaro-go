package main

import (
	"flag"
	"log"

	"github.com/Stalis/skaro-go/skarolib"
)

var host = flag.String("host", "skaro.stalis-dev.xyz:8000", "http service address")

func main() {
	log.Println("Starting...")
	skaro.StartConnection(*host)
}
