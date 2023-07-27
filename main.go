package main

import (
	"context"
	"log"
)

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggerService(service)

	_, err := service.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

}
