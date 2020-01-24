package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sinmetal/slog/v2"
)

type Logger struct {
	f *os.File
	l *log.Logger
}

func NewLogger(logName string, mode int) (*Logger, error) {
	switch mode {
	case 0:
		return &Logger{l: log.New(os.Stdout, "", 0)}, nil
	default:
		f, err := os.Create(fmt.Sprintf("/var/log/%s", logName))
		if err != nil {
			return nil, err
		}

		return &Logger{f: f, l: log.New(f, "", 0)}, nil
	}
}

func (l *Logger) WriteString(text string) {
	ctx := context.Background()
	ctx = slog.WithValue(ctx)
	slog.Info(ctx, text)
	slog.Flush(ctx)
}

func (l *Logger) Close() error {
	return l.f.Close()
}
