package gocaptcha

import (
	"context"
	"net/http"
)

type CaptchaSolver struct {
	provider IProvider
	settings *Settings
}

func NewCaptchaSolver(provider IProvider) *CaptchaSolver {
	return &CaptchaSolver{
		settings: NewSettings(),
		provider: provider,
	}
}

func (c *CaptchaSolver) SolveImageCaptcha(ctx context.Context, payload *ImageCaptchaPayload) (ICaptchaResponse, error) {
	return c.provider.SolveImageCaptcha(ctx, c.settings, payload)
}

func (c *CaptchaSolver) SolveRecaptchaV2(ctx context.Context, payload *RecaptchaV2Payload) (ICaptchaResponse, error) {
	return c.provider.SolveRecaptchaV2(ctx, c.settings, payload)
}

func (c *CaptchaSolver) SolveRecaptchaV3(ctx context.Context, payload *RecaptchaV3Payload) (ICaptchaResponse, error) {
	return c.provider.SolveRecaptchaV3(ctx, c.settings, payload)
}

func (c *CaptchaSolver) SolveHCaptcha(ctx context.Context, payload *HCaptchaPayload) (ICaptchaResponse, error) {
	return c.provider.SolveHCaptcha(ctx, c.settings, payload)
}

// SetClient will set the client that is used when interacting with APIs of providers.
func (c *CaptchaSolver) SetClient(client *http.Client) {
	c.settings.client = client
}

// SetInitialWaitTime sets the time in seconds that is being waited after submitting a task to a provider before polling
func (c *CaptchaSolver) SetInitialWaitTime(waitTime int) {
	c.settings.initialWaitTime = waitTime
}

// SetPollInterval sets the time in seconds that is being waited in between result polls
func (c *CaptchaSolver) SetPollInterval(interval int) {
	c.settings.pollInterval = interval
}

// SetMaxRetries sets the maximum amount of polling
func (c *CaptchaSolver) SetMaxRetries(maxRetries int) {
	c.settings.maxRetries = maxRetries
}
