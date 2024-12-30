package wallhaven

import (
	"net/http"
	"time"
)

const DEFAULT_CLIENT_TIMEOUT = 5 * time.Second

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func DefaultHttpClient() HttpClient {
	client := &http.Client{
		Timeout: DEFAULT_CLIENT_TIMEOUT,
	}
	return client
}
