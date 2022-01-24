package helpers

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"

	"github.com/dsnet/compress/brotli"
)

func BooleanToString(b bool) string {
	if b {
		return "1"
	}

	return "0"
}

func ReadResponseBody(response *http.Response) (string, error) {
	var reader io.ReadCloser
	var err error
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
	case "br":
		reader, err = brotli.NewReader(response.Body, nil)
	default:
		reader = response.Body
	}
	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(response.Body)

	return string(respBody), nil
}
