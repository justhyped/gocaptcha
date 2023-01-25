package gocaptcha

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/justhyped/gocaptcha/internal"
	"io"
	"net/http"
	"strconv"
)

type AntiCaptcha struct {
	baseUrl, apiKey string
}

func NewAntiCaptcha(apiKey string) *AntiCaptcha {
	return &AntiCaptcha{
		apiKey:  apiKey,
		baseUrl: "https://api.anti-captcha.com",
	}
}

// NewCustomAntiCaptcha can be used to change the baseUrl, some providers such as CapMonster, XEVil and CapSolver
// have the exact same API as AntiCaptcha, thus allowing you to use these providers with ease.
func NewCustomAntiCaptcha(baseUrl, apiKey string) *AntiCaptcha {
	return &AntiCaptcha{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

func (a *AntiCaptcha) SolveImageCaptcha(ctx context.Context, settings *Settings, payload *ImageCaptchaPayload) (ICaptchaResponse, error) {
	task := map[string]any{
		"type": "ImageToTextTask",
		"body": payload.Base64String,
		"case": payload.CaseSensitive,
	}

	result, err := a.solveTask(ctx, settings, task)
	if err != nil {
		return nil, err
	}

	result.reportBad = a.report("/reportIncorrectImageCaptcha", result.taskId, settings)
	return result, nil
}

func (a *AntiCaptcha) SolveRecaptchaV2(ctx context.Context, settings *Settings, payload *RecaptchaV2Payload) (ICaptchaResponse, error) {
	task := map[string]any{
		"type":        "NoCaptchaTaskProxyless",
		"websiteURL":  payload.EndpointUrl,
		"websiteKey":  payload.EndpointKey,
		"isInvisible": payload.IsInvisibleCaptcha}

	result, err := a.solveTask(ctx, settings, task)
	if err != nil {
		return nil, err
	}

	result.reportBad = a.report("/reportIncorrectRecaptcha", result.taskId, settings)
	result.reportGood = a.report("/reportCorrectRecaptcha", result.taskId, settings)
	return result, nil
}

func (a *AntiCaptcha) SolveRecaptchaV3(ctx context.Context, settings *Settings, payload *RecaptchaV3Payload) (ICaptchaResponse, error) {
	task := map[string]any{
		"type":       "RecaptchaV3TaskProxyless",
		"websiteURL": payload.EndpointUrl,
		"websiteKey": payload.EndpointKey,
		"minScore":   payload.MinScore,
	}

	result, err := a.solveTask(ctx, settings, task)
	if err != nil {
		return nil, err
	}

	result.reportBad = a.report("/reportIncorrectRecaptcha", result.taskId, settings)
	result.reportGood = a.report("/reportCorrectRecaptcha", result.taskId, settings)
	return result, nil
}

func (a *AntiCaptcha) SolveHCaptcha(ctx context.Context, settings *Settings, payload *HCaptchaPayload) (ICaptchaResponse, error) {
	task := map[string]any{
		"type":       "HCaptchaTaskProxyless",
		"websiteURL": payload.EndpointUrl,
		"websiteKey": payload.EndpointKey,
	}

	result, err := a.solveTask(ctx, settings, task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AntiCaptcha) SolveTurnstile(ctx context.Context, settings *Settings, payload *TurnstilePayload) (ICaptchaResponse, error) {
	task := map[string]any{
		"type":       "TurnstileTaskProxyless",
		"websiteURL": payload.EndpointUrl,
		"websiteKey": payload.EndpointKey,
	}

	result, err := a.solveTask(ctx, settings, task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AntiCaptcha) solveTask(ctx context.Context, settings *Settings, task map[string]any) (*CaptchaResponse, error) {
	taskId, err := a.createTask(ctx, settings, task)
	if err != nil {
		return nil, err
	}

	if err := internal.SleepContext(ctx, settings.initialWaitTime); err != nil {
		return nil, err
	}

	for i := 0; i < settings.maxRetries; i++ {
		answer, err := a.getTaskResult(ctx, settings, taskId)
		if err != nil {
			return nil, err
		}

		if answer != "" {
			return &CaptchaResponse{solution: answer, taskId: taskId}, nil
		}

		if err := internal.SleepContext(ctx, settings.pollInterval); err != nil {
			return nil, err
		}
	}

	return nil, errors.New("max tries exceeded")
}

func (a *AntiCaptcha) createTask(ctx context.Context, settings *Settings, task map[string]any) (string, error) {
	type antiCaptchaCreateResponse struct {
		ErrorID          int    `json:"errorId"`
		ErrorDescription string `json:"errorDescription"`
		TaskID           int    `json:"taskId"`
	}

	jsonValue, err := json.Marshal(map[string]any{"clientKey": a.apiKey, "task": task})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseUrl+"/createTask", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := settings.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseAsJSON antiCaptchaCreateResponse
	err = json.Unmarshal(respBody, &responseAsJSON)
	if err != nil {
		return "", err
	}

	if responseAsJSON.ErrorID == 0 {
		return strconv.Itoa(responseAsJSON.TaskID), nil
	} else {
		return "", errors.New(responseAsJSON.ErrorDescription)
	}
}

func (a *AntiCaptcha) getTaskResult(ctx context.Context, settings *Settings, taskId string) (string, error) {
	type antiCapSolution struct {
		RecaptchaResponse string `json:"gRecaptchaResponse"`
		Text              string `json:"text"`
	}

	type resultResponse struct {
		Status           string          `json:"status"`
		ErrorID          int             `json:"errorId"`
		ErrorDescription string          `json:"errorDescription"`
		Solution         antiCapSolution `json:"solution"`
	}

	resultData := map[string]string{"clientKey": a.apiKey, "taskId": taskId}
	jsonValue, err := json.Marshal(resultData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.baseUrl+"/getTaskResult", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	resp, err := settings.client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respJson resultResponse
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		return "", err
	}

	if respJson.ErrorID != 0 {
		return "", errors.New(respJson.ErrorDescription)
	}

	if respJson.Status == "ready" {
		if respJson.Solution.Text != "" {
			return respJson.Solution.Text, nil
		}

		if respJson.Solution.RecaptchaResponse != "" {
			return respJson.Solution.RecaptchaResponse, nil
		}
	}

	return "", nil
}

func (a *AntiCaptcha) report(path, taskId string, settings *Settings) func(ctx context.Context) error {
	type response struct {
		ErrorID          int64  `json:"errorId"`
		ErrorCode        string `json:"errorCode"`
		ErrorDescription string `json:"errorDescription"`
	}

	return func(ctx context.Context) error {
		payload := map[string]string{
			"clientKey": a.apiKey,
			"taskId":    taskId,
		}
		rawPayload, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseUrl+path, bytes.NewBuffer(rawPayload))
		if err != nil {
			return err
		}

		resp, err := settings.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var respJson response
		if err := json.Unmarshal(respBody, &respJson); err != nil {
			return err
		}

		if respJson.ErrorID != 0 {
			return fmt.Errorf("%v: %v", respJson.ErrorCode, respJson.ErrorDescription)
		}

		return nil
	}
}

type CaptchaResponse struct {
	solution, taskId  string
	isAlreadyReported bool

	// these are set internally for each captcha type
	reportGood func(ctx context.Context) error
	reportBad  func(ctx context.Context) error
}

func (a CaptchaResponse) Solution() string {
	return a.solution
}

func (a CaptchaResponse) ReportBad(ctx context.Context) error {
	if a.isAlreadyReported {
		return errors.New("already reported")
	}

	if a.reportBad == nil {
		return errors.New("not implemented for this captcha type")
	}

	return a.reportBad(ctx)
}

func (a CaptchaResponse) ReportGood(ctx context.Context) error {
	if a.isAlreadyReported {
		return errors.New("already reported")
	}

	if a.reportGood == nil {
		return errors.New("not implemented for this captcha type")
	}

	return a.reportGood(ctx)
}

var _ IProvider = (*AntiCaptcha)(nil)
var _ ICaptchaResponse = (*CaptchaResponse)(nil)
