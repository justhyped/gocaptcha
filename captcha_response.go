package gocaptcha

import (
	"context"
	"errors"
)

type ICaptchaResponse interface {
	// Solution will return the solution of the captcha as a string
	Solution() string

	// ReportBad reports the captcha to be invalid if the provider and captcha type support it.
	ReportBad(ctx context.Context) error

	// ReportGood reports the captcha to be valid if the provider and captcha type support it.
	ReportGood(ctx context.Context) error
}

type CaptchaResponse struct {
	solution, taskId  string
	isAlreadyReported bool

	// these are set internally for each captcha type
	reportGood func(ctx context.Context) error
	reportBad  func(ctx context.Context) error
}

func (a *CaptchaResponse) Solution() string {
	return a.solution
}

func (a *CaptchaResponse) ReportBad(ctx context.Context) error {
	if a.isAlreadyReported {
		return nil
	}

	if a.reportBad == nil {
		return errors.ErrUnsupported
	}

	if err := a.reportBad(ctx); err != nil {
		return err
	}

	a.isAlreadyReported = true

	return nil
}

func (a *CaptchaResponse) ReportGood(ctx context.Context) error {
	if a.isAlreadyReported {
		return nil
	}

	if a.reportGood == nil {
		return errors.ErrUnsupported
	}

	if err := a.reportGood(ctx); err != nil {
		return err
	}

	a.isAlreadyReported = true

	return nil
}

var _ ICaptchaResponse = (*CaptchaResponse)(nil)
