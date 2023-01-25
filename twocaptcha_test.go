package gocaptcha

import (
	"context"
	"testing"
)

func TestNewTwoCaptcha(t *testing.T) {
	cs := NewCaptchaSolver(NewTwoCaptcha("key"))
	cs.SetPollInterval(1)
	cs.SetInitialWaitTime(1)

	resp, err := cs.SolveRecaptchaV2(context.Background(), &RecaptchaV2Payload{
		EndpointUrl: "https://www.google.com/recaptcha/api2/demo",
		EndpointKey: "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(resp.Solution())
}
