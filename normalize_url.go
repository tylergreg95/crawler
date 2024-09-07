package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	host := parsedURL.Host
	path := strings.TrimSuffix(parsedURL.Path, "/")

	return host + path, nil
}
