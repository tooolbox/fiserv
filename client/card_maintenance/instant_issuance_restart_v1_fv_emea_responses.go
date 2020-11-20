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

// InstantIssuanceRestartV1FvEmeaReader is a Reader for the InstantIssuanceRestartV1FvEmea structure.
type InstantIssuanceRestartV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstantIssuanceRestartV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstantIssuanceRestartV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstantIssuanceRestartV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInstantIssuanceRestartV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstantIssuanceRestartV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstantIssuanceRestartV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInstantIssuanceRestartV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewInstantIssuanceRestartV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewInstantIssuanceRestartV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewInstantIssuanceRestartV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewInstantIssuanceRestartV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstantIssuanceRestartV1FvEmeaOK creates a InstantIssuanceRestartV1FvEmeaOK with default headers values
func NewInstantIssuanceRestartV1FvEmeaOK() *InstantIssuanceRestartV1FvEmeaOK {
	return &InstantIssuanceRestartV1FvEmeaOK{}
}

/*InstantIssuanceRestartV1FvEmeaOK handles this case with default header values.

successful operation
*/
type InstantIssuanceRestartV1FvEmeaOK struct {
	Payload *models.InstantIssuanceRestartResponse
}

func (o *InstantIssuanceRestartV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaOK) GetPayload() *models.InstantIssuanceRestartResponse {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstantIssuanceRestartResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaBadRequest creates a InstantIssuanceRestartV1FvEmeaBadRequest with default headers values
func NewInstantIssuanceRestartV1FvEmeaBadRequest() *InstantIssuanceRestartV1FvEmeaBadRequest {
	return &InstantIssuanceRestartV1FvEmeaBadRequest{}
}

/*InstantIssuanceRestartV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type InstantIssuanceRestartV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *InstantIssuanceRestartV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaUnauthorized creates a InstantIssuanceRestartV1FvEmeaUnauthorized with default headers values
func NewInstantIssuanceRestartV1FvEmeaUnauthorized() *InstantIssuanceRestartV1FvEmeaUnauthorized {
	return &InstantIssuanceRestartV1FvEmeaUnauthorized{}
}

/*InstantIssuanceRestartV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type InstantIssuanceRestartV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *InstantIssuanceRestartV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaForbidden creates a InstantIssuanceRestartV1FvEmeaForbidden with default headers values
func NewInstantIssuanceRestartV1FvEmeaForbidden() *InstantIssuanceRestartV1FvEmeaForbidden {
	return &InstantIssuanceRestartV1FvEmeaForbidden{}
}

/*InstantIssuanceRestartV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type InstantIssuanceRestartV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *InstantIssuanceRestartV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaNotFound creates a InstantIssuanceRestartV1FvEmeaNotFound with default headers values
func NewInstantIssuanceRestartV1FvEmeaNotFound() *InstantIssuanceRestartV1FvEmeaNotFound {
	return &InstantIssuanceRestartV1FvEmeaNotFound{}
}

/*InstantIssuanceRestartV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type InstantIssuanceRestartV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *InstantIssuanceRestartV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaTooManyRequests creates a InstantIssuanceRestartV1FvEmeaTooManyRequests with default headers values
func NewInstantIssuanceRestartV1FvEmeaTooManyRequests() *InstantIssuanceRestartV1FvEmeaTooManyRequests {
	return &InstantIssuanceRestartV1FvEmeaTooManyRequests{}
}

/*InstantIssuanceRestartV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type InstantIssuanceRestartV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *InstantIssuanceRestartV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaStatus452 creates a InstantIssuanceRestartV1FvEmeaStatus452 with default headers values
func NewInstantIssuanceRestartV1FvEmeaStatus452() *InstantIssuanceRestartV1FvEmeaStatus452 {
	return &InstantIssuanceRestartV1FvEmeaStatus452{}
}

/*InstantIssuanceRestartV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type InstantIssuanceRestartV1FvEmeaStatus452 struct {
}

func (o *InstantIssuanceRestartV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaStatus452 ", 452)
}

func (o *InstantIssuanceRestartV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaStatus453 creates a InstantIssuanceRestartV1FvEmeaStatus453 with default headers values
func NewInstantIssuanceRestartV1FvEmeaStatus453() *InstantIssuanceRestartV1FvEmeaStatus453 {
	return &InstantIssuanceRestartV1FvEmeaStatus453{}
}

/*InstantIssuanceRestartV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type InstantIssuanceRestartV1FvEmeaStatus453 struct {
}

func (o *InstantIssuanceRestartV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaStatus453 ", 453)
}

func (o *InstantIssuanceRestartV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaStatus455 creates a InstantIssuanceRestartV1FvEmeaStatus455 with default headers values
func NewInstantIssuanceRestartV1FvEmeaStatus455() *InstantIssuanceRestartV1FvEmeaStatus455 {
	return &InstantIssuanceRestartV1FvEmeaStatus455{}
}

/*InstantIssuanceRestartV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type InstantIssuanceRestartV1FvEmeaStatus455 struct {
}

func (o *InstantIssuanceRestartV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaStatus455 ", 455)
}

func (o *InstantIssuanceRestartV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantIssuanceRestartV1FvEmeaStatus465 creates a InstantIssuanceRestartV1FvEmeaStatus465 with default headers values
func NewInstantIssuanceRestartV1FvEmeaStatus465() *InstantIssuanceRestartV1FvEmeaStatus465 {
	return &InstantIssuanceRestartV1FvEmeaStatus465{}
}

/*InstantIssuanceRestartV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SERI1S : INSTANT ISSUANCE FUNCTIONALITY INACTIVE AT THE LOGO LEVEL<BR/>VPL5SERI2S : INSTANT ISSUANCE REQUESTS EXCEEDED FOR THE DAY<BR/>VPL5SERI3S : PLASTIC NOT TO BE GENERATED<BR/>VPL5SERI4S : INSTANT ISSUANCE NOT SUPPORTED FOR MULTI SCHEME<BR/>VPL5SERI5S : INSTANT ISSUANCE NOT SUPPORTED FOR PREPAID CARDS<BR/>VPL5SERI6S : INSTANT ISSUANCE NOT SUPPORTED WHEN HCS IS ACTIVE<BR/>VPL5SERI7S : INSTANT ISSUANCE NOT SUPPORTED FOR TRIAD CONTROLLED ACCOUNTS<BR/>VPL5SERI8S : INSTANT ISSUANCE SUPPORTED FOR CARD SCHEME 3 ONLY<BR/>VPL5SER41S : INVALID ORG ENTERED VALID VALUES ARE 001-998<BR/>VPL5SER42S : ACCOUNT NOT FOUND<BR/>VPL5SER43S : CARD NBR REQUIRED<BR/>VPL5SER44S : EMBOSSER NOT FOUND<BR/>VPL5SER45S : EMBOSSER PENDING ADD<BR/>VPL5SER46S : CARD SEQ INVALID VALID VALUES ARE 0001-9998<BR/>VPL5SER47S : REQUESTED ORGANIZATION COULD NOT BE DETERMINED<BR/>VPL5SER48S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPL5SER49S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SER50S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SER51S : SYSTEM RECORD NOT FOUND<BR/>VPL5SER52S : ORG RECORD NOT FOUND<BR/>VPL5SER53S : ORG RECORD PENDING ADD<BR/>VPL5SER54S : LOGO RECORD NOT FOUND<BR/>VPL5SER55S : LOGO RECORD PENDING ADD<BR/>VPL5SER56S : SMART CARD EMBOSSER NOT PRESENT<BR/>VPL5SER57S : RELATIONSHIP RECORD NOT FOUND<BR/>VPL5SER58S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SER59S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SERF2S : AMCR FILE NOT OPEN<BR/>VPL5SERF4S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SERF8S : AMED FILE NOT OPEN<BR/>VPL5SERF9S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SERFAS : SERVICE INPUT TO INSTANT ISS RESTART SERVICE IS INCORRECT LENGTH<BR/>VPL5SER91S : SYSTEM NOT IN NORMAL PROCESSIN
*/
type InstantIssuanceRestartV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *InstantIssuanceRestartV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/instantIssuanceRestart][%d] instantIssuanceRestartV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *InstantIssuanceRestartV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InstantIssuanceRestartV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}