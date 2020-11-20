// Code generated by go-swagger; DO NOT EDIT.

package loyalty_and_rewards_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// LoyaltyStatementSettingsUpdateV1FvEmeaReader is a Reader for the LoyaltyStatementSettingsUpdateV1FvEmea structure.
type LoyaltyStatementSettingsUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoyaltyStatementSettingsUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaOK creates a LoyaltyStatementSettingsUpdateV1FvEmeaOK with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaOK() *LoyaltyStatementSettingsUpdateV1FvEmeaOK {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaOK{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaOK struct {
	Payload *models.LoyaltyStatementSettingsUpdateResponse
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaOK) GetPayload() *models.LoyaltyStatementSettingsUpdateResponse {
	return o.Payload
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoyaltyStatementSettingsUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaBadRequest creates a LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaBadRequest() *LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized creates a LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized() *LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaForbidden creates a LoyaltyStatementSettingsUpdateV1FvEmeaForbidden with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaForbidden() *LoyaltyStatementSettingsUpdateV1FvEmeaForbidden {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaForbidden{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaNotFound creates a LoyaltyStatementSettingsUpdateV1FvEmeaNotFound with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaNotFound() *LoyaltyStatementSettingsUpdateV1FvEmeaNotFound {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaNotFound{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests creates a LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests() *LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus452 creates a LoyaltyStatementSettingsUpdateV1FvEmeaStatus452 with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus452() *LoyaltyStatementSettingsUpdateV1FvEmeaStatus452 {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaStatus452{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaStatus452 struct {
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaStatus452 ", 452)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus453 creates a LoyaltyStatementSettingsUpdateV1FvEmeaStatus453 with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus453() *LoyaltyStatementSettingsUpdateV1FvEmeaStatus453 {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaStatus453{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaStatus453 struct {
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaStatus453 ", 453)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus455 creates a LoyaltyStatementSettingsUpdateV1FvEmeaStatus455 with default headers values
func NewLoyaltyStatementSettingsUpdateV1FvEmeaStatus455() *LoyaltyStatementSettingsUpdateV1FvEmeaStatus455 {
	return &LoyaltyStatementSettingsUpdateV1FvEmeaStatus455{}
}

/*LoyaltyStatementSettingsUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LoyaltyStatementSettingsUpdateV1FvEmeaStatus455 struct {
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyStatementSettingsUpdate][%d] loyaltyStatementSettingsUpdateV1FvEmeaStatus455 ", 455)
}

func (o *LoyaltyStatementSettingsUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}