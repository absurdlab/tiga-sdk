package loginv1

import (
	"fmt"
	"strings"
)

// Routes helps configure API endpoints.
var Routes = route{}

type route struct{}

func (_ route) GetAuthenticationRequestEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/login", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) ReplyAuthenticationSuccessEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/login/success", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) ReplyAuthenticationFailureEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/login/failure", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func trimTrailingSlash(url string) string {
	return strings.TrimSuffix(url, "/")
}
