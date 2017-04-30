package main

import (
	"fmt"
	"skaro-im/skarolib-go"
)

func main() {
	fmt.Println("Starting...")
	skaro.StartConnection("skaro.stalis-dev.xyz:8000")
}
