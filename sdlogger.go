package main

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
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
	logger := l.c.Logger("logggggg/debug")
	defer func() {
		if err := logger.Flush(); err != nil {
			log.Printf("logging.Flush: %v\n", err)
		}
		log.Println("logging api: Flush")
	}()

	logger.Log(logging.Entry{
		Payload: struct{ Anything string }{
			Anything: "The payload can be any type!",
		},
		Severity: logging.Debug,
	})
}
