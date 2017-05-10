package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Stalis/libskaro-go"
)

type Config struct {
	Host string
}

func main() {
	log.Println("Starting...")
	inputFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("config.json not exists")
	}

	var configuration Config

	json.Unmarshal(inputFile, &configuration)

	skaro.StartConnection(configuration.Host)
}
