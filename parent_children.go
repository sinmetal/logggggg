package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Parent struct {
	Key       string
	Children  []*Child
	CreatedAt time.Time
}

type Child struct {
	Key   string
	Value interface{}
}

func LogParentChildren() {
	crn := []*Child{
		&Child{
			Key:   uuid.New().String(),
			Value: rand.Intn(100),
		},
		&Child{
			Key:   uuid.New().String(),
			Value: "Moge",
		},
	}
	p := &Parent{
		Key:       uuid.New().String(),
		Children:  crn,
		CreatedAt: time.Now(),
	}

	Output(p)
}
