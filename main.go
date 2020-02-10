package main

import (
	"context"
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
	outputBigNumber()

	msg := &MessageV1{
		LogKindName: "Sample",
		Version:     "v1.0.1",
		Body:        "Hello World",
	}
	LogMessageV1(ctx, msg)
}
