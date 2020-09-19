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

// DirectDebitInquiryV1FvEmeaReader is a Reader for the DirectDebitInquiryV1FvEmea structure.
type DirectDebitInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DirectDebitInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDirectDebitInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDirectDebitInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDirectDebitInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDirectDebitInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDirectDebitInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDirectDebitInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewDirectDebitInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewDirectDebitInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewDirectDebitInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewDirectDebitInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDirectDebitInquiryV1FvEmeaOK creates a DirectDebitInquiryV1FvEmeaOK with default headers values
func NewDirectDebitInquiryV1FvEmeaOK() *DirectDebitInquiryV1FvEmeaOK {
	return &DirectDebitInquiryV1FvEmeaOK{}
}

/*DirectDebitInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type DirectDebitInquiryV1FvEmeaOK struct {
	Payload *models.DirectDebitInquiryResponse
}

func (o *DirectDebitInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaOK) GetPayload() *models.DirectDebitInquiryResponse {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectDebitInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitInquiryV1FvEmeaBadRequest creates a DirectDebitInquiryV1FvEmeaBadRequest with default headers values
func NewDirectDebitInquiryV1FvEmeaBadRequest() *DirectDebitInquiryV1FvEmeaBadRequest {
	return &DirectDebitInquiryV1FvEmeaBadRequest{}
}

/*DirectDebitInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type DirectDebitInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitInquiryV1FvEmeaUnauthorized creates a DirectDebitInquiryV1FvEmeaUnauthorized with default headers values
func NewDirectDebitInquiryV1FvEmeaUnauthorized() *DirectDebitInquiryV1FvEmeaUnauthorized {
	return &DirectDebitInquiryV1FvEmeaUnauthorized{}
}

/*DirectDebitInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type DirectDebitInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitInquiryV1FvEmeaForbidden creates a DirectDebitInquiryV1FvEmeaForbidden with default headers values
func NewDirectDebitInquiryV1FvEmeaForbidden() *DirectDebitInquiryV1FvEmeaForbidden {
	return &DirectDebitInquiryV1FvEmeaForbidden{}
}

/*DirectDebitInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type DirectDebitInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitInquiryV1FvEmeaNotFound creates a DirectDebitInquiryV1FvEmeaNotFound with default headers values
func NewDirectDebitInquiryV1FvEmeaNotFound() *DirectDebitInquiryV1FvEmeaNotFound {
	return &DirectDebitInquiryV1FvEmeaNotFound{}
}

/*DirectDebitInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type DirectDebitInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitInquiryV1FvEmeaTooManyRequests creates a DirectDebitInquiryV1FvEmeaTooManyRequests with default headers values
func NewDirectDebitInquiryV1FvEmeaTooManyRequests() *DirectDebitInquiryV1FvEmeaTooManyRequests {
	return &DirectDebitInquiryV1FvEmeaTooManyRequests{}
}

/*DirectDebitInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type DirectDebitInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectDebitInquiryV1FvEmeaStatus452 creates a DirectDebitInquiryV1FvEmeaStatus452 with default headers values
func NewDirectDebitInquiryV1FvEmeaStatus452() *DirectDebitInquiryV1FvEmeaStatus452 {
	return &DirectDebitInquiryV1FvEmeaStatus452{}
}

/*DirectDebitInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type DirectDebitInquiryV1FvEmeaStatus452 struct {
}

func (o *DirectDebitInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaStatus452 ", 452)
}

func (o *DirectDebitInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDirectDebitInquiryV1FvEmeaStatus453 creates a DirectDebitInquiryV1FvEmeaStatus453 with default headers values
func NewDirectDebitInquiryV1FvEmeaStatus453() *DirectDebitInquiryV1FvEmeaStatus453 {
	return &DirectDebitInquiryV1FvEmeaStatus453{}
}

/*DirectDebitInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type DirectDebitInquiryV1FvEmeaStatus453 struct {
}

func (o *DirectDebitInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaStatus453 ", 453)
}

func (o *DirectDebitInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDirectDebitInquiryV1FvEmeaStatus455 creates a DirectDebitInquiryV1FvEmeaStatus455 with default headers values
func NewDirectDebitInquiryV1FvEmeaStatus455() *DirectDebitInquiryV1FvEmeaStatus455 {
	return &DirectDebitInquiryV1FvEmeaStatus455{}
}

/*DirectDebitInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type DirectDebitInquiryV1FvEmeaStatus455 struct {
}

func (o *DirectDebitInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaStatus455 ", 455)
}

func (o *DirectDebitInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDirectDebitInquiryV1FvEmeaStatus465 creates a DirectDebitInquiryV1FvEmeaStatus465 with default headers values
func NewDirectDebitInquiryV1FvEmeaStatus465() *DirectDebitInquiryV1FvEmeaStatus465 {
	return &DirectDebitInquiryV1FvEmeaStatus465{}
}

/*DirectDebitInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SDI01E : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 000-998<BR/>VPL5SDI01S : REQUESTED CARDorACCOUNT NUMBER IS NOT FOUND<BR/>VPL5SDI02E : ACCOUNT OR CARD NUMBER REQUIRED<BR/>VPL5SDI02S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SDI03E : INVALID FOREIGN USE INDICATOR - VALID VALUES ARE SPACE 0 OR 1   INVALID FOREIGN USE INDICATOR SPECIFIED - VALID VALUES ARE       SPACE 0 OR 1<BR/>VPL5SDI03S : FOREIGN ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SDI04S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SDI06S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SDIZ1S : SERVICE INPUT TO DIRECT DEBIT SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SDIZ3S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SDIZ4S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPL5SDIZ5S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SDIZ6S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SDIZ7S : ORGANIZATION FILE UNAVAILABLE<BR/>VPL5SDIZ8S : ACCOUNT BASE SEGMENT FILE UNAVAILABL
*/
type DirectDebitInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *DirectDebitInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/directDebitInquiry][%d] directDebitInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *DirectDebitInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *DirectDebitInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
