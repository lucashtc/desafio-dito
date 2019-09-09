package main

import (
	//"github.com/lucashtc/desafio-dito/autocomplete"
	"github.com/lucashtc/desafio-dito/timeline"
	"log"
)

func main() {
	//coletora.InitServer()
	if err := timeline.TimelineOrder(); err != nil {
		log.Fatal(err)
	}
}
