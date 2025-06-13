package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authentication info found")
	}

	parts := strings.Split(val, " ")
	if len(parts) != 2  {
		return "", errors.New("invalid authorization format")
	}
	return parts[1], nil
}
