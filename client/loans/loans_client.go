// Code generated by go-swagger; DO NOT EDIT.

package loans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new loans API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for loans API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AccountPlanlistV1FvEmea(params *AccountPlanlistV1FvEmeaParams) (*AccountPlanlistV1FvEmeaOK, error)

	BookOfferV1FvEmea(params *BookOfferV1FvEmeaParams) (*BookOfferV1FvEmeaOK, error)

	EppCancelV1FvEmea(params *EppCancelV1FvEmeaParams) (*EppCancelV1FvEmeaOK, error)

	EppTransactionPlanAddV1FvEmea(params *EppTransactionPlanAddV1FvEmeaParams) (*EppTransactionPlanAddV1FvEmeaOK, error)

	EppValidationV1FvEmea(params *EppValidationV1FvEmeaParams) (*EppValidationV1FvEmeaOK, error)

	FppAddV1FvEmea(params *FppAddV1FvEmeaParams) (*FppAddV1FvEmeaOK, error)

	FppQuoteGenerationV1FvEmea(params *FppQuoteGenerationV1FvEmeaParams) (*FppQuoteGenerationV1FvEmeaOK, error)

	LoanPlanAddV2FvEmea(params *LoanPlanAddV2FvEmeaParams) (*LoanPlanAddV2FvEmeaOK, error)

	LoanPlanInquiryV1FvEmea(params *LoanPlanInquiryV1FvEmeaParams) (*LoanPlanInquiryV1FvEmeaOK, error)

	PlanSettlementQuoteUpdateV1FvEmea(params *PlanSettlementQuoteUpdateV1FvEmeaParams) (*PlanSettlementQuoteUpdateV1FvEmeaOK, error)

	SettlementQuoteInquiryV1FvEmea(params *SettlementQuoteInquiryV1FvEmeaParams) (*SettlementQuoteInquiryV1FvEmeaOK, error)

	SettlementQuoteUpdateV1FvEmea(params *SettlementQuoteUpdateV1FvEmeaParams) (*SettlementQuoteUpdateV1FvEmeaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AccountPlanlistV1FvEmea accounts planlist

  This service is used to retrieve information on the available plans on an account. If the account has been transferred, the current account number must be provided. This service does not automatically switch to the new account.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) AccountPlanlistV1FvEmea(params *AccountPlanlistV1FvEmeaParams) (*AccountPlanlistV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountPlanlistV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "accountPlanlist_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/accountPlanlist",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountPlanlistV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountPlanlistV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accountPlanlist_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BookOfferV1FvEmea books offer

  Book Offer service is used for a client or 3rd party to book an offer in the system when the customer accepts their EPP offer that was send to the customer. This service will update the EPP plan in the Vision authorisation files.
*/
func (a *Client) BookOfferV1FvEmea(params *BookOfferV1FvEmeaParams) (*BookOfferV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBookOfferV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bookOffer_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/bookOffer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BookOfferV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BookOfferV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bookOffer_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EppCancelV1FvEmea epps cancel

  The Equal Payment Plan Close or Cancel service enables a customer  service representative to close or cancel an equal payment plan.
*/
func (a *Client) EppCancelV1FvEmea(params *EppCancelV1FvEmeaParams) (*EppCancelV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEppCancelV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "eppCancel_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/eppCancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EppCancelV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EppCancelV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for eppCancel_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EppTransactionPlanAddV1FvEmea epps transaction plan add

  The Equal Payment Plan Add service enables a Customer Service Representative to convert a sales transaction or a plan balance to an equal payment plan.
*/
func (a *Client) EppTransactionPlanAddV1FvEmea(params *EppTransactionPlanAddV1FvEmeaParams) (*EppTransactionPlanAddV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEppTransactionPlanAddV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "eppTransactionPlanAdd_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/eppTransactionPlanAdd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EppTransactionPlanAddV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EppTransactionPlanAddV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for eppTransactionPlanAdd_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EppValidationV1FvEmea epps validation

  The Equal Payment Plan Validation service validates the transaction or plan balance EPP conversion submitted by the Customer Service Representative.
*/
func (a *Client) EppValidationV1FvEmea(params *EppValidationV1FvEmeaParams) (*EppValidationV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEppValidationV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "eppValidation_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/eppValidation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EppValidationV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EppValidationV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for eppValidation_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FppAddV1FvEmea fpps add

  The Flexible Payment Processing Add or Update service allows a client to add or cancel the flexible payment processing (FPP) request for credit plan segments.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) FppAddV1FvEmea(params *FppAddV1FvEmeaParams) (*FppAddV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFppAddV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fppAdd_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/fppAdd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FppAddV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FppAddV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fppAdd_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FppQuoteGenerationV1FvEmea fpps quote generation

  Flexible Payment Plans FPP offers cardholders more control over the budgeting and repayment of their credit card account.  Cardholders are able to select transaction or balance amounts that are moved to a Flexible Payment Plan and select a payment amount to repay that portion of their balance over a period of time. A real time message returns a quote for fixed payments on a plan. The message invokes a service processing which calculates an estimated monthly fixed payment, which will pay down a fixed amount, including all accrued interest over a fixed period taking into account compounded daily interest and a period of time between initial plan debit and first account cycle.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) FppQuoteGenerationV1FvEmea(params *FppQuoteGenerationV1FvEmeaParams) (*FppQuoteGenerationV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFppQuoteGenerationV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fppQuoteGeneration_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/fppQuoteGeneration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FppQuoteGenerationV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FppQuoteGenerationV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fppQuoteGeneration_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LoanPlanAddV2FvEmea loans plan add

  The purpose of this service is to build loan plan data records. The service creates 4 loan plan types and the loan schedule records. The following records are created by this service Plan Segment Data, User Input Data, Amortization Input Data and Commission Subsidy plus the optional Loan Schedule record. Note a Loan Schedule record cannot be provided without a Plan Add record. The service will return a fail indicator F in the response message with an appropriate error code and description for records sent in error. The loan schedule record will be built automatically for ALOP, if Plan number, Term and Principal amount are given.
*/
func (a *Client) LoanPlanAddV2FvEmea(params *LoanPlanAddV2FvEmeaParams) (*LoanPlanAddV2FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoanPlanAddV2FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loanPlanAdd_v2_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v2/loanPlanAdd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoanPlanAddV2FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoanPlanAddV2FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for loanPlanAdd_v2_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LoanPlanInquiryV1FvEmea loans plan inquiry

  The CMS loan overview inquiry service is used to provide the loan data for an account from the plan segment file.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) LoanPlanInquiryV1FvEmea(params *LoanPlanInquiryV1FvEmeaParams) (*LoanPlanInquiryV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoanPlanInquiryV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loanPlanInquiry_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/loanPlanInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoanPlanInquiryV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoanPlanInquiryV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for loanPlanInquiry_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PlanSettlementQuoteUpdateV1FvEmea plans settlement quote update

  This service is used to update the Plan settlement information which enables the user to generate settlement quote for Partial early settlement on a Loan plan associated for an account.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) PlanSettlementQuoteUpdateV1FvEmea(params *PlanSettlementQuoteUpdateV1FvEmeaParams) (*PlanSettlementQuoteUpdateV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlanSettlementQuoteUpdateV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "planSettlementQuoteUpdate_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/planSettlementQuoteUpdate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PlanSettlementQuoteUpdateV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PlanSettlementQuoteUpdateV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for planSettlementQuoteUpdate_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SettlementQuoteInquiryV1FvEmea settlements quote inquiry

  This Service provides information regarding the status of a FirstVision account.<p><p>If an account has been transferred, the current FirstVision account number must be provided, as this service does not automatically switch to the new account. This service can be called with an account number. The service will retrieve settlement quote related information for the account.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) SettlementQuoteInquiryV1FvEmea(params *SettlementQuoteInquiryV1FvEmeaParams) (*SettlementQuoteInquiryV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettlementQuoteInquiryV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "settlementQuoteInquiry_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/settlementQuoteInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SettlementQuoteInquiryV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettlementQuoteInquiryV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for settlementQuoteInquiry_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SettlementQuoteUpdateV1FvEmea settlements quote update

  This service is used to update settlement quotation fields at the account level. The fields such as 'settlement quotation type'<p><p>P – Settlement quotation with penalties for interest, insurance and termination fees<p>W – Settlement quote waiving the interest, insurance and termination fees<p><p>Other fields like Loan pay off date can also be updated via this service.
*/
func (a *Client) SettlementQuoteUpdateV1FvEmea(params *SettlementQuoteUpdateV1FvEmeaParams) (*SettlementQuoteUpdateV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettlementQuoteUpdateV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "settlementQuoteUpdate_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/settlementQuoteUpdate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SettlementQuoteUpdateV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SettlementQuoteUpdateV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for settlementQuoteUpdate_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
