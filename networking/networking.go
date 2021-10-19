package networking

import (
	"bytes"
	"net/http"
	"time"
)

func SendSoap(endpoint, message string) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, err
	}

	return resp, nil
}
