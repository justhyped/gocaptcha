package gocaptcha

import (
	"net/http"
	"time"
)

type Settings struct {
	client                        *http.Client
	initialWaitTime, pollInterval time.Duration
	maxRetries                    int
}

func NewSettings() *Settings {
	return &Settings{
		client:          http.DefaultClient,
		initialWaitTime: 10 * time.Second,
		pollInterval:    5 * time.Second,
		maxRetries:      15,
	}
}
