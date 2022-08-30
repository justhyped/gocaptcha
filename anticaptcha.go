package gocaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/justhyped/gocaptcha/internal"
	"net/http"
	"strconv"
	"time"
)

func antiCaptchaSolveRecaptchaV2(payload *RecaptchaV2Payload) (*CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "api.anti-captcha.com"
	}

	captchaResponse := payload.CreateRecaptchaResponse()

	createTaskUrl := fmt.Sprintf("%v://%v/createTask", protocol, payload.CustomServiceUrl)
	getTaskUrl := fmt.Sprintf("%v://%v/getTaskResult", protocol, payload.CustomServiceUrl)

	typeTask := map[string]interface{}{
		"type":        "NoCaptchaTaskProxyless",
		"websiteURL":  payload.EndpointUrl,
		"websiteKey":  payload.EndpointKey,
		"isInvisible": payload.IsInvisibleCaptcha}

	createTask := map[string]interface{}{"clientKey": payload.ServiceApiKey, "task": typeTask}
	jsonValue, _ := json.Marshal(createTask)

	request, err := http.Post(createTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return &CaptchaResponse{}, err
	}

	response, err := internal.ReadResponseBody(request)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	responseAsJSON := antiCaptchaCreateResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	if responseAsJSON.ErrorID == 0 {
		captchaResponse.TaskId = strconv.Itoa(responseAsJSON.TaskID)
	} else {
		return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	resultData := map[string]string{"clientKey": payload.ServiceApiKey, "taskId": fmt.Sprint(captchaResponse.TaskId)}
	jsonValue, _ = json.Marshal(resultData)

	for captchaResponse.Solution == "" {
		request, err := http.Post(getTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return &CaptchaResponse{}, err
		}

		response, err := internal.ReadResponseBody(request)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		responseAsJSON := antiCaptchaResultResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		if responseAsJSON.ErrorID != 0 {
			return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
		}

		if responseAsJSON.Status == "ready" {
			captchaResponse.Solution = responseAsJSON.Solution.RecaptchaResponse
			return captchaResponse, nil
		}

		tries++
		if tries == payload.MaxRetries {
			return &CaptchaResponse{}, errors.New("captcha took longer than 115 seconds to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return &CaptchaResponse{}, errors.New("reached end of function")
}

func antiCaptchaSolveRecaptchaV3(payload *RecaptchaV3Payload) (*CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "api.anti-captcha.com"
	}

	captchaResponse := payload.CreateRecaptchaResponse()

	createTaskUrl := fmt.Sprintf("%v://%v/createTask", protocol, payload.CustomServiceUrl)
	getTaskUrl := fmt.Sprintf("%v://%v/getTaskResult", protocol, payload.CustomServiceUrl)

	typeTask := map[string]interface{}{
		"type":       "RecaptchaV3TaskProxyless",
		"websiteURL": payload.EndpointUrl,
		"websiteKey": payload.EndpointKey,
		"minScore":   0.3,
	}

	if payload.Action != "" {
		typeTask["pageAction"] = payload.Action
	}

	if payload.IsEnterprise {
		typeTask["isEnterprise"] = true
	}

	createTask := map[string]interface{}{"clientKey": payload.ServiceApiKey, "task": typeTask}
	jsonValue, _ := json.Marshal(createTask)

	request, err := http.Post(createTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return &CaptchaResponse{}, err
	}

	response, err := internal.ReadResponseBody(request)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	responseAsJSON := antiCaptchaCreateResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	if responseAsJSON.ErrorID == 0 {
		captchaResponse.TaskId = strconv.Itoa(responseAsJSON.TaskID)
	} else {
		return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	resultData := map[string]string{"clientKey": payload.ServiceApiKey, "taskId": fmt.Sprint(captchaResponse.TaskId)}
	jsonValue, _ = json.Marshal(resultData)

	for captchaResponse.Solution == "" {
		request, err := http.Post(getTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return &CaptchaResponse{}, err
		}

		response, err := internal.ReadResponseBody(request)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		responseAsJSON := antiCaptchaResultResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		if responseAsJSON.ErrorID != 0 {
			return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
		}

		if responseAsJSON.Status == "ready" {
			captchaResponse.Solution = responseAsJSON.Solution.RecaptchaResponse
			return captchaResponse, nil
		}

		tries++
		if tries == payload.MaxRetries {
			return &CaptchaResponse{}, errors.New("captcha took longer than 115 seconds to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return &CaptchaResponse{}, errors.New("reached end of function")
}

func antiCaptchaSolveImageCaptcha(payload *ImageCaptchaPayload) (*CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "api.anti-captcha.com"
	}

	captchaResponse := payload.CreateImageCaptchaResponse()

	createTaskUrl := fmt.Sprintf("%v://%v/createTask", protocol, payload.CustomServiceUrl)
	getTaskUrl := fmt.Sprintf("%v://%v/getTaskResult", protocol, payload.CustomServiceUrl)

	typeTask := map[string]interface{}{
		"type": "ImageToTextTask",
		"body": payload.Base64String,
		"case": payload.CaseSensitive,
	}

	createTask := map[string]interface{}{"clientKey": payload.ServiceApiKey, "task": typeTask}
	jsonValue, _ := json.Marshal(createTask)

	request, err := http.Post(createTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return &CaptchaResponse{}, err
	}

	response, err := internal.ReadResponseBody(request)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	responseAsJSON := antiCaptchaCreateResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	if responseAsJSON.ErrorID == 0 {
		captchaResponse.TaskId = strconv.Itoa(responseAsJSON.TaskID)
	} else {
		return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	resultData := map[string]string{"clientKey": payload.ServiceApiKey, "taskId": fmt.Sprint(captchaResponse.TaskId)}
	jsonValue, _ = json.Marshal(resultData)

	for captchaResponse.Solution == "" {
		request, err := http.Post(getTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return &CaptchaResponse{}, err
		}

		response, err := internal.ReadResponseBody(request)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		responseAsJSON := antiCaptchaResultResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		if responseAsJSON.ErrorID != 0 {
			return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
		}

		if responseAsJSON.Status == "ready" {
			captchaResponse.Solution = responseAsJSON.Solution.Text
			return captchaResponse, nil
		}

		tries++
		if tries == payload.MaxRetries {
			return &CaptchaResponse{}, errors.New("captcha took longer than 115 seconds to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return &CaptchaResponse{}, errors.New("reached end of function")
}

func antiCaptchaSolveHCaptcha(payload *HCaptchaPayload) (*CaptchaResponse, error) {
	var tries int

	protocol := "http"
	if payload.CustomServiceUrl == "" {
		protocol = "https"
		payload.CustomServiceUrl = "api.anti-captcha.com"
	}

	captchaResponse := payload.CreateHCaptchaResponse()

	createTaskUrl := fmt.Sprintf("%v://%v/createTask", protocol, payload.CustomServiceUrl)
	getTaskUrl := fmt.Sprintf("%v://%v/getTaskResult", protocol, payload.CustomServiceUrl)

	typeTask := map[string]interface{}{
		"type":       "HCaptchaTaskProxyless",
		"websiteURL": payload.EndpointUrl,
		"websiteKey": payload.EndpointKey,
	}

	createTask := map[string]interface{}{"clientKey": payload.ServiceApiKey, "task": typeTask}
	jsonValue, _ := json.Marshal(createTask)

	request, err := http.Post(createTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return &CaptchaResponse{}, err
	}

	response, err := internal.ReadResponseBody(request)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	responseAsJSON := antiCaptchaCreateResponse{}
	err = json.Unmarshal([]byte(response), &responseAsJSON)
	if err != nil {
		return &CaptchaResponse{}, err
	}

	if responseAsJSON.ErrorID == 0 {
		captchaResponse.TaskId = strconv.Itoa(responseAsJSON.TaskID)
	} else {
		return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
	}

	time.Sleep(time.Duration(payload.InitialWaitTime) * time.Second)

	resultData := map[string]string{"clientKey": payload.ServiceApiKey, "taskId": fmt.Sprint(captchaResponse.TaskId)}
	jsonValue, _ = json.Marshal(resultData)

	for captchaResponse.Solution == "" {
		request, err := http.Post(getTaskUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return &CaptchaResponse{}, err
		}

		response, err := internal.ReadResponseBody(request)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		responseAsJSON := antiCaptchaResultResponse{}
		err = json.Unmarshal([]byte(response), &responseAsJSON)
		if err != nil {
			return &CaptchaResponse{}, err
		}

		if responseAsJSON.ErrorID != 0 {
			return &CaptchaResponse{}, errors.New(responseAsJSON.ErrorDescription)
		}

		if responseAsJSON.Status == "ready" {
			captchaResponse.Solution = responseAsJSON.Solution.RecaptchaResponse
			return captchaResponse, nil
		}

		tries++
		if tries == payload.MaxRetries {
			return &CaptchaResponse{}, errors.New("captcha took longer than 115 seconds to solve")
		}
		time.Sleep(time.Duration(payload.PollInterval) * time.Second)
	}

	return &CaptchaResponse{}, errors.New("reached end of function")
}

type antiCaptchaCreateResponse struct {
	ErrorID          int    `json:"errorId"`
	ErrorDescription string `json:"errorDescription"`
	TaskID           int    `json:"taskId"`
}

type antiCaptchaResultResponse struct {
	Status           string          `json:"status"`
	ErrorID          int             `json:"errorId"`
	ErrorDescription string          `json:"errorDescription"`
	Solution         antiCapSolution `json:"solution"`
}

type antiCapSolution struct {
	RecaptchaResponse string `json:"gRecaptchaResponse"`
	Text              string `json:"text"`
}
