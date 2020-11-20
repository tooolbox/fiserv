// Code generated by go-swagger; DO NOT EDIT.

package card_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// CardIDToPanDecryptV1FvEmeaReader is a Reader for the CardIDToPanDecryptV1FvEmea structure.
type CardIDToPanDecryptV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CardIDToPanDecryptV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCardIDToPanDecryptV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCardIDToPanDecryptV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCardIDToPanDecryptV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCardIDToPanDecryptV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCardIDToPanDecryptV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCardIDToPanDecryptV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewCardIDToPanDecryptV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewCardIDToPanDecryptV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewCardIDToPanDecryptV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCardIDToPanDecryptV1FvEmeaOK creates a CardIDToPanDecryptV1FvEmeaOK with default headers values
func NewCardIDToPanDecryptV1FvEmeaOK() *CardIDToPanDecryptV1FvEmeaOK {
	return &CardIDToPanDecryptV1FvEmeaOK{}
}

/*CardIDToPanDecryptV1FvEmeaOK handles this case with default header values.

successful operation
*/
type CardIDToPanDecryptV1FvEmeaOK struct {
	Payload *models.CardIDToPanDecryptResponse
}

func (o *CardIDToPanDecryptV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *CardIDToPanDecryptV1FvEmeaOK) GetPayload() *models.CardIDToPanDecryptResponse {
	return o.Payload
}

func (o *CardIDToPanDecryptV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardIDToPanDecryptResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaBadRequest creates a CardIDToPanDecryptV1FvEmeaBadRequest with default headers values
func NewCardIDToPanDecryptV1FvEmeaBadRequest() *CardIDToPanDecryptV1FvEmeaBadRequest {
	return &CardIDToPanDecryptV1FvEmeaBadRequest{}
}

/*CardIDToPanDecryptV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type CardIDToPanDecryptV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *CardIDToPanDecryptV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *CardIDToPanDecryptV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardIDToPanDecryptV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaUnauthorized creates a CardIDToPanDecryptV1FvEmeaUnauthorized with default headers values
func NewCardIDToPanDecryptV1FvEmeaUnauthorized() *CardIDToPanDecryptV1FvEmeaUnauthorized {
	return &CardIDToPanDecryptV1FvEmeaUnauthorized{}
}

/*CardIDToPanDecryptV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type CardIDToPanDecryptV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *CardIDToPanDecryptV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *CardIDToPanDecryptV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardIDToPanDecryptV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaForbidden creates a CardIDToPanDecryptV1FvEmeaForbidden with default headers values
func NewCardIDToPanDecryptV1FvEmeaForbidden() *CardIDToPanDecryptV1FvEmeaForbidden {
	return &CardIDToPanDecryptV1FvEmeaForbidden{}
}

/*CardIDToPanDecryptV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type CardIDToPanDecryptV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *CardIDToPanDecryptV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *CardIDToPanDecryptV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardIDToPanDecryptV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaNotFound creates a CardIDToPanDecryptV1FvEmeaNotFound with default headers values
func NewCardIDToPanDecryptV1FvEmeaNotFound() *CardIDToPanDecryptV1FvEmeaNotFound {
	return &CardIDToPanDecryptV1FvEmeaNotFound{}
}

/*CardIDToPanDecryptV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type CardIDToPanDecryptV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *CardIDToPanDecryptV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *CardIDToPanDecryptV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardIDToPanDecryptV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaTooManyRequests creates a CardIDToPanDecryptV1FvEmeaTooManyRequests with default headers values
func NewCardIDToPanDecryptV1FvEmeaTooManyRequests() *CardIDToPanDecryptV1FvEmeaTooManyRequests {
	return &CardIDToPanDecryptV1FvEmeaTooManyRequests{}
}

/*CardIDToPanDecryptV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type CardIDToPanDecryptV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *CardIDToPanDecryptV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *CardIDToPanDecryptV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *CardIDToPanDecryptV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaStatus452 creates a CardIDToPanDecryptV1FvEmeaStatus452 with default headers values
func NewCardIDToPanDecryptV1FvEmeaStatus452() *CardIDToPanDecryptV1FvEmeaStatus452 {
	return &CardIDToPanDecryptV1FvEmeaStatus452{}
}

/*CardIDToPanDecryptV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type CardIDToPanDecryptV1FvEmeaStatus452 struct {
}

func (o *CardIDToPanDecryptV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaStatus452 ", 452)
}

func (o *CardIDToPanDecryptV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaStatus453 creates a CardIDToPanDecryptV1FvEmeaStatus453 with default headers values
func NewCardIDToPanDecryptV1FvEmeaStatus453() *CardIDToPanDecryptV1FvEmeaStatus453 {
	return &CardIDToPanDecryptV1FvEmeaStatus453{}
}

/*CardIDToPanDecryptV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type CardIDToPanDecryptV1FvEmeaStatus453 struct {
}

func (o *CardIDToPanDecryptV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaStatus453 ", 453)
}

func (o *CardIDToPanDecryptV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCardIDToPanDecryptV1FvEmeaStatus455 creates a CardIDToPanDecryptV1FvEmeaStatus455 with default headers values
func NewCardIDToPanDecryptV1FvEmeaStatus455() *CardIDToPanDecryptV1FvEmeaStatus455 {
	return &CardIDToPanDecryptV1FvEmeaStatus455{}
}

/*CardIDToPanDecryptV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type CardIDToPanDecryptV1FvEmeaStatus455 struct {
}

func (o *CardIDToPanDecryptV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/cardIdToPanDecrypt][%d] cardIdToPanDecryptV1FvEmeaStatus455 ", 455)
}

func (o *CardIDToPanDecryptV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}