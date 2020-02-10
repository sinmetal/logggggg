package main

import (
	"time"
)

func main() {
	t := time.NewTicker(6 * time.Minute)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			outputBigNumber()
		}
	}
}
