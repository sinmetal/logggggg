package main

import (
	"context"
	"time"
)

func main() {
	t := time.NewTicker(6 * time.Minute)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			ctx := context.Background()
			outputBigNumber()

			msg := &MessageV1{
				LogKindName: "Sample",
				Version:     "v1.0.0",
				Body:        "Hello World",
			}
			LogMessageV1(ctx, msg)
		}
	}
}
