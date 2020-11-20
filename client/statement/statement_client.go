// Code generated by go-swagger; DO NOT EDIT.

package statement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new statement API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for statement API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	StatementDateInquiryV1FvEmea(params *StatementDateInquiryV1FvEmeaParams) (*StatementDateInquiryV1FvEmeaOK, error)

	StatementPdfListV1FvEmea(params *StatementPdfListV1FvEmeaParams) (*StatementPdfListV1FvEmeaOK, error)

	StatementPdfURLV1FvEmea(params *StatementPdfURLV1FvEmeaParams) (*StatementPdfURLV1FvEmeaOK, error)

	StatementReprintUpdateV1FvEmea(params *StatementReprintUpdateV1FvEmeaParams) (*StatementReprintUpdateV1FvEmeaOK, error)

	StatementSummaryInquiryV1FvEmea(params *StatementSummaryInquiryV1FvEmeaParams) (*StatementSummaryInquiryV1FvEmeaOK, error)

	StatementSummaryInquiryV2FvEmea(params *StatementSummaryInquiryV2FvEmeaParams) (*StatementSummaryInquiryV2FvEmeaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StatementDateInquiryV1FvEmea statements date inquiry

  The Statement Date Inquiry Service will allow a user to inquire on a Statement date and its status, noting statement status is either detail or summary.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) StatementDateInquiryV1FvEmea(params *StatementDateInquiryV1FvEmeaParams) (*StatementDateInquiryV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatementDateInquiryV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statementDateInquiry_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/statementDateInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatementDateInquiryV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatementDateInquiryV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statementDateInquiry_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatementPdfListV1FvEmea this service is used to request statement list for an account

  Financial institution(issuer) part of First Vision system can request this service by providing cardholder account number, statement date and type. As a pre-requisite FICS system configuration should exist for the client with necessary level of setup. Based on the type of request, number of internal checks and actions will be performed on the cardholders account and the response will be sent back.
*/
func (a *Client) StatementPdfListV1FvEmea(params *StatementPdfListV1FvEmeaParams) (*StatementPdfListV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatementPdfListV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statementPdfList_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/statementPdfList",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatementPdfListV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatementPdfListV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statementPdfList_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatementPdfURLV1FvEmea this service is used to request statement document u r ls for an account given a start date and statement type

  Financial institution(issuer) part of First Vision system can request this service by providing cardholder account number, statement date and type. As a pre-requisite FICS system configuration should exist for the client with necessary level of setup. Based on the type of request, number of internal checks and actions will be performed on the cardholders account and the response will be sent back.
*/
func (a *Client) StatementPdfURLV1FvEmea(params *StatementPdfURLV1FvEmeaParams) (*StatementPdfURLV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatementPdfURLV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statementPdfUrl_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/statementPdfUrl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatementPdfURLV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatementPdfURLV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statementPdfUrl_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatementReprintUpdateV1FvEmea statements reprint update

  This service will write a memo record to the ASM history file for each statement requested using ASSD as the action code for the memo.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) StatementReprintUpdateV1FvEmea(params *StatementReprintUpdateV1FvEmeaParams) (*StatementReprintUpdateV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatementReprintUpdateV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statementReprintUpdate_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/statementReprintUpdate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatementReprintUpdateV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatementReprintUpdateV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statementReprintUpdate_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatementSummaryInquiryV1FvEmea statements summary inquiry

  This service is deprecated and will be supported until January 2021. The new version contains changes in date formats of certain fields to Gregorian (YYYYMMDD). Please switch to Version 2 of this API before that time. <BR/><BR/>This service is used to retrieve information from the CMS online statement. The response provides a summary of account activity that does not include the individual transaction details. This service can be called with an account number or a card number.<p>The statement inquiry can be requested in two formats, detail or monthly. If the request is sent with detail option then the exact statement date is a mandatory input in request. If the option is monthly then the relative statement number can be passed in the request as per the below.<p><p>00 - Current statement<p>01 - Previous months statement<p>02 - Second previous months statement, etc.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) StatementSummaryInquiryV1FvEmea(params *StatementSummaryInquiryV1FvEmeaParams) (*StatementSummaryInquiryV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatementSummaryInquiryV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statementSummaryInquiry_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/statementSummaryInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatementSummaryInquiryV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatementSummaryInquiryV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statementSummaryInquiry_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatementSummaryInquiryV2FvEmea statements summary inquiry

  The Statement Inquiry service provides information from the CMS online statement. The response provides a summary of account activity that does not include the individual transaction details. This service can be called with an account number or a card number.<p>The statement inquiry can be requested in two formats, detail or monthly. If the request is sent with detail option then the exact statement date is a mandatory input in request. If the option is monthly then the relative statement number can be passed in the request as per the below.<p><p>00 - Current statement<p><p>01 - Previous months statement<p><p>02 - Second previous months statement, etc.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) StatementSummaryInquiryV2FvEmea(params *StatementSummaryInquiryV2FvEmeaParams) (*StatementSummaryInquiryV2FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatementSummaryInquiryV2FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statementSummaryInquiry_v2_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v2/statementSummaryInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatementSummaryInquiryV2FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatementSummaryInquiryV2FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statementSummaryInquiry_v2_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}