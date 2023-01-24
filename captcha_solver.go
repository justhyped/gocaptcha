package gocaptcha

import (
	"errors"
)

// SolveImageCaptcha solves an image captcha
func SolveImageCaptcha(payload *ImageCaptchaPayload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	imageSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		imageSolution, err = twoCaptchaSolveImageCaptcha(payload)
	case "AntiCaptcha":
		imageSolution, err = antiCaptchaSolveImageCaptcha(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		imageSolution, err = antiCaptchaSolveImageCaptcha(payload)
	}

	return imageSolution, err
}

// SolveRecaptchaV2 solves recaptcha V2
func SolveRecaptchaV2(payload *RecaptchaV2Payload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	captchaSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		captchaSolution, err = twoCaptchaSolveRecaptchaV2(payload)
	case "AntiCaptcha":
		captchaSolution, err = antiCaptchaSolveRecaptchaV2(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		captchaSolution, err = antiCaptchaSolveRecaptchaV2(payload)
	}

	return captchaSolution, err
}

// SolveRecaptchaV3 solves recaptcha V3
func SolveRecaptchaV3(payload *RecaptchaV3Payload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	captchaSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		captchaSolution, err = twoCaptchaSolveRecaptchaV3(payload)
	case "AntiCaptcha":
		captchaSolution, err = antiCaptchaSolveRecaptchaV3(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		captchaSolution, err = antiCaptchaSolveRecaptchaV3(payload)
	}

	return captchaSolution, err
}

// SolveHCaptcha solves hCaptcha
func SolveHCaptcha(payload *HCaptchaPayload) (*CaptchaResponse, error) {
	payload.SetDefaultValues()

	captchaSolution, err := &CaptchaResponse{}, errors.New("unsupported captcha service")

	switch payload.ServiceName {
	case "2Captcha":
		captchaSolution, err = twoCaptchaSolveHCaptcha(payload)
	case "AntiCaptcha":
		captchaSolution, err = antiCaptchaSolveHCaptcha(payload)
	case "CapMonster Cloud":
		// CapMonster Cloud has the same api
		// as AntiCaptcha so we just alter the api endpoint
		payload.CustomServiceUrl = "api.capmonster.cloud"
		captchaSolution, err = antiCaptchaSolveHCaptcha(payload)
	}

	return captchaSolution, err
}
