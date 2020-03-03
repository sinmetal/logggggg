package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Miracle struct {
	Message       string
	Parent        *Parent
	AmbiguousType []*AmbiguousType
}

func LogMiracle() {
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

	var l []*AmbiguousType
	{
		v := struct {
			Name  string
			Value int
		}{
			Name:  uuid.New().String(),
			Value: 1,
		}

		l = []*AmbiguousType{
			&AmbiguousType{Value: "Hello World"},
			&AmbiguousType{Value: -100},
			&AmbiguousType{Value: 0.10},
			&AmbiguousType{Value: time.Now()},
			&AmbiguousType{Value: v},
		}
	}

	v := &Miracle{
		Message:       "このログに出会えるとは運がいいね！",
		Parent:        p,
		AmbiguousType: l,
	}

	Output(v)
}
