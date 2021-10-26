package logoutv1

import (
	"fmt"
	"strings"
)

// Routes helps configure API endpoints.
var Routes = route{}

type route struct{}

func (_ route) GetLogoutRequestEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/logout/session/%s", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) ReplyLogoutSuccessEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/logout/session/%s/success", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) ResumeLogoutURL(baseUrl string, logoutSessionId string) string {
	return fmt.Sprintf("%s/oauth/logout/session/%s/resume", trimTrailingSlash(baseUrl), logoutSessionId)
}

func trimTrailingSlash(url string) string {
	return strings.TrimSuffix(url, "/")
}
