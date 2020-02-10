package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type MessageV1 struct {
	LogKindName string
	Version     string
	Body        string
}

func LogMessageV1(ctx context.Context, msg *MessageV1) {
	e := json.NewEncoder(os.Stdout)
	if err := e.Encode(msg); err != nil {
		fmt.Printf("failed json.Encode... %+v\n", err)
	}
}
