package it

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"

	"github.com/atlassian/voyager/pkg/opsgenieIntegrationManager"
	"github.com/atlassian/voyager/pkg/util"
	"github.com/atlassian/voyager/pkg/util/logz"
	"github.com/atlassian/voyager/pkg/util/pkiutil"
	"github.com/atlassian/voyager/pkg/util/testutil"
	"github.com/stretchr/testify/require"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	microsServerURL = "https://micros.prod.atl-paas.net"
)

// You can run this file from root using
// `go test -v pkg/opsgenieIntegrationManager/it/client_manual_test.go`
// NOTE: THIS WILL CREATE INTEGRATIONS IF NONE EXIST
func TestGetIntegrations(t *testing.T) {
	t.Parallel()

	// Prepare ASAP secrets from Kubernetes Secret
	asapCreatorSecret := getSecret(t)
	ctx := testutil.ContextWithLogger(t)
	testLogger := logz.RetrieveLoggerFromContext(ctx)
	asapConfig, asapErr := pkiutil.NewASAPClientConfigFromKubernetesSecret(asapCreatorSecret)
	require.NoError(t, asapErr)

	client := util.HTTPClient()
	c := opsgenieIntegrationManager.New(testLogger, client, asapConfig, parseURL(t, microsServerURL))

	// Get Service Attributes
	_, resp, err := c.GetOrCreateIntegrations(ctx, "Platform SRE")
	require.NoError(t, err)

	testLogger.Sugar().Infof("Number of returned integrations: %v", len(resp.Integrations))
	testLogger.Sugar().Infof("Response: %#v", resp)

	bytes, _ := json.Marshal(resp)
	testLogger.Sugar().Infof("Attributes JSON: %#v", string(bytes))
}

// data should be "export OPSGENIE_YAML=$(kubectl -n voyager get secrets asap-creator -o yaml)"
func getSecret(t *testing.T) *v1.Secret {
	data := os.Getenv("OPSGENIE_YAML") //Envvar containing the yaml contents of the secret

	decode := scheme.Codecs.UniversalDeserializer().Decode
	destination := &v1.Secret{}
	_, _, err := decode([]byte(data), nil, destination)
	require.NoError(t, err)
	return destination
}

func parseURL(t *testing.T, urlstr string) *url.URL {
	urlobj, err := url.Parse(urlstr)
	require.NoError(t, err)
	return urlobj
}
