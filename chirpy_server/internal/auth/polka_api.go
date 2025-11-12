package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	apiString := header.Get("Authorization")
	if apiString == "" {
		return "", fmt.Errorf("auth header not found")
	}
	key := strings.TrimPrefix(apiString, "ApiKey ")
	return key, nil
}