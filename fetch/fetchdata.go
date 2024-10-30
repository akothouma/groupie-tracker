package fetch

import (
	"fmt"
	"io"
	"net/http"
)

// FetchData takes a url and returns its response body, and an error if there's any
func Fetch(url string) ([]byte, error) {
	body := []byte{}
	var body_err error

	response, artists_err := http.Get(url)
	if artists_err != nil {
		return nil, fmt.Errorf("Error making a get request to the artists api endpoint: %s", artists_err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, body_err = io.ReadAll(response.Body)
		if body_err != nil {
			return nil, fmt.Errorf("Error reading response body: %s", body_err)
		}
	}

	return body, nil
}