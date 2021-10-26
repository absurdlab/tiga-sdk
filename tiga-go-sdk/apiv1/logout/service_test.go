package logoutv1_test

import (
	"context"
	commonv1 "github.com/absurdlab/tiga-sdk/tiga-go-sdk/apiv1/common"
	logoutv1 "github.com/absurdlab/tiga-sdk/tiga-go-sdk/apiv1/logout"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2/clientcredentials"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

const tigaBaseURL = "https://test.tiga.elan-vision.com"

func TestDefaultService(t *testing.T) {
	suite.Run(t, new(defaultServiceTestSuite))
}

type defaultServiceTestSuite struct {
	suite.Suite
}

func (s *defaultServiceTestSuite) TestGetLogout() {
	service := newService()

	cases := []struct {
		name      string
		sessionId string
		assert    func(t *testing.T, req *logoutv1.LogoutRequest, err error)
	}{
		{
			name:      "ok",
			sessionId: "ok",
			assert: func(t *testing.T, req *logoutv1.LogoutRequest, err error) {
				if assert.NoError(t, err) && assert.NotNil(t, req) {
					assert.Equal(t, "test", req.Client.ClientName)
				}
			},
		},
		{
			name:      "bad",
			sessionId: "bad",
			assert: func(t *testing.T, req *logoutv1.LogoutRequest, err error) {
				var e commonv1.ErrorResponseWrapper
				if assert.ErrorAs(t, err, &e) && assert.Nil(t, req) {
					assert.Equal(t, "bad_request", e.Error())
				}
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			resp, err := service.GetLogout(
				context.Background(),
				logoutv1.Routes.GetLogoutRequestEndpoint(tigaBaseURL, c.sessionId),
			)
			c.assert(t, resp, err)
		})
	}
}

func (s *defaultServiceTestSuite) SetupSuite() {
	httpmock.Activate()
	s.Require().NoError(setupMock())
}

func (s *defaultServiceTestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func newService() logoutv1.Service {
	return logoutv1.NewService(&clientcredentials.Config{
		ClientID:     "test",
		ClientSecret: "s3cret",
		TokenURL:     tigaBaseURL + "/oauth/token",
		Scopes:       []string{"tiga:logout"},
	})
}

func setupMock() error {
	f, err := os.Open("testdata/mock.yaml")
	if err != nil {
		return err
	}

	type record struct {
		URL    string `yaml:"url"`
		Method string `yaml:"method"`
		Status int    `yaml:"status"`
		Body   string `yaml:"body"`
	}

	spec := make([]record, 0)
	if err := yaml.NewDecoder(f).Decode(&spec); err != nil {
		return err
	}

	for _, r := range spec {
		httpmock.RegisterResponder(r.Method, r.URL, httpmock.NewStringResponder(r.Status, r.Body))
	}

	return nil
}
