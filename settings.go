package gocaptcha

import "net/http"

type Settings struct {
	client                                    *http.Client
	initialWaitTime, pollInterval, maxRetries int
}

func NewSettings() *Settings {
	return &Settings{
		client:          http.DefaultClient,
		initialWaitTime: 10,
		pollInterval:    5,
		maxRetries:      15,
	}
}
