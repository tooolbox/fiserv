// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CollectionsAccountInquiryV1FvEmea(params *CollectionsAccountInquiryV1FvEmeaParams) (*CollectionsAccountInquiryV1FvEmeaOK, error)

	CollectionsPromiseToPayAddV1FvEmea(params *CollectionsPromiseToPayAddV1FvEmeaParams) (*CollectionsPromiseToPayAddV1FvEmeaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CollectionsAccountInquiryV1FvEmea collections account inquiry

  The CTA Account Information Inquiry service is used to retrieve data held at the account level for a FirstVision account. This service can be called with Collection Organization and Account Number to access the correct records in CTA and CMS. If called with a card number, details will be returned for the account number to which the card number is linked.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) CollectionsAccountInquiryV1FvEmea(params *CollectionsAccountInquiryV1FvEmeaParams) (*CollectionsAccountInquiryV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCollectionsAccountInquiryV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "collectionsAccountInquiry_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/collectionsAccountInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CollectionsAccountInquiryV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CollectionsAccountInquiryV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for collectionsAccountInquiry_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CollectionsPromiseToPayAddV1FvEmea collections promise to pay add

  The CTA Promise to Pay update service is used to capture the customer data when customer promises to pay by a certain date. This date is held at the account level in FirstVision. This service can be called with Collection Organization and Account Number to access the correct records in CTA and CMS. If called with a card number, details will be returned for the account number to which the card number is linked. <BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) CollectionsPromiseToPayAddV1FvEmea(params *CollectionsPromiseToPayAddV1FvEmeaParams) (*CollectionsPromiseToPayAddV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCollectionsPromiseToPayAddV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "collectionsPromiseToPayAdd_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/collectionsPromiseToPayAdd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CollectionsPromiseToPayAddV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CollectionsPromiseToPayAddV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for collectionsPromiseToPayAdd_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
