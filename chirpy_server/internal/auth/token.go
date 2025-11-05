package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	bearerString := headers.Get("Authorization")
	if bearerString == "" {
		return "", fmt.Errorf("auth header not found")
	}
	token := strings.TrimPrefix(bearerString, "Bearer ")
	return token, nil
}