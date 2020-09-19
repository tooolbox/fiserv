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

// AddressListInquiryV1FvEmeaReader is a Reader for the AddressListInquiryV1FvEmea structure.
type AddressListInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressListInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressListInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddressListInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddressListInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddressListInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddressListInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddressListInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewAddressListInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewAddressListInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewAddressListInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddressListInquiryV1FvEmeaOK creates a AddressListInquiryV1FvEmeaOK with default headers values
func NewAddressListInquiryV1FvEmeaOK() *AddressListInquiryV1FvEmeaOK {
	return &AddressListInquiryV1FvEmeaOK{}
}

/*AddressListInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type AddressListInquiryV1FvEmeaOK struct {
	Payload *models.AddressListInquiryResponse
}

func (o *AddressListInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *AddressListInquiryV1FvEmeaOK) GetPayload() *models.AddressListInquiryResponse {
	return o.Payload
}

func (o *AddressListInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressListInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressListInquiryV1FvEmeaBadRequest creates a AddressListInquiryV1FvEmeaBadRequest with default headers values
func NewAddressListInquiryV1FvEmeaBadRequest() *AddressListInquiryV1FvEmeaBadRequest {
	return &AddressListInquiryV1FvEmeaBadRequest{}
}

/*AddressListInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type AddressListInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *AddressListInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *AddressListInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressListInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressListInquiryV1FvEmeaUnauthorized creates a AddressListInquiryV1FvEmeaUnauthorized with default headers values
func NewAddressListInquiryV1FvEmeaUnauthorized() *AddressListInquiryV1FvEmeaUnauthorized {
	return &AddressListInquiryV1FvEmeaUnauthorized{}
}

/*AddressListInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type AddressListInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *AddressListInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *AddressListInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressListInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressListInquiryV1FvEmeaForbidden creates a AddressListInquiryV1FvEmeaForbidden with default headers values
func NewAddressListInquiryV1FvEmeaForbidden() *AddressListInquiryV1FvEmeaForbidden {
	return &AddressListInquiryV1FvEmeaForbidden{}
}

/*AddressListInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type AddressListInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *AddressListInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *AddressListInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressListInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressListInquiryV1FvEmeaNotFound creates a AddressListInquiryV1FvEmeaNotFound with default headers values
func NewAddressListInquiryV1FvEmeaNotFound() *AddressListInquiryV1FvEmeaNotFound {
	return &AddressListInquiryV1FvEmeaNotFound{}
}

/*AddressListInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type AddressListInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *AddressListInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *AddressListInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressListInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressListInquiryV1FvEmeaTooManyRequests creates a AddressListInquiryV1FvEmeaTooManyRequests with default headers values
func NewAddressListInquiryV1FvEmeaTooManyRequests() *AddressListInquiryV1FvEmeaTooManyRequests {
	return &AddressListInquiryV1FvEmeaTooManyRequests{}
}

/*AddressListInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type AddressListInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *AddressListInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddressListInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AddressListInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressListInquiryV1FvEmeaStatus452 creates a AddressListInquiryV1FvEmeaStatus452 with default headers values
func NewAddressListInquiryV1FvEmeaStatus452() *AddressListInquiryV1FvEmeaStatus452 {
	return &AddressListInquiryV1FvEmeaStatus452{}
}

/*AddressListInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type AddressListInquiryV1FvEmeaStatus452 struct {
}

func (o *AddressListInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaStatus452 ", 452)
}

func (o *AddressListInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressListInquiryV1FvEmeaStatus453 creates a AddressListInquiryV1FvEmeaStatus453 with default headers values
func NewAddressListInquiryV1FvEmeaStatus453() *AddressListInquiryV1FvEmeaStatus453 {
	return &AddressListInquiryV1FvEmeaStatus453{}
}

/*AddressListInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type AddressListInquiryV1FvEmeaStatus453 struct {
}

func (o *AddressListInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaStatus453 ", 453)
}

func (o *AddressListInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressListInquiryV1FvEmeaStatus455 creates a AddressListInquiryV1FvEmeaStatus455 with default headers values
func NewAddressListInquiryV1FvEmeaStatus455() *AddressListInquiryV1FvEmeaStatus455 {
	return &AddressListInquiryV1FvEmeaStatus455{}
}

/*AddressListInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type AddressListInquiryV1FvEmeaStatus455 struct {
}

func (o *AddressListInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/addressListInquiry][%d] addressListInquiryV1FvEmeaStatus455 ", 455)
}

func (o *AddressListInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
