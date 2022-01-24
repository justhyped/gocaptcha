package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/justhyped/gocaptcha"
	"github.com/justhyped/gocaptcha/helpers"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func TwoCaptchaSolveRecaptchaV2(payload *gocaptcha.RecaptchaV2Payload) (*gocaptcha.CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "2captcha.com"
	}

	captchaResponse := payload.CreateRecaptchaResponse()

	var submitUrl = fmt.Sprint(protocol, "://", payload.CustomServiceUrl, "/in.php")
	var pollUrl = fmt.Sprint(protocol, "://", payload.CustomServiceUrl, "/res.php")

	urlParams := url.Values{}
	urlParams.Set("key", payload.ServiceApiKey)
	urlParams.Set("json", "1")
	urlParams.Set("method", "userrecaptcha")
	urlParams.Set("googlekey", payload.EndpointKey)
	urlParams.Set("pageurl", payload.EndpointUrl)
	urlParams.Set("invisible", helpers.BooleanToString(payload.IsInvisibleCaptcha))

	request, err := http.Get(fmt.Sprint(submitUrl, "?", urlParams.Encode()))
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	response, err := helpers.ReadResponseBody(request)
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	responseAsJSON := twoCaptchaResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	if responseAsJSON.Status == 1 {
		captchaResponse.TaskId = responseAsJSON.Request
	} else {
		return &gocaptcha.CaptchaResponse{}, errors.New(responseAsJSON.ErrorText)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	for captchaResponse.Solution == "" {
		urlParams = url.Values{}
		urlParams.Set("key", payload.ServiceApiKey)
		urlParams.Set("json", "1")
		urlParams.Set("action", "get")
		urlParams.Set("id", captchaResponse.TaskId)

		request, err = http.Get(fmt.Sprint(pollUrl, "?", urlParams.Encode()))
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		response, err := helpers.ReadResponseBody(request)
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		responseAsJSON := twoCaptchaResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		if responseAsJSON.Status == 1 {
			captchaResponse.Solution = responseAsJSON.Request
			return captchaResponse, nil
		}

		if responseAsJSON.Status == 0 && !strings.Contains(responseAsJSON.Request, "NOT_READY") {
			return &gocaptcha.CaptchaResponse{}, errors.New(responseAsJSON.Request)
		}

		tries++
		if tries == payload.MaxRetries {
			return &gocaptcha.CaptchaResponse{}, errors.New("captcha took too long to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return captchaResponse, nil
}

func TwoCaptchaSolveRecaptchaV3(payload *gocaptcha.RecaptchaV3Payload) (*gocaptcha.CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "2captcha.com"
	}

	captchaResponse := payload.CreateRecaptchaResponse()

	var submitUrl = fmt.Sprint(protocol, "://", payload.CustomServiceUrl, "/in.php")
	var pollUrl = fmt.Sprint(protocol, "://", payload.CustomServiceUrl, "/res.php")

	urlParams := url.Values{}
	urlParams.Set("key", payload.ServiceApiKey)
	urlParams.Set("json", "1")
	urlParams.Set("method", "userrecaptcha")
	urlParams.Set("version", "v3")
	urlParams.Set("googlekey", payload.EndpointKey)
	urlParams.Set("pageurl", payload.EndpointUrl)

	if payload.Action != "" {
		urlParams.Set("action", payload.Action)
	}

	if payload.IsEnterprise {
		urlParams.Set("enterprise", "1")
	}

	request, err := http.Get(fmt.Sprint(submitUrl, "?", urlParams.Encode()))
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	response, err := helpers.ReadResponseBody(request)
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	responseAsJSON := twoCaptchaResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	if responseAsJSON.Status == 1 {
		captchaResponse.TaskId = responseAsJSON.Request
	} else {
		return &gocaptcha.CaptchaResponse{}, errors.New(responseAsJSON.ErrorText)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	for captchaResponse.Solution == "" {
		urlParams = url.Values{}
		urlParams.Set("key", payload.ServiceApiKey)
		urlParams.Set("json", "1")
		urlParams.Set("action", "get")
		urlParams.Set("id", captchaResponse.TaskId)

		request, err = http.Get(fmt.Sprint(pollUrl, "?", urlParams.Encode()))
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		response, err := helpers.ReadResponseBody(request)
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		responseAsJSON := twoCaptchaResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		if responseAsJSON.Status == 1 {
			captchaResponse.Solution = responseAsJSON.Request
			return captchaResponse, nil
		}

		if responseAsJSON.Status == 0 && !strings.Contains(responseAsJSON.Request, "NOT_READY") {
			return &gocaptcha.CaptchaResponse{}, errors.New(responseAsJSON.Request)
		}

		tries++
		if tries == payload.MaxRetries {
			return &gocaptcha.CaptchaResponse{}, errors.New("captcha took too long to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return captchaResponse, nil
}

func TwoCaptchaSolveImageCaptcha(payload *gocaptcha.ImageCaptchaPayload) (*gocaptcha.CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "2captcha.com"
	}

	captchaResponse := payload.CreateImageCaptchaResponse()

	var submitUrl = fmt.Sprint(protocol, "://", payload.CustomServiceUrl, "/in.php")
	var pollUrl = fmt.Sprint(protocol, "://", payload.CustomServiceUrl, "/res.php")

	urlParams := url.Values{}
	urlParams.Set("key", payload.ServiceApiKey)
	urlParams.Set("json", "1")
	urlParams.Set("method", "base64")
	urlParams.Set("body", payload.Base64String)

	if payload.InstructionsForSolver != "" {
		urlParams.Set("textinstructions", payload.InstructionsForSolver)
	}

	if payload.CaseSensitive {
		urlParams.Set("regsense", "1")
	}

	request, err := http.Get(fmt.Sprint(submitUrl, "?", urlParams.Encode()))
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	response, err := helpers.ReadResponseBody(request)
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	responseAsJSON := twoCaptchaResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &gocaptcha.CaptchaResponse{}, err
	}

	if responseAsJSON.Status == 1 {
		captchaResponse.TaskId = responseAsJSON.Request
	} else {
		return &gocaptcha.CaptchaResponse{}, errors.New(responseAsJSON.ErrorText)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	for captchaResponse.Solution == "" {
		urlParams = url.Values{}
		urlParams.Set("key", payload.ServiceApiKey)
		urlParams.Set("json", "1")
		urlParams.Set("action", "get")
		urlParams.Set("id", captchaResponse.TaskId)

		request, err = http.Get(fmt.Sprint(pollUrl, "?", urlParams.Encode()))
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		response, err := helpers.ReadResponseBody(request)
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		responseAsJSON := twoCaptchaResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &gocaptcha.CaptchaResponse{}, err
		}

		if responseAsJSON.Status == 1 {
			captchaResponse.Solution = responseAsJSON.Request
			return captchaResponse, nil
		}

		if responseAsJSON.Status == 0 && !strings.Contains(responseAsJSON.Request, "NOT_READY") {
			return &gocaptcha.CaptchaResponse{}, errors.New(responseAsJSON.Request)
		}

		tries++
		if tries == payload.MaxRetries {
			return &gocaptcha.CaptchaResponse{}, errors.New("captcha took too long to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return captchaResponse, nil
}

type twoCaptchaResponse struct {
	Status    int    `json:"status"`
	Request   string `json:"request"`
	ErrorText string `json:"error_text"`
}
