package opsgenieIntegrationManager

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"bitbucket.org/atlassianlabs/restclient"
	"github.com/atlassian/voyager/pkg/util"
	"github.com/atlassian/voyager/pkg/util/httputil"
	"github.com/atlassian/voyager/pkg/util/pkiutil"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	asapAudience     = "micros-server"
	asapSubject      = ""
	integrationsPath = "/api/v1/opsgenie/integrations"
)

type Client struct {
	logger     *zap.Logger
	httpClient *http.Client
	asap       pkiutil.ASAP
	rm         *restclient.RequestMutator
}

func New(logger *zap.Logger, httpClient *http.Client, asap pkiutil.ASAP, baseURL *url.URL) *Client {
	rm := restclient.NewRequestMutator(
		restclient.BaseURL(baseURL.String()),
	)
	return &Client{
		logger:     logger,
		httpClient: httpClient,
		asap:       asap,
		rm:         rm,
	}
}

// Gets OpsGenie integrations
// return codes:
// - 201: Service created
// - 400: Bad request to Opsgenie
// - 401: Unauthorized
// - 404: Not found returned by Opsgenie. Does the specified Opsgenie team exist?
// - 422: Semantic error in request to Opsgenie.
// - 429: Rate limited by Opsgenie.
// - 500: Internal server error
func (c *Client) GetOrCreateIntegrations(ctx context.Context, teamName string) (bool /* retriable */, *IntegrationsResponse, error) {
	req, err := c.rm.NewRequest(
		pkiutil.AuthenticateWithASAP(c.asap, asapAudience, asapSubject),
		restclient.Method(http.MethodGet),
		restclient.JoinPath(fmt.Sprintf("%s/%s", integrationsPath, teamName)),
		restclient.Context(ctx),
	)
	if err != nil {
		return false, nil, errors.Wrap(err, "failed to create get integrations request")
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return false, nil, errors.Wrap(err, "failed to execute get integrations request")
	}

	retriable := false
	switch response.StatusCode {
	case http.StatusInternalServerError:
		retriable = true
	case 429:
		retriable = true
	}

	defer util.CloseSilently(response.Body)
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return retriable, nil, errors.Wrap(err, "failed to read response body")
	}

	if response.StatusCode != http.StatusOK {
		message := fmt.Sprintf("failed to get integrations for team %q. Response: %s", teamName, respBody)
		return retriable, nil, clientError(response.StatusCode, message)
	}

	var parsedBody IntegrationsResponse
	err = json.Unmarshal(respBody, &parsedBody)
	if err != nil {
		return retriable, nil, errors.Wrap(err, "failed to unmarshal response body")
	}

	return retriable, &parsedBody, nil
}

func clientError(statusCode int, message string) error {
	switch statusCode {
	case http.StatusNotFound:
		return httputil.NewNotFound(message)
	case http.StatusBadRequest:
		return httputil.NewBadRequest(message)
	default:
		return httputil.NewUnknown(fmt.Sprintf("%s (%s)", message, http.StatusText(statusCode)))
	}
}
