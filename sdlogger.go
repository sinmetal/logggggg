package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/logging"
	"github.com/google/uuid"
)

type SDLogger struct {
	projectID string
	c         *logging.Client
}

func NewSDLogger(ctx context.Context, projectID string) (*SDLogger, error) {
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}
	client.OnError = func(err error) {
		log.Printf("logging.OnError: %v\n", err)
	}
	return &SDLogger{projectID: projectID, c: client}, nil
}

func (l *SDLogger) Write() {
	logger := l.c.Logger(uuid.New().String())
	defer func() {
		if err := logger.Flush(); err != nil {
			log.Printf("logging.Flush: %v\n")
		}
	}()

	logger.Log(logging.Entry{
		LogName: fmt.Sprintf("projects/%s/logs/logggggg/debug"),
		Payload: struct{ Anything string }{
			Anything: "The payload can be any type!",
		},
		Severity: logging.Debug,
	})
}
