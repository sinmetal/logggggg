package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/vvakame/sdlog/buildlog"
)

type MessageV1 struct {
	LogKindName string
	Version     string
	Body        string
}

func LogMessageV1(ctx context.Context, msg *MessageV1) {
	le := buildlog.NewLogEntry(ctx)
	j, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("failed json.Marshal... %+v\n", err)
		return
	}
	le.Message = string(j)

	e := json.NewEncoder(os.Stdout)
	if err := e.Encode(le); err != nil {
		fmt.Printf("failed json.Encode... %+v\n", err)
	}
}
