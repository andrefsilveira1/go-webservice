package main

import (
	"context"
	"fmt"
	"time"
)

type LoggerService struct {
	next Service
}

func NewLoggerService(next Service) Service {
	return &LoggerService{
		next: next,
	}
}

func (s *LoggerService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("FACT: %v | ERR: %s | TOOK: %v", fact, err, time.Since(start))
	}(time.Now())
	return s.next.GetCatFact(ctx)
}
