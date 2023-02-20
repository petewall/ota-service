package lib

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Host string
}

func (c *Client) Update(mac, firmware, version string) ([]byte, error) {
	updateURL, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse host (%s): %w", c.Host, err)
	}

	updateURL.Path = "/update"
	params := updateURL.Query()
	params.Set("firmware", firmware)
	params.Set("version", version)
	updateURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, updateURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build update request: %w", err)
	}

	req.Header.Set("x-esp8266-sta-mac", mac)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send update request: %w", err)
	}

	if res.StatusCode == http.StatusNotModified {
		return nil, nil
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update request failed: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read firmware binary: %w", err)
	}

	return body, nil
}
