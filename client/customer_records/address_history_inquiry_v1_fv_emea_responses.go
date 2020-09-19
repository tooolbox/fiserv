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

// AddressHistoryInquiryV1FvEmeaReader is a Reader for the AddressHistoryInquiryV1FvEmea structure.
type AddressHistoryInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressHistoryInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressHistoryInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddressHistoryInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddressHistoryInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddressHistoryInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddressHistoryInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddressHistoryInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewAddressHistoryInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewAddressHistoryInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewAddressHistoryInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddressHistoryInquiryV1FvEmeaOK creates a AddressHistoryInquiryV1FvEmeaOK with default headers values
func NewAddressHistoryInquiryV1FvEmeaOK() *AddressHistoryInquiryV1FvEmeaOK {
	return &AddressHistoryInquiryV1FvEmeaOK{}
}

/*AddressHistoryInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type AddressHistoryInquiryV1FvEmeaOK struct {
	Payload *models.AddressHistoryInquiryResponse
}

func (o *AddressHistoryInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *AddressHistoryInquiryV1FvEmeaOK) GetPayload() *models.AddressHistoryInquiryResponse {
	return o.Payload
}

func (o *AddressHistoryInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressHistoryInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaBadRequest creates a AddressHistoryInquiryV1FvEmeaBadRequest with default headers values
func NewAddressHistoryInquiryV1FvEmeaBadRequest() *AddressHistoryInquiryV1FvEmeaBadRequest {
	return &AddressHistoryInquiryV1FvEmeaBadRequest{}
}

/*AddressHistoryInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type AddressHistoryInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *AddressHistoryInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *AddressHistoryInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressHistoryInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaUnauthorized creates a AddressHistoryInquiryV1FvEmeaUnauthorized with default headers values
func NewAddressHistoryInquiryV1FvEmeaUnauthorized() *AddressHistoryInquiryV1FvEmeaUnauthorized {
	return &AddressHistoryInquiryV1FvEmeaUnauthorized{}
}

/*AddressHistoryInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type AddressHistoryInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *AddressHistoryInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *AddressHistoryInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressHistoryInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaForbidden creates a AddressHistoryInquiryV1FvEmeaForbidden with default headers values
func NewAddressHistoryInquiryV1FvEmeaForbidden() *AddressHistoryInquiryV1FvEmeaForbidden {
	return &AddressHistoryInquiryV1FvEmeaForbidden{}
}

/*AddressHistoryInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type AddressHistoryInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *AddressHistoryInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *AddressHistoryInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressHistoryInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaNotFound creates a AddressHistoryInquiryV1FvEmeaNotFound with default headers values
func NewAddressHistoryInquiryV1FvEmeaNotFound() *AddressHistoryInquiryV1FvEmeaNotFound {
	return &AddressHistoryInquiryV1FvEmeaNotFound{}
}

/*AddressHistoryInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type AddressHistoryInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *AddressHistoryInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *AddressHistoryInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressHistoryInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaTooManyRequests creates a AddressHistoryInquiryV1FvEmeaTooManyRequests with default headers values
func NewAddressHistoryInquiryV1FvEmeaTooManyRequests() *AddressHistoryInquiryV1FvEmeaTooManyRequests {
	return &AddressHistoryInquiryV1FvEmeaTooManyRequests{}
}

/*AddressHistoryInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type AddressHistoryInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *AddressHistoryInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddressHistoryInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressHistoryInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaStatus452 creates a AddressHistoryInquiryV1FvEmeaStatus452 with default headers values
func NewAddressHistoryInquiryV1FvEmeaStatus452() *AddressHistoryInquiryV1FvEmeaStatus452 {
	return &AddressHistoryInquiryV1FvEmeaStatus452{}
}

/*AddressHistoryInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type AddressHistoryInquiryV1FvEmeaStatus452 struct {
}

func (o *AddressHistoryInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaStatus452 ", 452)
}

func (o *AddressHistoryInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaStatus453 creates a AddressHistoryInquiryV1FvEmeaStatus453 with default headers values
func NewAddressHistoryInquiryV1FvEmeaStatus453() *AddressHistoryInquiryV1FvEmeaStatus453 {
	return &AddressHistoryInquiryV1FvEmeaStatus453{}
}

/*AddressHistoryInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type AddressHistoryInquiryV1FvEmeaStatus453 struct {
}

func (o *AddressHistoryInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaStatus453 ", 453)
}

func (o *AddressHistoryInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressHistoryInquiryV1FvEmeaStatus455 creates a AddressHistoryInquiryV1FvEmeaStatus455 with default headers values
func NewAddressHistoryInquiryV1FvEmeaStatus455() *AddressHistoryInquiryV1FvEmeaStatus455 {
	return &AddressHistoryInquiryV1FvEmeaStatus455{}
}

/*AddressHistoryInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type AddressHistoryInquiryV1FvEmeaStatus455 struct {
}

func (o *AddressHistoryInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressHistoryInquiry][%d] addressHistoryInquiryV1FvEmeaStatus455 ", 455)
}

func (o *AddressHistoryInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
