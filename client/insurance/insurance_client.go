// Code generated by go-swagger; DO NOT EDIT.

package insurance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new insurance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for insurance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	InsuranceAddV1FvEmea(params *InsuranceAddV1FvEmeaParams) (*InsuranceAddV1FvEmeaOK, error)

	InsuranceInquiryV2FvEmea(params *InsuranceInquiryV2FvEmeaParams) (*InsuranceInquiryV2FvEmeaOK, error)

	InsuranceStatusUpdateV1FvEmea(params *InsuranceStatusUpdateV1FvEmeaParams) (*InsuranceStatusUpdateV1FvEmeaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InsuranceAddV1FvEmea insurances add

  The Add Insurance Service provides the ability to Add insurance products for an account. <BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) InsuranceAddV1FvEmea(params *InsuranceAddV1FvEmeaParams) (*InsuranceAddV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInsuranceAddV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "insuranceAdd_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/insuranceAdd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InsuranceAddV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InsuranceAddV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for insuranceAdd_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InsuranceInquiryV2FvEmea insurances inquiry

  The Insurance Enquiry service provides Insurance related information for a specific account. This service provides details such as Insurance Product, Premium, Status, Reason for cancellation etc. <BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) InsuranceInquiryV2FvEmea(params *InsuranceInquiryV2FvEmeaParams) (*InsuranceInquiryV2FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInsuranceInquiryV2FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "insuranceInquiry_v2_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v2/insuranceInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InsuranceInquiryV2FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InsuranceInquiryV2FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for insuranceInquiry_v2_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InsuranceStatusUpdateV1FvEmea insurances status update

  The Change Insurance Service allows updates of insurance products for an account. This service will emulate functionality used to update the status of an insurance product for inclusion or exclusion of coverage. A FOREIGN USE indicator will be provided in the input message to allow the requesting source to specify whether the status update applies to the local or foreign account when the account is in a dual currency processing organization. The default will be to the local account.  If the account is in a dual currency, each account must be updated individually.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) InsuranceStatusUpdateV1FvEmea(params *InsuranceStatusUpdateV1FvEmeaParams) (*InsuranceStatusUpdateV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInsuranceStatusUpdateV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "insuranceStatusUpdate_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/insuranceStatusUpdate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InsuranceStatusUpdateV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InsuranceStatusUpdateV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for insuranceStatusUpdate_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
