package congress

import (
	"net/http"
	"net/url"
	"os"

	"github.com/opnscty/go-httpclient/gohttp"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.propublica.com",
	Path:   "congress/v1",
}

// [TODO] Can't return pointer here??
func Client() gohttp.Client {
	headers := make(http.Header)
	headers.Set("X-API-KEY", os.Getenv("PROPUBLICA_CONGRESS_API_KEY"))

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetUserAgent(os.Getenv("USER_AGENT")).
		Build()

	return client
}
