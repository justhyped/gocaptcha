package gocaptcha

import "context"

type IProvider interface {
	SolveImageCaptcha(ctx context.Context, settings *Settings, payload *ImageCaptchaPayload) (ICaptchaResponse, error)
	SolveRecaptchaV2(ctx context.Context, settings *Settings, payload *RecaptchaV2Payload) (ICaptchaResponse, error)
	SolveRecaptchaV3(ctx context.Context, settings *Settings, payload *RecaptchaV3Payload) (ICaptchaResponse, error)
	SolveHCaptcha(ctx context.Context, settings *Settings, payload *HCaptchaPayload) (ICaptchaResponse, error)
}

type ICaptchaResponse interface {
	Solution() string
	ReportBad(ctx context.Context) error
	ReportGood(ctx context.Context) error
}
