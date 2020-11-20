// Code generated by go-swagger; DO NOT EDIT.

package customer_records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// CustomerAccountsListV1FvEmeaReader is a Reader for the CustomerAccountsListV1FvEmea structure.
type CustomerAccountsListV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAccountsListV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerAccountsListV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomerAccountsListV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCustomerAccountsListV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCustomerAccountsListV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCustomerAccountsListV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCustomerAccountsListV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewCustomerAccountsListV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewCustomerAccountsListV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewCustomerAccountsListV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomerAccountsListV1FvEmeaOK creates a CustomerAccountsListV1FvEmeaOK with default headers values
func NewCustomerAccountsListV1FvEmeaOK() *CustomerAccountsListV1FvEmeaOK {
	return &CustomerAccountsListV1FvEmeaOK{}
}

/*CustomerAccountsListV1FvEmeaOK handles this case with default header values.

successful operation
*/
type CustomerAccountsListV1FvEmeaOK struct {
	Payload *models.CustomerAccountsListResponse
}

func (o *CustomerAccountsListV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *CustomerAccountsListV1FvEmeaOK) GetPayload() *models.CustomerAccountsListResponse {
	return o.Payload
}

func (o *CustomerAccountsListV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerAccountsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountsListV1FvEmeaBadRequest creates a CustomerAccountsListV1FvEmeaBadRequest with default headers values
func NewCustomerAccountsListV1FvEmeaBadRequest() *CustomerAccountsListV1FvEmeaBadRequest {
	return &CustomerAccountsListV1FvEmeaBadRequest{}
}

/*CustomerAccountsListV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type CustomerAccountsListV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *CustomerAccountsListV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerAccountsListV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CustomerAccountsListV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountsListV1FvEmeaUnauthorized creates a CustomerAccountsListV1FvEmeaUnauthorized with default headers values
func NewCustomerAccountsListV1FvEmeaUnauthorized() *CustomerAccountsListV1FvEmeaUnauthorized {
	return &CustomerAccountsListV1FvEmeaUnauthorized{}
}

/*CustomerAccountsListV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type CustomerAccountsListV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *CustomerAccountsListV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *CustomerAccountsListV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CustomerAccountsListV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountsListV1FvEmeaForbidden creates a CustomerAccountsListV1FvEmeaForbidden with default headers values
func NewCustomerAccountsListV1FvEmeaForbidden() *CustomerAccountsListV1FvEmeaForbidden {
	return &CustomerAccountsListV1FvEmeaForbidden{}
}

/*CustomerAccountsListV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type CustomerAccountsListV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *CustomerAccountsListV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *CustomerAccountsListV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CustomerAccountsListV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountsListV1FvEmeaNotFound creates a CustomerAccountsListV1FvEmeaNotFound with default headers values
func NewCustomerAccountsListV1FvEmeaNotFound() *CustomerAccountsListV1FvEmeaNotFound {
	return &CustomerAccountsListV1FvEmeaNotFound{}
}

/*CustomerAccountsListV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type CustomerAccountsListV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *CustomerAccountsListV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *CustomerAccountsListV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CustomerAccountsListV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountsListV1FvEmeaTooManyRequests creates a CustomerAccountsListV1FvEmeaTooManyRequests with default headers values
func NewCustomerAccountsListV1FvEmeaTooManyRequests() *CustomerAccountsListV1FvEmeaTooManyRequests {
	return &CustomerAccountsListV1FvEmeaTooManyRequests{}
}

/*CustomerAccountsListV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type CustomerAccountsListV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *CustomerAccountsListV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *CustomerAccountsListV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CustomerAccountsListV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountsListV1FvEmeaStatus452 creates a CustomerAccountsListV1FvEmeaStatus452 with default headers values
func NewCustomerAccountsListV1FvEmeaStatus452() *CustomerAccountsListV1FvEmeaStatus452 {
	return &CustomerAccountsListV1FvEmeaStatus452{}
}

/*CustomerAccountsListV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type CustomerAccountsListV1FvEmeaStatus452 struct {
}

func (o *CustomerAccountsListV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaStatus452 ", 452)
}

func (o *CustomerAccountsListV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerAccountsListV1FvEmeaStatus453 creates a CustomerAccountsListV1FvEmeaStatus453 with default headers values
func NewCustomerAccountsListV1FvEmeaStatus453() *CustomerAccountsListV1FvEmeaStatus453 {
	return &CustomerAccountsListV1FvEmeaStatus453{}
}

/*CustomerAccountsListV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type CustomerAccountsListV1FvEmeaStatus453 struct {
}

func (o *CustomerAccountsListV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaStatus453 ", 453)
}

func (o *CustomerAccountsListV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerAccountsListV1FvEmeaStatus455 creates a CustomerAccountsListV1FvEmeaStatus455 with default headers values
func NewCustomerAccountsListV1FvEmeaStatus455() *CustomerAccountsListV1FvEmeaStatus455 {
	return &CustomerAccountsListV1FvEmeaStatus455{}
}

/*CustomerAccountsListV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type CustomerAccountsListV1FvEmeaStatus455 struct {
}

func (o *CustomerAccountsListV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/customerAccountsList][%d] customerAccountsListV1FvEmeaStatus455 ", 455)
}

func (o *CustomerAccountsListV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}