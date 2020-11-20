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

// InstantCreditFlagUpdateV1FvEmeaReader is a Reader for the InstantCreditFlagUpdateV1FvEmea structure.
type InstantCreditFlagUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstantCreditFlagUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstantCreditFlagUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstantCreditFlagUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInstantCreditFlagUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstantCreditFlagUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstantCreditFlagUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInstantCreditFlagUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewInstantCreditFlagUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewInstantCreditFlagUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewInstantCreditFlagUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewInstantCreditFlagUpdateV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstantCreditFlagUpdateV1FvEmeaOK creates a InstantCreditFlagUpdateV1FvEmeaOK with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaOK() *InstantCreditFlagUpdateV1FvEmeaOK {
	return &InstantCreditFlagUpdateV1FvEmeaOK{}
}

/*InstantCreditFlagUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type InstantCreditFlagUpdateV1FvEmeaOK struct {
	Payload *models.InstantCreditFlagUpdateResponse
}

func (o *InstantCreditFlagUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaOK) GetPayload() *models.InstantCreditFlagUpdateResponse {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstantCreditFlagUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaBadRequest creates a InstantCreditFlagUpdateV1FvEmeaBadRequest with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaBadRequest() *InstantCreditFlagUpdateV1FvEmeaBadRequest {
	return &InstantCreditFlagUpdateV1FvEmeaBadRequest{}
}

/*InstantCreditFlagUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type InstantCreditFlagUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *InstantCreditFlagUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaUnauthorized creates a InstantCreditFlagUpdateV1FvEmeaUnauthorized with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaUnauthorized() *InstantCreditFlagUpdateV1FvEmeaUnauthorized {
	return &InstantCreditFlagUpdateV1FvEmeaUnauthorized{}
}

/*InstantCreditFlagUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type InstantCreditFlagUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *InstantCreditFlagUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaForbidden creates a InstantCreditFlagUpdateV1FvEmeaForbidden with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaForbidden() *InstantCreditFlagUpdateV1FvEmeaForbidden {
	return &InstantCreditFlagUpdateV1FvEmeaForbidden{}
}

/*InstantCreditFlagUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type InstantCreditFlagUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *InstantCreditFlagUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaNotFound creates a InstantCreditFlagUpdateV1FvEmeaNotFound with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaNotFound() *InstantCreditFlagUpdateV1FvEmeaNotFound {
	return &InstantCreditFlagUpdateV1FvEmeaNotFound{}
}

/*InstantCreditFlagUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type InstantCreditFlagUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *InstantCreditFlagUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaTooManyRequests creates a InstantCreditFlagUpdateV1FvEmeaTooManyRequests with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaTooManyRequests() *InstantCreditFlagUpdateV1FvEmeaTooManyRequests {
	return &InstantCreditFlagUpdateV1FvEmeaTooManyRequests{}
}

/*InstantCreditFlagUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type InstantCreditFlagUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *InstantCreditFlagUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaStatus452 creates a InstantCreditFlagUpdateV1FvEmeaStatus452 with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaStatus452() *InstantCreditFlagUpdateV1FvEmeaStatus452 {
	return &InstantCreditFlagUpdateV1FvEmeaStatus452{}
}

/*InstantCreditFlagUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type InstantCreditFlagUpdateV1FvEmeaStatus452 struct {
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaStatus452 ", 452)
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaStatus453 creates a InstantCreditFlagUpdateV1FvEmeaStatus453 with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaStatus453() *InstantCreditFlagUpdateV1FvEmeaStatus453 {
	return &InstantCreditFlagUpdateV1FvEmeaStatus453{}
}

/*InstantCreditFlagUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type InstantCreditFlagUpdateV1FvEmeaStatus453 struct {
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaStatus453 ", 453)
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaStatus455 creates a InstantCreditFlagUpdateV1FvEmeaStatus455 with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaStatus455() *InstantCreditFlagUpdateV1FvEmeaStatus455 {
	return &InstantCreditFlagUpdateV1FvEmeaStatus455{}
}

/*InstantCreditFlagUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type InstantCreditFlagUpdateV1FvEmeaStatus455 struct {
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaStatus455 ", 455)
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantCreditFlagUpdateV1FvEmeaStatus465 creates a InstantCreditFlagUpdateV1FvEmeaStatus465 with default headers values
func NewInstantCreditFlagUpdateV1FvEmeaStatus465() *InstantCreditFlagUpdateV1FvEmeaStatus465 {
	return &InstantCreditFlagUpdateV1FvEmeaStatus465{}
}

/*InstantCreditFlagUpdateV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SIFZ1S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SIFZ2S : SERVICE INPUT TO INSTANT CREDIT SERVICE  IS AN INCORRECT LENGTH<BR/>VPL5SIFZ3S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SIF02S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SIF03S : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 000-998<BR/>VPL5SIF04S : REQUESTED ORGANIZATION COULD NOT BE DETERMINED<BR/>VPL5SIF05S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPL5SIF06S : ORGANIZATION NUMBER NOT FOUND ON FILE<BR/>VPL5SIF07S : ACCOUNT NUMBER IS NOT NUMERIC OR EQUALS SPACES<BR/>VPL5SIF08S : ACCOUNT NBR MUST BE GREATER THAN ZEROES<BR/>VPL5SIF09S : ACCOUNT NOT FOUND ON FILE<BR/>VPL5SIF10S : INVALID ACCOUNT BASE SEGMNENT RECORD<BR/>VPL5SIF11S : INSTANT CREDIT IS NOT ACTIVE FOR THIS LOGO<BR/>VPL5SIF12S : INSTANT CREDIT FLAG IS NOT NUMERIC<BR/>VPL5SIF13S : VALID VALUES ARE 0 AND 9<BR/>VPL5SIF14S : INSTANT CREDIT CAN'T BE CHANGED TO 0 CURRENT FLAG HAS A VALUE 9<BR/>VPL5SIF15S : INSTANT CREDIT CAN'T BE CHANGED TO 9 CURRENT FLAG IS NOT SET AS 1<BR/>VPL5SIF16S : CURRENT INSTANT CREDIT NUMBER NOT FOUND ON FIL
*/
type InstantCreditFlagUpdateV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantCreditFlagUpdate][%d] instantCreditFlagUpdateV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantCreditFlagUpdateV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}