// Generated code
// run `make generate` to update

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/atlassian/voyager/pkg/apis/composition/v1"
	"github.com/atlassian/voyager/pkg/composition/client/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CompositionV1Interface interface {
	RESTClient() rest.Interface
	ServiceDescriptorsGetter
}

// CompositionV1Client is used to interact with features provided by the composition.voyager.atl-paas.net group.
type CompositionV1Client struct {
	restClient rest.Interface
}

func (c *CompositionV1Client) ServiceDescriptors() ServiceDescriptorInterface {
	return newServiceDescriptors(c)
}

// NewForConfig creates a new CompositionV1Client for the given config.
func NewForConfig(c *rest.Config) (*CompositionV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CompositionV1Client{client}, nil
}

// NewForConfigOrDie creates a new CompositionV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CompositionV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CompositionV1Client for the given RESTClient.
func New(c rest.Interface) *CompositionV1Client {
	return &CompositionV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CompositionV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
