package gocaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/justhyped/gocaptcha/internal"
	"net/http"
	"net/url"
)

func (response *CaptchaResponse) ReportGoodRecaptcha() {
	switch response.Service {
	case "2Captcha":
		response.reportRecaptchaTwoCaptcha("reportgood")
	case "AntiCaptcha":
		response.reportRecaptchaAntiCaptcha("reportCorrectRecaptcha")
	}
}

func (response *CaptchaResponse) ReportBadRecaptcha() {
	switch response.Service {
	case "2Captcha":
		response.reportRecaptchaTwoCaptcha("reportbad")
	case "AntiCaptcha":
		response.reportRecaptchaAntiCaptcha("reportIncorrectRecaptcha")
	}
}

func (response *CaptchaResponse) reportRecaptchaTwoCaptcha(goodOrBad string) {
	protocol := "http"
	if response.Endpoint == "" {
		protocol = "https"
		response.Endpoint = "2captcha.com"
	}

	urlParams := url.Values{}
	urlParams.Set("key", response.ServiceApiKey)
	urlParams.Set("action", goodOrBad)
	urlParams.Set("id", response.TaskId)

	request, err := http.Get(fmt.Sprint(protocol, "://", response.Endpoint, "/res.php?", urlParams.Encode()))
	if err != nil {
		return
	}

	_, _ = internal.ReadResponseBody(request)
}

func (response *CaptchaResponse) reportRecaptchaAntiCaptcha(goodOrBad string) {
	protocol := "http"
	if response.Endpoint == "" {
		protocol = "https"
		response.Endpoint = "api.anti-captcha.com"
	}

	createTask := map[string]interface{}{"clientKey": response.ServiceApiKey, "taskId": response.TaskId}
	jsonValue, _ := json.Marshal(createTask)

	request, err := http.Post(fmt.Sprint(protocol, "://", response.Endpoint, "/", goodOrBad),
		"application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}

	_, _ = internal.ReadResponseBody(request)
}
