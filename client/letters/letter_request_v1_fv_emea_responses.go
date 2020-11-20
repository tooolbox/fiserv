// Code generated by go-swagger; DO NOT EDIT.

package letters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// LetterRequestV1FvEmeaReader is a Reader for the LetterRequestV1FvEmea structure.
type LetterRequestV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LetterRequestV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLetterRequestV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLetterRequestV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLetterRequestV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLetterRequestV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLetterRequestV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLetterRequestV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLetterRequestV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLetterRequestV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLetterRequestV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLetterRequestV1FvEmeaOK creates a LetterRequestV1FvEmeaOK with default headers values
func NewLetterRequestV1FvEmeaOK() *LetterRequestV1FvEmeaOK {
	return &LetterRequestV1FvEmeaOK{}
}

/*LetterRequestV1FvEmeaOK handles this case with default header values.

successful operation
*/
type LetterRequestV1FvEmeaOK struct {
	Payload *models.LetterRequestResponse
}

func (o *LetterRequestV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LetterRequestV1FvEmeaOK) GetPayload() *models.LetterRequestResponse {
	return o.Payload
}

func (o *LetterRequestV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LetterRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLetterRequestV1FvEmeaBadRequest creates a LetterRequestV1FvEmeaBadRequest with default headers values
func NewLetterRequestV1FvEmeaBadRequest() *LetterRequestV1FvEmeaBadRequest {
	return &LetterRequestV1FvEmeaBadRequest{}
}

/*LetterRequestV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LetterRequestV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LetterRequestV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LetterRequestV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LetterRequestV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLetterRequestV1FvEmeaUnauthorized creates a LetterRequestV1FvEmeaUnauthorized with default headers values
func NewLetterRequestV1FvEmeaUnauthorized() *LetterRequestV1FvEmeaUnauthorized {
	return &LetterRequestV1FvEmeaUnauthorized{}
}

/*LetterRequestV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LetterRequestV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LetterRequestV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LetterRequestV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LetterRequestV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLetterRequestV1FvEmeaForbidden creates a LetterRequestV1FvEmeaForbidden with default headers values
func NewLetterRequestV1FvEmeaForbidden() *LetterRequestV1FvEmeaForbidden {
	return &LetterRequestV1FvEmeaForbidden{}
}

/*LetterRequestV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LetterRequestV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LetterRequestV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LetterRequestV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LetterRequestV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLetterRequestV1FvEmeaNotFound creates a LetterRequestV1FvEmeaNotFound with default headers values
func NewLetterRequestV1FvEmeaNotFound() *LetterRequestV1FvEmeaNotFound {
	return &LetterRequestV1FvEmeaNotFound{}
}

/*LetterRequestV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LetterRequestV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LetterRequestV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LetterRequestV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LetterRequestV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLetterRequestV1FvEmeaTooManyRequests creates a LetterRequestV1FvEmeaTooManyRequests with default headers values
func NewLetterRequestV1FvEmeaTooManyRequests() *LetterRequestV1FvEmeaTooManyRequests {
	return &LetterRequestV1FvEmeaTooManyRequests{}
}

/*LetterRequestV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LetterRequestV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LetterRequestV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LetterRequestV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LetterRequestV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLetterRequestV1FvEmeaStatus452 creates a LetterRequestV1FvEmeaStatus452 with default headers values
func NewLetterRequestV1FvEmeaStatus452() *LetterRequestV1FvEmeaStatus452 {
	return &LetterRequestV1FvEmeaStatus452{}
}

/*LetterRequestV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LetterRequestV1FvEmeaStatus452 struct {
}

func (o *LetterRequestV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaStatus452 ", 452)
}

func (o *LetterRequestV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLetterRequestV1FvEmeaStatus453 creates a LetterRequestV1FvEmeaStatus453 with default headers values
func NewLetterRequestV1FvEmeaStatus453() *LetterRequestV1FvEmeaStatus453 {
	return &LetterRequestV1FvEmeaStatus453{}
}

/*LetterRequestV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LetterRequestV1FvEmeaStatus453 struct {
}

func (o *LetterRequestV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaStatus453 ", 453)
}

func (o *LetterRequestV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLetterRequestV1FvEmeaStatus455 creates a LetterRequestV1FvEmeaStatus455 with default headers values
func NewLetterRequestV1FvEmeaStatus455() *LetterRequestV1FvEmeaStatus455 {
	return &LetterRequestV1FvEmeaStatus455{}
}

/*LetterRequestV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LetterRequestV1FvEmeaStatus455 struct {
}

func (o *LetterRequestV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/letterRequest][%d] letterRequestV1FvEmeaStatus455 ", 455)
}

func (o *LetterRequestV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}