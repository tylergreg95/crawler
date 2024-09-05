package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	host := parsedURL.Host
	path := parsedURL.Path

	fmt.Println(host + path)

	return host + path, nil
}
