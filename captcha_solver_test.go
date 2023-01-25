package gocaptcha

import (
	"context"
	"net/http"
	"testing"
)

func TestNewCaptchaSolver(t *testing.T) {
	ctx := context.Background()

	cs := NewCaptchaSolver(NewAntiCaptcha("test"))
	cs.SetClient(http.DefaultClient)
	cs.SetInitialWaitTime(10)
	cs.SetMaxRetries(5)
	cs.SetPollInterval(20)

	resp, err := cs.SolveRecaptchaV2(ctx, &RecaptchaV2Payload{
		EndpointUrl: "",
		EndpointKey: "",
	})
	if err != nil {
		t.Error(err)
	}

	resp.Solution() // gets the answer or recaptcha token etc
	_ = resp.ReportBad(ctx)
	_ = resp.ReportGood(ctx)
}
