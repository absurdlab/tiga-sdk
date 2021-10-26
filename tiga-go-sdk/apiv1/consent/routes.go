package consentv1

import (
	"fmt"
	"strings"
)

// Routes helps configure API endpoints.
var Routes = route{}

type route struct{}

func (_ route) GetConsentRequestEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/consent", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) ReplyConsentSuccessEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/consent/success", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) ReplyConsentFailureEndpoint(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/consent/failure", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func trimTrailingSlash(url string) string {
	return strings.TrimSuffix(url, "/")
}
