package main

import (
	"time"
	"log"
)

func main() {
	configuration, err := LoadConfiguration()
	if err != nil {
		panic(err)
	}
	log.Println(configuration.AssetEndpointURL)

	t := time.NewTicker(time.Second)
	for {
		log.Println("executed")
		<-t.C
	}
}
