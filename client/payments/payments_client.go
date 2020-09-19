// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	FasterPaymentV1FvEmea(params *FasterPaymentV1FvEmeaParams) (*FasterPaymentV1FvEmeaOK, error)

	MoveFundsV1FvEmea(params *MoveFundsV1FvEmeaParams) (*MoveFundsV1FvEmeaOK, error)

	PaymentDueEstimationV1FvEmea(params *PaymentDueEstimationV1FvEmeaParams) (*PaymentDueEstimationV1FvEmeaOK, error)

	PaymentHistoryInquiryV1FvEmea(params *PaymentHistoryInquiryV1FvEmeaParams) (*PaymentHistoryInquiryV1FvEmeaOK, error)

	PaymentSettingsV1FvEmea(params *PaymentSettingsV1FvEmeaParams) (*PaymentSettingsV1FvEmeaOK, error)

	PendingPaymentDeleteV1FvEmea(params *PendingPaymentDeleteV1FvEmeaParams) (*PendingPaymentDeleteV1FvEmeaOK, error)

	PendingPaymentListV1FvEmea(params *PendingPaymentListV1FvEmeaParams) (*PendingPaymentListV1FvEmeaOK, error)

	ProcessPaymentV1FvEmea(params *ProcessPaymentV1FvEmeaParams) (*ProcessPaymentV1FvEmeaOK, error)

	RealTimePaymentWithMemoV2FvEmea(params *RealTimePaymentWithMemoV2FvEmeaParams) (*RealTimePaymentWithMemoV2FvEmeaOK, error)

	RealTimePaymentV1FvEmea(params *RealTimePaymentV1FvEmeaParams) (*RealTimePaymentV1FvEmeaOK, error)

	RealTimePaymentwithMemoV1FvEmea(params *RealTimePaymentwithMemoV1FvEmeaParams) (*RealTimePaymentwithMemoV1FvEmeaOK, error)

	ReoccuringPaymentDeleteV1FvEmea(params *ReoccuringPaymentDeleteV1FvEmeaParams) (*ReoccuringPaymentDeleteV1FvEmeaOK, error)

	ReoccuringPaymentListV1FvEmea(params *ReoccuringPaymentListV1FvEmeaParams) (*ReoccuringPaymentListV1FvEmeaOK, error)

	StoredPayingCardDeleteV1FvEmea(params *StoredPayingCardDeleteV1FvEmeaParams) (*StoredPayingCardDeleteV1FvEmeaOK, error)

	StoredPayingCardUpdateV1FvEmea(params *StoredPayingCardUpdateV1FvEmeaParams) (*StoredPayingCardUpdateV1FvEmeaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FasterPaymentV1FvEmea fasters payment

  The purpose of this service is to apply payment and payment reversal transactions immediately to credit card accounts  when it is received as a Faster Payment transaction from the external remitter. The Faster Payment functionality supports duplicate checking for payment transactions. The duplicate payments i.e. same Org, Account and FPS ID, will be rejected and reported. The Faster Payment functionality also supports immediate payment reversals within the same business day. It ensures that payments are only reversed if they have been applied, matching payment transaction with same Org, Account and FPS ID. For successful payment transactions, memo credit will be applied and account Open To Buy ARIQ01 will be increased by that much amount. For successful payment reversal transaction, memo debit will be applied and account OTB ARIQ01 will be decreased by that much amount. The service will return a fail indicator F in the response message with an appropriate error code and description for duplicate payments and non-matching payment reversal transactions. The batch process will extract Faster Payment transactions and post it to CMS accounts.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) FasterPaymentV1FvEmea(params *FasterPaymentV1FvEmeaParams) (*FasterPaymentV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFasterPaymentV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fasterPayment_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/fasterPayment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FasterPaymentV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FasterPaymentV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fasterPayment_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MoveFundsV1FvEmea this service is used to perform the balance transfer or money transfers via the move funds service message which in turn is a f i c s direct service

  Financial institution(issuer) part of First Vision system can request this service by providing all the required details to perform BT/MT. As a pre-requisite FICS system configuration should exist for the client with necessary level of setup. Based on the type of request, number of internal checks and actions will be performed on the cardholders account and the response will be sent back.
*/
func (a *Client) MoveFundsV1FvEmea(params *MoveFundsV1FvEmeaParams) (*MoveFundsV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveFundsV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "moveFunds_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/moveFunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveFundsV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MoveFundsV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for moveFunds_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PaymentDueEstimationV1FvEmea payments due estimation

  Estimated payment due calculation service<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) PaymentDueEstimationV1FvEmea(params *PaymentDueEstimationV1FvEmeaParams) (*PaymentDueEstimationV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentDueEstimationV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "paymentDueEstimation_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/paymentDueEstimation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentDueEstimationV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PaymentDueEstimationV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for paymentDueEstimation_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PaymentHistoryInquiryV1FvEmea payments history inquiry

  The Account Payment History Inquiry service allows a calling program to request Payment History information associated with an Account.  All of the Account Payment History information that can be seen online on the Account Payment History screens ARPH01 and ARPH02 is included in the output message. The Organization Parameters must be set to retain payment history or a standard Fail error message will be returned in the Output Message.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) PaymentHistoryInquiryV1FvEmea(params *PaymentHistoryInquiryV1FvEmeaParams) (*PaymentHistoryInquiryV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentHistoryInquiryV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "paymentHistoryInquiry_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/paymentHistoryInquiry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentHistoryInquiryV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PaymentHistoryInquiryV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for paymentHistoryInquiry_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PaymentSettingsV1FvEmea retrieves initial settings based on a vision card or account number e g current stored cards issuer department settings
*/
func (a *Client) PaymentSettingsV1FvEmea(params *PaymentSettingsV1FvEmeaParams) (*PaymentSettingsV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentSettingsV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "paymentSettings_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/paymentSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentSettingsV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PaymentSettingsV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for paymentSettings_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PendingPaymentDeleteV1FvEmea this will cancel any number of continuous payment authority c p as and or future payments
*/
func (a *Client) PendingPaymentDeleteV1FvEmea(params *PendingPaymentDeleteV1FvEmeaParams) (*PendingPaymentDeleteV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPendingPaymentDeleteV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pendingPaymentDelete_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/pendingPaymentDelete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PendingPaymentDeleteV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PendingPaymentDeleteV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pendingPaymentDelete_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PendingPaymentListV1FvEmea this will provide a list of all continuous payment authority c p as and future payments set on an account
*/
func (a *Client) PendingPaymentListV1FvEmea(params *PendingPaymentListV1FvEmeaParams) (*PendingPaymentListV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPendingPaymentListV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pendingPaymentList_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/pendingPaymentList",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PendingPaymentListV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PendingPaymentListV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pendingPaymentList_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProcessPaymentV1FvEmea allows you to make a payment to a card or account number on first vision or a bank product bank account loan mortgage etc includes support for 3 d secure an option to store the paying card for future use an option to use a stored card for payment and various other configuration settings use this service for adding a stored card without making a payment by setting the payment amount to zero a card check where the amount 0 needs to be done before storing a card this service also support 3 d s2 0 version fields and the conditional mandatory fields have to be passed in order for banks to be compliant for 3 d s2 0 version
*/
func (a *Client) ProcessPaymentV1FvEmea(params *ProcessPaymentV1FvEmeaParams) (*ProcessPaymentV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProcessPaymentV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "processPayment_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/processPayment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProcessPaymentV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProcessPaymentV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for processPayment_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RealTimePaymentWithMemoV2FvEmea reals time payment with memo

  This service is used to post the payment and payment reversal transactions via passing appropriate service type indicator in the request (PYMT for payment and PYRV for the reversals). For all successfully processed requests, the OTB of the cardholders is updated on a real time basis, which allows cardholders to make transaction as per the new increased/decreased OTB. All accepted payments and reversals are processed in the overnight batch automatically.<p><p>Each payment requires a unique payment reference number to be sent in the request message to avoid duplicate messages being received and processed. If the payment reference number has already been processed by the service on that day, the subsequent request will be rejected.<p><p>In order to reverse a payment the original payment reference number must be provided in the request.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) RealTimePaymentWithMemoV2FvEmea(params *RealTimePaymentWithMemoV2FvEmeaParams) (*RealTimePaymentWithMemoV2FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRealTimePaymentWithMemoV2FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "realTimePaymentWithMemo_v2_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v2/realTimePaymentWithMemo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RealTimePaymentWithMemoV2FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RealTimePaymentWithMemoV2FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for realTimePaymentWithMemo_v2_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RealTimePaymentV1FvEmea reals time payment

  This service is used to post the payment and payment reversal transactions via passing appropriate service type indicator in the request (PYMT for payment and PYRV for the reversals). For all successfully processed requests, the OTB of the cardholders is updated on a real time basis, which allows cardholders to make transaction as per the new increased/decreased OTB. All accepted payments and reversals are processed in the overnight batch automatically.<p><p>Each payment requires a unique payment reference number to be sent in the request message to avoid duplicate messages being received and processed. If the payment reference number has already been processed by the service on that day, the subsequent request will be rejected.<p><p>In order to reverse a payment the original payment reference number must be provided in the request.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) RealTimePaymentV1FvEmea(params *RealTimePaymentV1FvEmeaParams) (*RealTimePaymentV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRealTimePaymentV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "realTimePayment_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/realTimePayment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RealTimePaymentV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RealTimePaymentV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for realTimePayment_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RealTimePaymentwithMemoV1FvEmea reals time paymentwith memo

  This service is deprecated and will be supported until January 2021. Please switch to Real Time Payment With Memo version 2 before this date. Note: End point changed from realTimePaymentwithMemo to realTimePaymentWithMemo. <BR/><BR/>This service is used to post the payment and payment reversal transactions via passing appropriate service type indicator in the request (PYMT for payment and PYRV for the reversals). For all successfully processed requests, the OTB of the cardholders is updated on a real time basis, which allows cardholders to make transaction as per the new increased/decreased OTB. All accepted payments and reversals are processed in the overnight batch automatically.<p><p>Each payment requires a unique payment reference number to be sent in the request message to avoid duplicate messages being received and processed. If the payment reference number has already been processed by the service on that day, the subsequent request will be rejected.<p><p>In order to reverse a payment the original payment reference number must be provided in the request.<BR/><BR/><p>Fields that are not provided in the Request object will be initialised to their default values. All numeric fields are initialised to zero and alphanumeric fields initialised to spaces</p>
*/
func (a *Client) RealTimePaymentwithMemoV1FvEmea(params *RealTimePaymentwithMemoV1FvEmeaParams) (*RealTimePaymentwithMemoV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRealTimePaymentwithMemoV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "realTimePaymentwithMemo_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/realTimePaymentwithMemo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RealTimePaymentwithMemoV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RealTimePaymentwithMemoV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for realTimePaymentwithMemo_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReoccuringPaymentDeleteV1FvEmea this will cancel a continuous payment authority so that future payments are not taken from the customer account
*/
func (a *Client) ReoccuringPaymentDeleteV1FvEmea(params *ReoccuringPaymentDeleteV1FvEmeaParams) (*ReoccuringPaymentDeleteV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReoccuringPaymentDeleteV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reoccuringPaymentDelete_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/reoccuringPaymentDelete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReoccuringPaymentDeleteV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReoccuringPaymentDeleteV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reoccuringPaymentDelete_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReoccuringPaymentListV1FvEmea this will provides a list of active continuous payment authority c p as set up for a customer
*/
func (a *Client) ReoccuringPaymentListV1FvEmea(params *ReoccuringPaymentListV1FvEmeaParams) (*ReoccuringPaymentListV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReoccuringPaymentListV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reoccuringPaymentList_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/reoccuringPaymentList",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReoccuringPaymentListV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReoccuringPaymentListV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reoccuringPaymentList_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StoredPayingCardDeleteV1FvEmea this will cancel a clients card that was previously stored in the first pay system
*/
func (a *Client) StoredPayingCardDeleteV1FvEmea(params *StoredPayingCardDeleteV1FvEmeaParams) (*StoredPayingCardDeleteV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStoredPayingCardDeleteV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "storedPayingCardDelete_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/storedPayingCardDelete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StoredPayingCardDeleteV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StoredPayingCardDeleteV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storedPayingCardDelete_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StoredPayingCardUpdateV1FvEmea sets a stored card as being the default card in the first pay system
*/
func (a *Client) StoredPayingCardUpdateV1FvEmea(params *StoredPayingCardUpdateV1FvEmeaParams) (*StoredPayingCardUpdateV1FvEmeaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStoredPayingCardUpdateV1FvEmeaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "storedPayingCardUpdate_v1_fv_emea",
		Method:             "POST",
		PathPattern:        "/fv_emea/v1/storedPayingCardUpdate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StoredPayingCardUpdateV1FvEmeaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StoredPayingCardUpdateV1FvEmeaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storedPayingCardUpdate_v1_fv_emea: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
