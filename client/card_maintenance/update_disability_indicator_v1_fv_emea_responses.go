// Code generated by go-swagger; DO NOT EDIT.

package card_maintenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// UpdateDisabilityIndicatorV1FvEmeaReader is a Reader for the UpdateDisabilityIndicatorV1FvEmea structure.
type UpdateDisabilityIndicatorV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDisabilityIndicatorV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDisabilityIndicatorV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDisabilityIndicatorV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateDisabilityIndicatorV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDisabilityIndicatorV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDisabilityIndicatorV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateDisabilityIndicatorV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewUpdateDisabilityIndicatorV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewUpdateDisabilityIndicatorV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewUpdateDisabilityIndicatorV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewUpdateDisabilityIndicatorV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDisabilityIndicatorV1FvEmeaOK creates a UpdateDisabilityIndicatorV1FvEmeaOK with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaOK() *UpdateDisabilityIndicatorV1FvEmeaOK {
	return &UpdateDisabilityIndicatorV1FvEmeaOK{}
}

/*UpdateDisabilityIndicatorV1FvEmeaOK handles this case with default header values.

successful operation
*/
type UpdateDisabilityIndicatorV1FvEmeaOK struct {
	Payload *models.UpdateDisabilityIndicatorResponse
}

func (o *UpdateDisabilityIndicatorV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaOK) GetPayload() *models.UpdateDisabilityIndicatorResponse {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDisabilityIndicatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaBadRequest creates a UpdateDisabilityIndicatorV1FvEmeaBadRequest with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaBadRequest() *UpdateDisabilityIndicatorV1FvEmeaBadRequest {
	return &UpdateDisabilityIndicatorV1FvEmeaBadRequest{}
}

/*UpdateDisabilityIndicatorV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type UpdateDisabilityIndicatorV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *UpdateDisabilityIndicatorV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaUnauthorized creates a UpdateDisabilityIndicatorV1FvEmeaUnauthorized with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaUnauthorized() *UpdateDisabilityIndicatorV1FvEmeaUnauthorized {
	return &UpdateDisabilityIndicatorV1FvEmeaUnauthorized{}
}

/*UpdateDisabilityIndicatorV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type UpdateDisabilityIndicatorV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *UpdateDisabilityIndicatorV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaForbidden creates a UpdateDisabilityIndicatorV1FvEmeaForbidden with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaForbidden() *UpdateDisabilityIndicatorV1FvEmeaForbidden {
	return &UpdateDisabilityIndicatorV1FvEmeaForbidden{}
}

/*UpdateDisabilityIndicatorV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type UpdateDisabilityIndicatorV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *UpdateDisabilityIndicatorV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaNotFound creates a UpdateDisabilityIndicatorV1FvEmeaNotFound with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaNotFound() *UpdateDisabilityIndicatorV1FvEmeaNotFound {
	return &UpdateDisabilityIndicatorV1FvEmeaNotFound{}
}

/*UpdateDisabilityIndicatorV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type UpdateDisabilityIndicatorV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *UpdateDisabilityIndicatorV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaTooManyRequests creates a UpdateDisabilityIndicatorV1FvEmeaTooManyRequests with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaTooManyRequests() *UpdateDisabilityIndicatorV1FvEmeaTooManyRequests {
	return &UpdateDisabilityIndicatorV1FvEmeaTooManyRequests{}
}

/*UpdateDisabilityIndicatorV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type UpdateDisabilityIndicatorV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *UpdateDisabilityIndicatorV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaStatus452 creates a UpdateDisabilityIndicatorV1FvEmeaStatus452 with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaStatus452() *UpdateDisabilityIndicatorV1FvEmeaStatus452 {
	return &UpdateDisabilityIndicatorV1FvEmeaStatus452{}
}

/*UpdateDisabilityIndicatorV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type UpdateDisabilityIndicatorV1FvEmeaStatus452 struct {
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaStatus452 ", 452)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaStatus453 creates a UpdateDisabilityIndicatorV1FvEmeaStatus453 with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaStatus453() *UpdateDisabilityIndicatorV1FvEmeaStatus453 {
	return &UpdateDisabilityIndicatorV1FvEmeaStatus453{}
}

/*UpdateDisabilityIndicatorV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type UpdateDisabilityIndicatorV1FvEmeaStatus453 struct {
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaStatus453 ", 453)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaStatus455 creates a UpdateDisabilityIndicatorV1FvEmeaStatus455 with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaStatus455() *UpdateDisabilityIndicatorV1FvEmeaStatus455 {
	return &UpdateDisabilityIndicatorV1FvEmeaStatus455{}
}

/*UpdateDisabilityIndicatorV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type UpdateDisabilityIndicatorV1FvEmeaStatus455 struct {
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaStatus455 ", 455)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDisabilityIndicatorV1FvEmeaStatus465 creates a UpdateDisabilityIndicatorV1FvEmeaStatus465 with default headers values
func NewUpdateDisabilityIndicatorV1FvEmeaStatus465() *UpdateDisabilityIndicatorV1FvEmeaStatus465 {
	return &UpdateDisabilityIndicatorV1FvEmeaStatus465{}
}

/*UpdateDisabilityIndicatorV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SDE01E : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 001-998<BR/>VPL5SDE01S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SDE02E : CARD NUMBER IS NOT NUMERIC OR EQUALS ZEROES<BR/>VPL5SDE02S : CARD NUMBER NOT FOUND<BR/>VPL5SDE03S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SDE05E : CARD SEQUENCE NUMBER MUST BE NUMERIC VALID VALUES ARE 0001-9998<BR/>VPL5SDE06E : INVALID DISABILITY INDICATOR VALUE<BR/>VPL5SDE14S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SDEZ1S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SDEZ2S : SERVICE INPUT TO CARD UPDATE SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SDEZ4S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SDEZ5S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPL5SDEZ6S : SYSTEM RECORD NOT FOUND<BR/>VPL5SDEZ7S : ORGANIZATION FILE UNAVAILABLE<BR/>VPL5SDEZ8S : EMBOSSER CARD FILE UNAVAILABLE<BR/>VPL5SDEZ9S : ACCOUNT BASE SEGMENT FILE UNAVAILABLE<BR/>VPL5SDEZAS : ORGANIZATION NUMBER NOT DETERMINE
*/
type UpdateDisabilityIndicatorV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/updateDisabilityIndicator][%d] updateDisabilityIndicatorV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UpdateDisabilityIndicatorV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}