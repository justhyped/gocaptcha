package gocaptcha

import (
	"errors"
	"github.com/justhyped/gocaptcha/providers"
)

// SolveImageCaptcha solves image captchas
func SolveImageCaptcha(payload *ImageCaptchaPayload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	imageSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		imageSolution, err = providers.TwoCaptchaSolveImageCaptcha(payload)
	case "AntiCaptcha":
		imageSolution, err = providers.AntiCaptchaSolveImageCaptcha(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		imageSolution, err = providers.AntiCaptchaSolveImageCaptcha(payload)
	}

	return imageSolution, err
}

// SolveRecaptchaV2 solves recaptcha V2
func SolveRecaptchaV2(payload *RecaptchaV2Payload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	captchaSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		captchaSolution, err = providers.TwoCaptchaSolveRecaptchaV2(payload)
	case "AntiCaptcha":
		captchaSolution, err = providers.AntiCaptchaSolveRecaptchaV2(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		captchaSolution, err = providers.AntiCaptchaSolveRecaptchaV2(payload)
	}

	return captchaSolution, err
}

// SolveRecaptchaV3 solves recaptcha V3
func SolveRecaptchaV3(payload *RecaptchaV3Payload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	captchaSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		captchaSolution, err = providers.TwoCaptchaSolveRecaptchaV3(payload)
	case "AntiCaptcha":
		captchaSolution, err = providers.AntiCaptchaSolveRecaptchaV3(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		captchaSolution, err = providers.AntiCaptchaSolveRecaptchaV3(payload)
	}

	return captchaSolution, err
}
