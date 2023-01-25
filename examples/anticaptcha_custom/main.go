package main

import (
	"context"
	"fmt"
	"github.com/justhyped/gocaptcha"
	"net/http"
	"time"
)

func main() {
	ctx := context.TODO()
	cs := gocaptcha.NewCaptchaSolver(gocaptcha.NewCustomAntiCaptcha("https://api.capmonster.cloud", "key"))

	// optional changes
	cs.SetPollInterval(time.Second * 10)
	cs.SetClient(&http.Client{})
	cs.SetInitialWaitTime(time.Second * 10)
	cs.SetMaxRetries(5)

	img, err := cs.SolveImageCaptcha(ctx, &gocaptcha.ImageCaptchaPayload{
		Base64String: "base64here",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(img.Solution())

	if err := img.ReportBad(ctx); err != nil {
		panic(err)
	}
	if err := img.ReportGood(ctx); err != nil {
		panic(err)
	}
}
