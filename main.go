package main

import (
	"log"
)

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggerService(service)
	ApiServer := NewApiServer(service)
	log.Fatal(ApiServer.Start(":3001"))

}
