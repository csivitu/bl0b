package ctftime

import (
	"fmt"
	"net/http"
)

// CTFtime struct with BaseURL
type CTFtime struct {
	BaseURL string
}

// New creates a new CTFtime struct with a pre-defined BaseURL
func New() *CTFtime {
	return &CTFtime{
		BaseURL: "https://ctftime.org/api/v1",
	}
}

// Get is used to perform a GET request on a CTFtime endpoint
func (ctf *CTFtime) Get(endpoint string) (*http.Response, error) {
	client := &http.Client{}

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s%s", ctf.BaseURL, endpoint),
		nil,
	)

	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")

	res, err := client.Do(req)

	return res, err
}
