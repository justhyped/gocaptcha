package internal

import (
	"context"
	"time"
)

// SleepWithContext sleeps for the specified duration but returns early with an error if the context is cancelled.
func SleepWithContext(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}
