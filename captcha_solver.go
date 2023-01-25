package gocaptcha

import (
	"context"
	"net/http"
	"time"
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

// SolveImageCaptcha uses the provider to fetch the solution of a captcha.
//
// The function returns ICaptchaResponse that has .Solution(), .ReportBad() and .ReportGood() that can be used
// to get the answer or report the quality of the captcha to the provider.
func (c *CaptchaSolver) SolveImageCaptcha(ctx context.Context, payload *ImageCaptchaPayload) (ICaptchaResponse, error) {
	return c.provider.SolveImageCaptcha(ctx, c.settings, payload)
}

// SolveRecaptchaV2 uses the provider to fetch the solution of a captcha.
//
// The function returns ICaptchaResponse that has .Solution(), .ReportBad() and .ReportGood() that can be used
// to get the answer or report the quality of the captcha to the provider.
func (c *CaptchaSolver) SolveRecaptchaV2(ctx context.Context, payload *RecaptchaV2Payload) (ICaptchaResponse, error) {
	return c.provider.SolveRecaptchaV2(ctx, c.settings, payload)
}

// SolveRecaptchaV3 uses the provider to fetch the solution of a captcha.
//
// The function returns ICaptchaResponse that has .Solution(), .ReportBad() and .ReportGood() that can be used
// to get the answer or report the quality of the captcha to the provider.
func (c *CaptchaSolver) SolveRecaptchaV3(ctx context.Context, payload *RecaptchaV3Payload) (ICaptchaResponse, error) {
	return c.provider.SolveRecaptchaV3(ctx, c.settings, payload)
}

// SolveHCaptcha uses the provider to fetch the solution of a captcha.
//
// The function returns ICaptchaResponse that has .Solution(), .ReportBad() and .ReportGood() that can be used
// to get the answer or report the quality of the captcha to the provider.
func (c *CaptchaSolver) SolveHCaptcha(ctx context.Context, payload *HCaptchaPayload) (ICaptchaResponse, error) {
	return c.provider.SolveHCaptcha(ctx, c.settings, payload)
}

// SetClient will set the client that is used when interacting with APIs of providers.
func (c *CaptchaSolver) SetClient(client *http.Client) {
	c.settings.client = client
}

// SetInitialWaitTime sets the time that is being waited after submitting a task to a provider before polling
func (c *CaptchaSolver) SetInitialWaitTime(waitTime time.Duration) {
	c.settings.initialWaitTime = waitTime
}

// SetPollInterval sets the time that is being waited in between result polls
func (c *CaptchaSolver) SetPollInterval(interval time.Duration) {
	c.settings.pollInterval = interval
}

// SetMaxRetries sets the maximum amount of polling
func (c *CaptchaSolver) SetMaxRetries(maxRetries int) {
	c.settings.maxRetries = maxRetries
}
