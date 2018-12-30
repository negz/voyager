// Code generated by go-swagger; DO NOT EDIT.

package resolve

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new resolve API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resolve API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Resolve resolves

Resolve a given location set into a set of matching Mapping names to their associated release details expanded.
*/
func (a *Client) Resolve(params *ResolveParams) (*ResolveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResolveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resolve",
		Method:             "GET",
		PathPattern:        "/v1/resolve",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream", "application/x-protobuf", "application/xml", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream", "application/x-protobuf", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResolveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResolveOK), nil

}

/*
ResolveBatch resolves batch

Resolve a given location set into a set of matching Mapping names to their associated release details expanded.
*/
func (a *Client) ResolveBatch(params *ResolveBatchParams) (*ResolveBatchOK, *ResolveBatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResolveBatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resolveBatch",
		Method:             "GET",
		PathPattern:        "/v1/resolve/batch",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream", "application/x-protobuf", "application/xml", "text/html"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream", "application/x-protobuf", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResolveBatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ResolveBatchOK:
		return value, nil, nil
	case *ResolveBatchNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
