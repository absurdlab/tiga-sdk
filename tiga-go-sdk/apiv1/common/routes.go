package commonv1

import (
	"fmt"
	"strings"
)

// Routes helps configure API endpoints.
var Routes = route{}

type route struct{}

func (_ route) ResumeAuthorizeURL(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/resume", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func trimTrailingSlash(url string) string {
	return strings.TrimSuffix(url, "/")
}
