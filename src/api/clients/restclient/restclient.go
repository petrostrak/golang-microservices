package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	JSONBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(JSONBytes))
	req.Header = headers

	client := http.Client{}
	return client.Do(req)
}
