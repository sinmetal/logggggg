package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.NewTicker(6 * time.Minute)
	defer t.Stop()

	run(context.Background()) // Ticker待つのがめんどいので、とりあえず一発出す

	for {
		select {
		case <-t.C:
			ctx := context.Background()
			run(ctx)
		}
	}
}

func run(ctx context.Context) {
	fmt.Println("Start Run")
	defer fmt.Println("End Run")

	outputBigNumber()

	msg := &MessageV1{
		LogKindName: "Sample",
		Version:     "v1.0.1",
		Body:        "Hello World",
	}
	LogMessageV1(ctx, msg)

	LogAmbiguousType()
	LogBigMap()
	LogKV()
	LogMultiMap()
	LogParentChildren()
	LogDimensionalArrays()

	if rand.Intn(1000000) < 10 {
		LogMiracle()
	}
}
