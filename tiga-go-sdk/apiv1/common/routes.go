package commonv1

import (
	"fmt"
	"strings"
)

// Routes helps configure API endpoints.
var Routes = route{}

type route struct{}

func (_ route) AuthorizeURL(baseUrl string) string {
	return fmt.Sprintf("%s/oauth/authorize", trimTrailingSlash(baseUrl))
}

func (_ route) ResumeAuthorizeURL(baseUrl string, authorizeSessionId string) string {
	return fmt.Sprintf("%s/oauth/authorize/session/%s/resume", trimTrailingSlash(baseUrl), authorizeSessionId)
}

func (_ route) TokenURL(baseUrl string) string {
	return fmt.Sprintf("%s/oauth/token", trimTrailingSlash(baseUrl))
}

func (_ route) TokenRevokeURL(baseUrl string) string {
	return fmt.Sprintf("%s/oauth/token/revoke", trimTrailingSlash(baseUrl))
}

func (_ route) UserInfoURL(baseUrl string) string {
	return fmt.Sprintf("%s/oauth/token/userinfo", trimTrailingSlash(baseUrl))
}

func (_ route) DiscoveryURL(baseUrl string) string {
	return fmt.Sprintf("%s/.well-known/openid-configuration", trimTrailingSlash(baseUrl))
}

func (_ route) JSONWebKeySetURL(baseUrl string) string {
	return fmt.Sprintf("%s/.well-known/jwks.json", trimTrailingSlash(baseUrl))
}

func trimTrailingSlash(url string) string {
	return strings.TrimSuffix(url, "/")
}
