package main

import (
	"context"
	"github.com/justhyped/gocaptcha"
)

func main() {
	cs := gocaptcha.NewCaptchaSolver(CustomProvider{})

	img, err := cs.SolveImageCaptcha(context.TODO(), &gocaptcha.ImageCaptchaPayload{
		Base64String: "",
	})
	if err != nil {
		panic(err)
	}

	_ = img.Solution()
	_ = img.ReportBad(context.TODO())
	_ = img.ReportGood(context.TODO())
}

type CustomProvider struct {
}

func (c CustomProvider) SolveImageCaptcha(ctx context.Context, settings *gocaptcha.Settings, payload *gocaptcha.ImageCaptchaPayload) (gocaptcha.ICaptchaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomProvider) SolveRecaptchaV2(ctx context.Context, settings *gocaptcha.Settings, payload *gocaptcha.RecaptchaV2Payload) (gocaptcha.ICaptchaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomProvider) SolveRecaptchaV3(ctx context.Context, settings *gocaptcha.Settings, payload *gocaptcha.RecaptchaV3Payload) (gocaptcha.ICaptchaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomProvider) SolveHCaptcha(ctx context.Context, settings *gocaptcha.Settings, payload *gocaptcha.HCaptchaPayload) (gocaptcha.ICaptchaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomProvider) SolveTurnstile(ctx context.Context, settings *gocaptcha.Settings, payload *gocaptcha.TurnstilePayload) (gocaptcha.ICaptchaResponse, error) {
	//TODO implement me
	panic("implement me")
}

type CustomResponse struct {
}

func (c CustomResponse) Solution() string {
	//TODO implement me
	panic("implement me")
}

func (c CustomResponse) ReportBad(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c CustomResponse) ReportGood(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

// Interface guard, used to make sure the struct implements all required methods
var _ gocaptcha.IProvider = (*CustomProvider)(nil)
var _ gocaptcha.ICaptchaResponse = (*CustomResponse)(nil)
