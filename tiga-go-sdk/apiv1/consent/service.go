package consentv1

import (
	"bytes"
	"context"
	"encoding/json"
	commonv1 "github.com/absurdlab/tiga-sdk/tiga-go-sdk/apiv1/common"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// ParamRequest is the name of the request parameter that contains the URL to fetch ConsentRequest.
const ParamRequest = "tiga_consent"

// Service abstracts consent API endpoints.
type Service interface {
	// GetConsent retrieves the ConsentRequest data from Tiga via a URL link provided on tiga during
	// the redirection to the Dyna consent entrypoint.
	GetConsent(ctx context.Context, url string) (*ConsentRequest, error)
	// ReplyConsentSuccess posts back the successful ConsentCallback to Tiga.
	ReplyConsentSuccess(ctx context.Context, req *ConsentRequest, payload *ConsentCallback) error
	// ReplyConsentFailure posts back the failure ErrorCallback to Tiga.
	ReplyConsentFailure(ctx context.Context, req *ConsentRequest, payload *commonv1.ErrorCallback) error
}

type Option func(s *defaultService)

func WithTransport(transport http.RoundTripper) Option {
	return func(s *defaultService) {
		s.transport = transport
	}
}

// NewService returns a new instance of the default implementation of Service.
func NewService(oauth *clientcredentials.Config, options ...Option) Service {
	svc := &defaultService{oauth: oauth}

	for _, opt := range options {
		opt(svc)
	}

	if oauth == nil {
		panic("oauth config is required")
	}

	if svc.transport == nil {
		svc.transport = http.DefaultTransport
	}

	return svc
}

type defaultService struct {
	oauth     *clientcredentials.Config
	transport http.RoundTripper
}

func (s *defaultService) GetConsent(ctx context.Context, url string) (*ConsentRequest, error) {
	var ar ConsentRequest
	if err := s.getRequest(ctx, url, &ar); err != nil {
		return nil, err
	}
	return &ar, nil
}

func (s *defaultService) ReplyConsentSuccess(ctx context.Context, req *ConsentRequest, payload *ConsentCallback) error {
	return s.postCallback(ctx, req.Links.SuccessCallbackUrl, payload)
}

func (s *defaultService) ReplyConsentFailure(ctx context.Context, req *ConsentRequest, payload *commonv1.ErrorCallback) error {
	return s.postCallback(ctx, req.Links.ErrorCallbackUrl, payload)
}

func (s *defaultService) postCallback(ctx context.Context, url string, body interface{}) error {
	raw, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := s.httpClient(ctx).Post(url, "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	var errorResp commonv1.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
		return err
	}

	return commonv1.WrapErrorResponse(&errorResp)
}

func (s *defaultService) getRequest(ctx context.Context, url string, dest interface{}) error {
	resp, err := s.httpClient(ctx).Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&dest); err != nil {
			return err
		}
		return nil
	}

	var errorResp commonv1.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
		return err
	}

	return commonv1.WrapErrorResponse(&errorResp)
}

func (s *defaultService) httpClient(ctx context.Context) *http.Client {
	return &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ReuseTokenSource(nil, s.oauth.TokenSource(ctx)),
			Base:   s.transport,
		},
	}
}
