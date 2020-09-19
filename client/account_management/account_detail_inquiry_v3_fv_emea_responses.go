// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// AccountDetailInquiryV3FvEmeaReader is a Reader for the AccountDetailInquiryV3FvEmea structure.
type AccountDetailInquiryV3FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountDetailInquiryV3FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountDetailInquiryV3FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountDetailInquiryV3FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountDetailInquiryV3FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccountDetailInquiryV3FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountDetailInquiryV3FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAccountDetailInquiryV3FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewAccountDetailInquiryV3FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewAccountDetailInquiryV3FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewAccountDetailInquiryV3FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountDetailInquiryV3FvEmeaOK creates a AccountDetailInquiryV3FvEmeaOK with default headers values
func NewAccountDetailInquiryV3FvEmeaOK() *AccountDetailInquiryV3FvEmeaOK {
	return &AccountDetailInquiryV3FvEmeaOK{}
}

/*AccountDetailInquiryV3FvEmeaOK handles this case with default header values.

successful operation
*/
type AccountDetailInquiryV3FvEmeaOK struct {
	Payload *models.AccountDetailInquiryResponse3
}

func (o *AccountDetailInquiryV3FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaOK  %+v", 200, o.Payload)
}

func (o *AccountDetailInquiryV3FvEmeaOK) GetPayload() *models.AccountDetailInquiryResponse3 {
	return o.Payload
}

func (o *AccountDetailInquiryV3FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountDetailInquiryResponse3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDetailInquiryV3FvEmeaBadRequest creates a AccountDetailInquiryV3FvEmeaBadRequest with default headers values
func NewAccountDetailInquiryV3FvEmeaBadRequest() *AccountDetailInquiryV3FvEmeaBadRequest {
	return &AccountDetailInquiryV3FvEmeaBadRequest{}
}

/*AccountDetailInquiryV3FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type AccountDetailInquiryV3FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *AccountDetailInquiryV3FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *AccountDetailInquiryV3FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountDetailInquiryV3FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDetailInquiryV3FvEmeaUnauthorized creates a AccountDetailInquiryV3FvEmeaUnauthorized with default headers values
func NewAccountDetailInquiryV3FvEmeaUnauthorized() *AccountDetailInquiryV3FvEmeaUnauthorized {
	return &AccountDetailInquiryV3FvEmeaUnauthorized{}
}

/*AccountDetailInquiryV3FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type AccountDetailInquiryV3FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *AccountDetailInquiryV3FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountDetailInquiryV3FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountDetailInquiryV3FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDetailInquiryV3FvEmeaForbidden creates a AccountDetailInquiryV3FvEmeaForbidden with default headers values
func NewAccountDetailInquiryV3FvEmeaForbidden() *AccountDetailInquiryV3FvEmeaForbidden {
	return &AccountDetailInquiryV3FvEmeaForbidden{}
}

/*AccountDetailInquiryV3FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type AccountDetailInquiryV3FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *AccountDetailInquiryV3FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *AccountDetailInquiryV3FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountDetailInquiryV3FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDetailInquiryV3FvEmeaNotFound creates a AccountDetailInquiryV3FvEmeaNotFound with default headers values
func NewAccountDetailInquiryV3FvEmeaNotFound() *AccountDetailInquiryV3FvEmeaNotFound {
	return &AccountDetailInquiryV3FvEmeaNotFound{}
}

/*AccountDetailInquiryV3FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type AccountDetailInquiryV3FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *AccountDetailInquiryV3FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *AccountDetailInquiryV3FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountDetailInquiryV3FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDetailInquiryV3FvEmeaTooManyRequests creates a AccountDetailInquiryV3FvEmeaTooManyRequests with default headers values
func NewAccountDetailInquiryV3FvEmeaTooManyRequests() *AccountDetailInquiryV3FvEmeaTooManyRequests {
	return &AccountDetailInquiryV3FvEmeaTooManyRequests{}
}

/*AccountDetailInquiryV3FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type AccountDetailInquiryV3FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *AccountDetailInquiryV3FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *AccountDetailInquiryV3FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountDetailInquiryV3FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountDetailInquiryV3FvEmeaStatus452 creates a AccountDetailInquiryV3FvEmeaStatus452 with default headers values
func NewAccountDetailInquiryV3FvEmeaStatus452() *AccountDetailInquiryV3FvEmeaStatus452 {
	return &AccountDetailInquiryV3FvEmeaStatus452{}
}

/*AccountDetailInquiryV3FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type AccountDetailInquiryV3FvEmeaStatus452 struct {
}

func (o *AccountDetailInquiryV3FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaStatus452 ", 452)
}

func (o *AccountDetailInquiryV3FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountDetailInquiryV3FvEmeaStatus453 creates a AccountDetailInquiryV3FvEmeaStatus453 with default headers values
func NewAccountDetailInquiryV3FvEmeaStatus453() *AccountDetailInquiryV3FvEmeaStatus453 {
	return &AccountDetailInquiryV3FvEmeaStatus453{}
}

/*AccountDetailInquiryV3FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type AccountDetailInquiryV3FvEmeaStatus453 struct {
}

func (o *AccountDetailInquiryV3FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaStatus453 ", 453)
}

func (o *AccountDetailInquiryV3FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountDetailInquiryV3FvEmeaStatus455 creates a AccountDetailInquiryV3FvEmeaStatus455 with default headers values
func NewAccountDetailInquiryV3FvEmeaStatus455() *AccountDetailInquiryV3FvEmeaStatus455 {
	return &AccountDetailInquiryV3FvEmeaStatus455{}
}

/*AccountDetailInquiryV3FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type AccountDetailInquiryV3FvEmeaStatus455 struct {
}

func (o *AccountDetailInquiryV3FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v3/accountDetailInquiry][%d] accountDetailInquiryV3FvEmeaStatus455 ", 455)
}

func (o *AccountDetailInquiryV3FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
