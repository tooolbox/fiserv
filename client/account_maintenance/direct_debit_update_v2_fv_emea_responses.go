// Code generated by go-swagger; DO NOT EDIT.

package account_maintenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// DirectDebitUpdateV2FvEmeaReader is a Reader for the DirectDebitUpdateV2FvEmea structure.
type DirectDebitUpdateV2FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DirectDebitUpdateV2FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDirectDebitUpdateV2FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDirectDebitUpdateV2FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDirectDebitUpdateV2FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDirectDebitUpdateV2FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDirectDebitUpdateV2FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDirectDebitUpdateV2FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewDirectDebitUpdateV2FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewDirectDebitUpdateV2FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewDirectDebitUpdateV2FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDirectDebitUpdateV2FvEmeaOK creates a DirectDebitUpdateV2FvEmeaOK with default headers values
func NewDirectDebitUpdateV2FvEmeaOK() *DirectDebitUpdateV2FvEmeaOK {
	return &DirectDebitUpdateV2FvEmeaOK{}
}

/*DirectDebitUpdateV2FvEmeaOK handles this case with default header values.

successful operation
*/
type DirectDebitUpdateV2FvEmeaOK struct {
	Payload *models.DirectDebitUpdateResponse2
}

func (o *DirectDebitUpdateV2FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaOK  %+v", 200, o.Payload)
}

func (o *DirectDebitUpdateV2FvEmeaOK) GetPayload() *models.DirectDebitUpdateResponse2 {
	return o.Payload
}

func (o *DirectDebitUpdateV2FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectDebitUpdateResponse2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitUpdateV2FvEmeaBadRequest creates a DirectDebitUpdateV2FvEmeaBadRequest with default headers values
func NewDirectDebitUpdateV2FvEmeaBadRequest() *DirectDebitUpdateV2FvEmeaBadRequest {
	return &DirectDebitUpdateV2FvEmeaBadRequest{}
}

/*DirectDebitUpdateV2FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type DirectDebitUpdateV2FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitUpdateV2FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *DirectDebitUpdateV2FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitUpdateV2FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitUpdateV2FvEmeaUnauthorized creates a DirectDebitUpdateV2FvEmeaUnauthorized with default headers values
func NewDirectDebitUpdateV2FvEmeaUnauthorized() *DirectDebitUpdateV2FvEmeaUnauthorized {
	return &DirectDebitUpdateV2FvEmeaUnauthorized{}
}

/*DirectDebitUpdateV2FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type DirectDebitUpdateV2FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitUpdateV2FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *DirectDebitUpdateV2FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitUpdateV2FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitUpdateV2FvEmeaForbidden creates a DirectDebitUpdateV2FvEmeaForbidden with default headers values
func NewDirectDebitUpdateV2FvEmeaForbidden() *DirectDebitUpdateV2FvEmeaForbidden {
	return &DirectDebitUpdateV2FvEmeaForbidden{}
}

/*DirectDebitUpdateV2FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type DirectDebitUpdateV2FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitUpdateV2FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *DirectDebitUpdateV2FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitUpdateV2FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitUpdateV2FvEmeaNotFound creates a DirectDebitUpdateV2FvEmeaNotFound with default headers values
func NewDirectDebitUpdateV2FvEmeaNotFound() *DirectDebitUpdateV2FvEmeaNotFound {
	return &DirectDebitUpdateV2FvEmeaNotFound{}
}

/*DirectDebitUpdateV2FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type DirectDebitUpdateV2FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitUpdateV2FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *DirectDebitUpdateV2FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitUpdateV2FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitUpdateV2FvEmeaTooManyRequests creates a DirectDebitUpdateV2FvEmeaTooManyRequests with default headers values
func NewDirectDebitUpdateV2FvEmeaTooManyRequests() *DirectDebitUpdateV2FvEmeaTooManyRequests {
	return &DirectDebitUpdateV2FvEmeaTooManyRequests{}
}

/*DirectDebitUpdateV2FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type DirectDebitUpdateV2FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitUpdateV2FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DirectDebitUpdateV2FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitUpdateV2FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitUpdateV2FvEmeaStatus452 creates a DirectDebitUpdateV2FvEmeaStatus452 with default headers values
func NewDirectDebitUpdateV2FvEmeaStatus452() *DirectDebitUpdateV2FvEmeaStatus452 {
	return &DirectDebitUpdateV2FvEmeaStatus452{}
}

/*DirectDebitUpdateV2FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type DirectDebitUpdateV2FvEmeaStatus452 struct {
}

func (o *DirectDebitUpdateV2FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaStatus452 ", 452)
}

func (o *DirectDebitUpdateV2FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDirectDebitUpdateV2FvEmeaStatus453 creates a DirectDebitUpdateV2FvEmeaStatus453 with default headers values
func NewDirectDebitUpdateV2FvEmeaStatus453() *DirectDebitUpdateV2FvEmeaStatus453 {
	return &DirectDebitUpdateV2FvEmeaStatus453{}
}

/*DirectDebitUpdateV2FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type DirectDebitUpdateV2FvEmeaStatus453 struct {
}

func (o *DirectDebitUpdateV2FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaStatus453 ", 453)
}

func (o *DirectDebitUpdateV2FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDirectDebitUpdateV2FvEmeaStatus455 creates a DirectDebitUpdateV2FvEmeaStatus455 with default headers values
func NewDirectDebitUpdateV2FvEmeaStatus455() *DirectDebitUpdateV2FvEmeaStatus455 {
	return &DirectDebitUpdateV2FvEmeaStatus455{}
}

/*DirectDebitUpdateV2FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type DirectDebitUpdateV2FvEmeaStatus455 struct {
}

func (o *DirectDebitUpdateV2FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/directDebitUpdate][%d] directDebitUpdateV2FvEmeaStatus455 ", 455)
}

func (o *DirectDebitUpdateV2FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}