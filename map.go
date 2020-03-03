package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type MultiMap struct {
	Tag map[string]string
	Any map[string]interface{}
}

func LogKV() {
	m := make(map[string]string)
	m["Hoge"] = uuid.New().String()
	m["Fuga"] = "Greenだよ"

	Output(m)
}

func LogBigMap() {
	m := make(map[string]int64)
	for i := 0; i < 1000; i++ {
		m[uuid.New().String()] = rand.Int63n(1000000)
	}

	Output(m)
}

func LogMultiMap() {
	tm := make(map[string]string)
	tm["Hoge"] = "Hoge"
	tm["Fuga"] = "Fuga"

	am := make(map[string]interface{})
	am["Hoge"] = 100
	am["Fuga"] = time.Now()
	am["Moge"] = []string{"Momomo", "Bobobobo"}

	mm := &MultiMap{
		Tag: tm,
		Any: am,
	}

	Output(mm)
}
