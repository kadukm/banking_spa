package utils

import (
	"net/http"
	"strings"
)

func MIMEContentTypeIsJSON(request *http.Request) bool {
	contentType := request.Header.Get("Content-Type")
	return strings.EqualFold(contentType, "application/json")
}
