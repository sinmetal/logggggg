package main

import (
	"time"

	"github.com/google/uuid"
)

type AmbiguousType struct {
	Value interface{}
}

func LogAmbiguousType() {
	v := struct {
		Name  string
		Value int
	}{
		Name:  uuid.New().String(),
		Value: 1,
	}

	l := []*AmbiguousType{
		&AmbiguousType{Value: "Hello World"},
		&AmbiguousType{Value: -100},
		&AmbiguousType{Value: 0.10},
		&AmbiguousType{Value: time.Now()},
		&AmbiguousType{Value: v},
	}

	for _, v := range l {
		Output(v)
	}

	Output(l)
}
