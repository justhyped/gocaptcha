package gocaptcha

import (
	"context"
	"testing"
)

func TestNewCaptchaSolver(t *testing.T) {
	ctx := context.Background()

	cs := NewCaptchaSolver(NewCustomAntiCaptcha("https://api.capmonster.cloud", "key"))

	resp, err := cs.SolveRecaptchaV2(ctx, &RecaptchaV2Payload{
		EndpointUrl: "https://www.google.com/recaptcha/api2/demo",
		EndpointKey: "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(resp.Solution()) // gets the answer or recaptcha token etc
}
