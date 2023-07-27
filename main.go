package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggerService(service)

	fact, err := service.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", fact)
}
