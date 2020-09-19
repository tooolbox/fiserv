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

// EmbosserUpdateV1FvEmeaReader is a Reader for the EmbosserUpdateV1FvEmea structure.
type EmbosserUpdateV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmbosserUpdateV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmbosserUpdateV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEmbosserUpdateV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEmbosserUpdateV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEmbosserUpdateV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEmbosserUpdateV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewEmbosserUpdateV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewEmbosserUpdateV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewEmbosserUpdateV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewEmbosserUpdateV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewEmbosserUpdateV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEmbosserUpdateV1FvEmeaOK creates a EmbosserUpdateV1FvEmeaOK with default headers values
func NewEmbosserUpdateV1FvEmeaOK() *EmbosserUpdateV1FvEmeaOK {
	return &EmbosserUpdateV1FvEmeaOK{}
}

/*EmbosserUpdateV1FvEmeaOK handles this case with default header values.

successful operation
*/
type EmbosserUpdateV1FvEmeaOK struct {
	Payload *models.EmbosserUpdateResponse
}

func (o *EmbosserUpdateV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaOK) GetPayload() *models.EmbosserUpdateResponse {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmbosserUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmbosserUpdateV1FvEmeaBadRequest creates a EmbosserUpdateV1FvEmeaBadRequest with default headers values
func NewEmbosserUpdateV1FvEmeaBadRequest() *EmbosserUpdateV1FvEmeaBadRequest {
	return &EmbosserUpdateV1FvEmeaBadRequest{}
}

/*EmbosserUpdateV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type EmbosserUpdateV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *EmbosserUpdateV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmbosserUpdateV1FvEmeaUnauthorized creates a EmbosserUpdateV1FvEmeaUnauthorized with default headers values
func NewEmbosserUpdateV1FvEmeaUnauthorized() *EmbosserUpdateV1FvEmeaUnauthorized {
	return &EmbosserUpdateV1FvEmeaUnauthorized{}
}

/*EmbosserUpdateV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type EmbosserUpdateV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *EmbosserUpdateV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmbosserUpdateV1FvEmeaForbidden creates a EmbosserUpdateV1FvEmeaForbidden with default headers values
func NewEmbosserUpdateV1FvEmeaForbidden() *EmbosserUpdateV1FvEmeaForbidden {
	return &EmbosserUpdateV1FvEmeaForbidden{}
}

/*EmbosserUpdateV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type EmbosserUpdateV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *EmbosserUpdateV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmbosserUpdateV1FvEmeaNotFound creates a EmbosserUpdateV1FvEmeaNotFound with default headers values
func NewEmbosserUpdateV1FvEmeaNotFound() *EmbosserUpdateV1FvEmeaNotFound {
	return &EmbosserUpdateV1FvEmeaNotFound{}
}

/*EmbosserUpdateV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type EmbosserUpdateV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *EmbosserUpdateV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmbosserUpdateV1FvEmeaTooManyRequests creates a EmbosserUpdateV1FvEmeaTooManyRequests with default headers values
func NewEmbosserUpdateV1FvEmeaTooManyRequests() *EmbosserUpdateV1FvEmeaTooManyRequests {
	return &EmbosserUpdateV1FvEmeaTooManyRequests{}
}

/*EmbosserUpdateV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type EmbosserUpdateV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *EmbosserUpdateV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmbosserUpdateV1FvEmeaStatus452 creates a EmbosserUpdateV1FvEmeaStatus452 with default headers values
func NewEmbosserUpdateV1FvEmeaStatus452() *EmbosserUpdateV1FvEmeaStatus452 {
	return &EmbosserUpdateV1FvEmeaStatus452{}
}

/*EmbosserUpdateV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type EmbosserUpdateV1FvEmeaStatus452 struct {
}

func (o *EmbosserUpdateV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaStatus452 ", 452)
}

func (o *EmbosserUpdateV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmbosserUpdateV1FvEmeaStatus453 creates a EmbosserUpdateV1FvEmeaStatus453 with default headers values
func NewEmbosserUpdateV1FvEmeaStatus453() *EmbosserUpdateV1FvEmeaStatus453 {
	return &EmbosserUpdateV1FvEmeaStatus453{}
}

/*EmbosserUpdateV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type EmbosserUpdateV1FvEmeaStatus453 struct {
}

func (o *EmbosserUpdateV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaStatus453 ", 453)
}

func (o *EmbosserUpdateV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmbosserUpdateV1FvEmeaStatus455 creates a EmbosserUpdateV1FvEmeaStatus455 with default headers values
func NewEmbosserUpdateV1FvEmeaStatus455() *EmbosserUpdateV1FvEmeaStatus455 {
	return &EmbosserUpdateV1FvEmeaStatus455{}
}

/*EmbosserUpdateV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type EmbosserUpdateV1FvEmeaStatus455 struct {
}

func (o *EmbosserUpdateV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaStatus455 ", 455)
}

func (o *EmbosserUpdateV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmbosserUpdateV1FvEmeaStatus465 creates a EmbosserUpdateV1FvEmeaStatus465 with default headers values
func NewEmbosserUpdateV1FvEmeaStatus465() *EmbosserUpdateV1FvEmeaStatus465 {
	return &EmbosserUpdateV1FvEmeaStatus465{}
}

/*EmbosserUpdateV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SEUZ1S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SEUZ2S : SERVICE INPUT TO CARD UPDATE SERVICE IS AN INCORRECT LENGTH<BR/>VPL5SEUZ4S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SEU14S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SEU01E : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 000-998<BR/>VPL5SEU02E : CARD NUMBER IS NOT NUMERIC OR EQUALS ZEROES<BR/>VPL5SEU05E : CARD SEQUENCE NUMBER MUST BE NUMERIC VALID VALUES ARE 0001-9998<BR/>VPL5SEU03E : INVALID FOREIGN USE INDICATOR VALID VALUES ARE SPACE0 OR 1<BR/>VPL5SEU38E : INVALID REQUEST FOR PIN RTY CNT RESET- 0 OR 1 ONLY<BR/>VPL5SEU04S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SEUZ5S : REQUESTED ORG NUMBER IS NOT FOUND<BR/>VPL5SEUZ6S : SYSTEM RECORD NOT FOUND<BR/>VPL5SEUZ7S : ORGANIZATION FILE UNAVAILABLE<BR/>VPL5SEU01S : ORGANIZATION NUMBER NOT FOUND<BR/>VPL5SEU16I : NO FOREIGN ORG FOUND IN ORG RECORD<BR/>VPL5SEUZ8S : EMBOSSER CARD FILE UNAVAILABLE<BR/>VPL5SEU02S : CARD NUMBER NOT FOUND<BR/>VPL5SEUZ9S : ACCOUNT BASE SEGMENT FILE UNAVAILABLE<BR/>VPL5SEU03S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SEUZCS : LOGO RECORD NOT FOUND<BR/>VPL5SEUZBS : CUSTOMER NAME AND ADDRESS FILE UNAVAILABLE<BR/>VPL5SEU15S : CUSTOMER NAME AND ADDRESS NOT FOUND<BR/>VPL5SEU21E : EMBOSSER NAME 1 CANNOT BE SPACES<BR/>VPL5SEU08E : USER 4 FIELD IS NOT NUMERIC<BR/>VPL5SEU09E : USER DATE 1 IS NOT NUMERIC<BR/>VPL5SEU10E : USER DATE 1 IS NOT A VALID DATE<BR/>VPL5SEU11E : USER DATE 2 IS NOT NUMERIC<BR/>VPL5SEU12E : USER DATE 2 IS NOT A VALID DATE<BR/>VPL5SEU13E : CARDHOLDER FLAG INVALID VALID VALUES ARE SPACE 0 TO 9<BR/>VPL5SEU17E : LANG CODE INVALID.VALID VALUES ARE ALPHA-3 CODES FROM SS40WS<BR/>VPL5SEU18E : INVALID ISO LANGUAGE CODE<BR/>VPL5SEU19E : ISSUE DELIVERY OPTION IS NOT NUMERIC<BR/>VPL5SEU20E : REISSUE DELIVERY OPTION IS NOT NUMERIC<BR/>VPL5SEU26E : EMBLEM ID IS NOT NUMERIC<BR/>VPL5SEU27E : POS SERVICE CODE IS NOT NUMERIC<BR/>VPL5SEU30E : PIN COMM PREF METHOD NOT ALLOWED FOR PRODUCT<BR/>VPL5SEU29E : PIN COMM PREF FIELD MUST BE EITHER 0 1 OR 2<BR/>VPL5SEU32E : PIN OVRD METHOD NOT ALLOWED FOR PRODUCT<BR/>VPL5SEU31E : PIN COMM OVRD FIELD MUST BE EITHER 0 1 2 OR 4<BR/>VPL5SEU33E : PIN COMM PREF OVRD CANNOT BE THE SAME AS PIN COMM PREF<BR/>VPL5SEU34E : MOBILE NUMBER HAS INVALID CHARACTERS<BR/>VPL5SEU35E : MEMORABLE WORD HAS INVALID CHARACTERS<BR/>VPL5SEU36E : MOBILE NUMBER IS MANDATORY FOR PIN BY SMS<BR/>VPL5SEU37E : MEMORABLE WORD IS MANDATORY FOR PIN BY SMS<BR/>VPL5SEU49E : VALID VALUES FOR AUTH ATM OTC RTL FREQ IS NUMERIC 0 1 2 OR 3<BR/>VPL5SEU40E : AUTH ATM CASH AMOUNT IS NOT NUMERIC<BR/>VPL5SEU41E : AUTH ATM CASH NUMBER IS NOT NUMERIC<BR/>VPL5SEU42E : AUTH ATM CASH SINGLE TXN IS NOT NUMERIC<BR/>VPL5SEU43E : AUTH OTC CASH AMOUNT IS NOT NUMERIC<BR/>VPL5SEU70E : AUTH OTC CASH NUMBER IS NOT NUMERIC<BR/>VPL5SEU45E : AUTH OTC CASH SINGLE IS NOT NUMERIC<BR/>VPL5SEU46E : AUTH RETAIL AMOUNT IS NOT NUMERIC<BR/>VPL5SEU47E : AUTH RETAIL NUMBER IS NOT NUMERIC<BR/>VPL5SEU48E : AUTH RETAIL SINGLE TXN IS NOT NUMERIC<BR/>VPL5SEU71E : AUTH TOTAL TXN AMOUNT IS NOT NUMERIC<BR/>VPL5SEU72E : AUTH TOTAL TXN NUMBER IS NOT NUMERIC<BR/>VPL5SEUZDS : BASE SEGMENT EMBOSSER CROSS-REFERENCE FILE UNAVAILABLE<BR/>VPL5SEU73E : TXN EOMM LIMIT IS NOT NUMERIC<BR/>VPL5SEU74E : TYPE OF CARD IS NOT NUMERIC<BR/>VPL5SEU75E : REQUESTED CARD TYPE IS NOT NUMERIC<BR/>VPL5SEU76E : INPUT SDP PROCESS REQUEST IS NOT NUMERIC<BR/>VPL5SEU77E : INVALID SDP REQUESTVALID REQUEST IS 0 1 3 5 7 OR 9<BR/>VPL5SEU78E : EMBLM ID IS NOT NUMERIC<BR/>VPL5SEU79E : ECOMM NUMBER IS NOT NUMERIC<BR/>VPL5SEU80E : ECOMM AMOUNT IS NOT NUMERIC<BR/>VPL5SEU81E : AUTH ECOMM AMOUNT ACCUMULATION IS NOT NUMERIC<BR/>VPL5SEU82E : AUTH ECOMM NUMBER ACCUMULATION IS NOT NUMERIC<BR/>VPL5SEU83E : SDP CYCLE IS IN PROGRES; UPDATE FOR TYPE OF CARD NOT ALLOWE
*/
type EmbosserUpdateV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *EmbosserUpdateV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/embosserUpdate][%d] embosserUpdateV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *EmbosserUpdateV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *EmbosserUpdateV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
