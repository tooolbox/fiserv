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

// OfferAddV1FvEmeaReader is a Reader for the OfferAddV1FvEmea structure.
type OfferAddV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OfferAddV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOfferAddV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOfferAddV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOfferAddV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOfferAddV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOfferAddV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOfferAddV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewOfferAddV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewOfferAddV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewOfferAddV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewOfferAddV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOfferAddV1FvEmeaOK creates a OfferAddV1FvEmeaOK with default headers values
func NewOfferAddV1FvEmeaOK() *OfferAddV1FvEmeaOK {
	return &OfferAddV1FvEmeaOK{}
}

/*OfferAddV1FvEmeaOK handles this case with default header values.

successful operation
*/
type OfferAddV1FvEmeaOK struct {
	Payload *models.OfferAddResponse
}

func (o *OfferAddV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *OfferAddV1FvEmeaOK) GetPayload() *models.OfferAddResponse {
	return o.Payload
}

func (o *OfferAddV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OfferAddResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferAddV1FvEmeaBadRequest creates a OfferAddV1FvEmeaBadRequest with default headers values
func NewOfferAddV1FvEmeaBadRequest() *OfferAddV1FvEmeaBadRequest {
	return &OfferAddV1FvEmeaBadRequest{}
}

/*OfferAddV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type OfferAddV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *OfferAddV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *OfferAddV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferAddV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferAddV1FvEmeaUnauthorized creates a OfferAddV1FvEmeaUnauthorized with default headers values
func NewOfferAddV1FvEmeaUnauthorized() *OfferAddV1FvEmeaUnauthorized {
	return &OfferAddV1FvEmeaUnauthorized{}
}

/*OfferAddV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type OfferAddV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *OfferAddV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *OfferAddV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferAddV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferAddV1FvEmeaForbidden creates a OfferAddV1FvEmeaForbidden with default headers values
func NewOfferAddV1FvEmeaForbidden() *OfferAddV1FvEmeaForbidden {
	return &OfferAddV1FvEmeaForbidden{}
}

/*OfferAddV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type OfferAddV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *OfferAddV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *OfferAddV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferAddV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferAddV1FvEmeaNotFound creates a OfferAddV1FvEmeaNotFound with default headers values
func NewOfferAddV1FvEmeaNotFound() *OfferAddV1FvEmeaNotFound {
	return &OfferAddV1FvEmeaNotFound{}
}

/*OfferAddV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type OfferAddV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *OfferAddV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *OfferAddV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferAddV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferAddV1FvEmeaTooManyRequests creates a OfferAddV1FvEmeaTooManyRequests with default headers values
func NewOfferAddV1FvEmeaTooManyRequests() *OfferAddV1FvEmeaTooManyRequests {
	return &OfferAddV1FvEmeaTooManyRequests{}
}

/*OfferAddV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type OfferAddV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *OfferAddV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *OfferAddV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferAddV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferAddV1FvEmeaStatus452 creates a OfferAddV1FvEmeaStatus452 with default headers values
func NewOfferAddV1FvEmeaStatus452() *OfferAddV1FvEmeaStatus452 {
	return &OfferAddV1FvEmeaStatus452{}
}

/*OfferAddV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type OfferAddV1FvEmeaStatus452 struct {
}

func (o *OfferAddV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaStatus452 ", 452)
}

func (o *OfferAddV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOfferAddV1FvEmeaStatus453 creates a OfferAddV1FvEmeaStatus453 with default headers values
func NewOfferAddV1FvEmeaStatus453() *OfferAddV1FvEmeaStatus453 {
	return &OfferAddV1FvEmeaStatus453{}
}

/*OfferAddV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type OfferAddV1FvEmeaStatus453 struct {
}

func (o *OfferAddV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaStatus453 ", 453)
}

func (o *OfferAddV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOfferAddV1FvEmeaStatus455 creates a OfferAddV1FvEmeaStatus455 with default headers values
func NewOfferAddV1FvEmeaStatus455() *OfferAddV1FvEmeaStatus455 {
	return &OfferAddV1FvEmeaStatus455{}
}

/*OfferAddV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type OfferAddV1FvEmeaStatus455 struct {
}

func (o *OfferAddV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaStatus455 ", 455)
}

func (o *OfferAddV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOfferAddV1FvEmeaStatus465 creates a OfferAddV1FvEmeaStatus465 with default headers values
func NewOfferAddV1FvEmeaStatus465() *OfferAddV1FvEmeaStatus465 {
	return &OfferAddV1FvEmeaStatus465{}
}

/*OfferAddV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPEBSD101E : INVALID ORG ORG  NOT NUMERIC<BR/>VPEBSD102E : INVALID ACCOUNT ACCOUNT  NOT NUMERIC<BR/>VPEBSD103E : OFFER ID MUST BE PRESENT TO ADD OFFER<BR/>VPEBSD104S : ORG IN ADD PENDING STATUS<BR/>VPEBSD105S : ORG RECORD STATUS INCOMPELETE OR TO BE PURGED<BR/>VPEBSD106S : FILE ERROR WHILE READING OFFER ORG RECORD<BR/>VPEBSD107S : GETMAIN FOR OFFER ADD INPUT AREA FAILED<BR/>VPEBSD108S : GETMAIN FOR OFFER ADD OUPUT AREA FAILED<BR/>VPEBSD109S : CANNOT CONTINUE FURTHER PROCESSING AS MAX ERROR COUNT REACHED.<BR/>VPEBSD1Z1S : GET MAIN FAILED FOR THE INPUT COPYBOOK<BR/>VPEBSD1Z2S : SERVICE INPUT LENGTH ERROR<BR/>VPEBSD1Z3S : INVALID USER  SIGNON NOT VALID<BR/>VPEBSD1Z4S : FILE TABLE NOT FOUND ON SMFT FILE<BR/>VPEBSD1Z5S : SYSTEM IN AFTER HOURS STATE PLEASE SUBMIT LATER<BR/>VPEBSD1Z7S : USER DOES NOT HAVE ACCESS TO THIS ORG<BR/>VPEBSD1Z8S : FILE TABLE NOT FOUND ON SMFT FILE<BR/>VPEBSD1Z9S : SYSTEM IN AFTER HOURS STATE PLEASE SUBMIT LATER<BR/>VPLBSAE01E : ORG MUST BE NUMERIC AND VALID VALUES ARE 001 TO 998<BR/>VPLBSAE02E : INPUT ACCOUNT NUMBER IS NOT NUMERIC OR EQUALS SPACES<BR/>VPLBSAE03E : INVALID OFFER ID<BR/>VPLBSAE04E : PRIORITY MUST BE NUMERIC VALID VALUES ARE 1-99999<BR/>VPLBSAE05E : ENROLL ACTION MUST BE NUMERIC VALID VALUES ARE 0-4<BR/>VPLBSAE06E : START DATE INVALID MUST BE NUMERIC AND GREATER THAN ZERO<BR/>VPLBSAE07E : EXPIRE DATE INVALID MUST BE NUMERIC AND GREATER THAN 0<BR/>VPLBSAE11E : EFFECTIVE DATE INVALID MUST BE NUMERIC AND GREATER THAN 0<BR/>VPLBSAE12S : ORG MUST BE NUMERIC AND VALID VALUES ARE 001 TO 998<BR/>VPLBSAE13S : OMS FILE-TABLE USED BY THIS SERVICE IS MISSING<BR/>VPLBSAE14S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUT<BR/>VPLBSAE16S : ORG RECORD NOT FOUND<BR/>VPLBSAE17S : ORG RECORD NOT FOUND<BR/>VPLBSAE18S : ORG RECORD NOT FOUND<BR/>VPLBSAE20S : OFFER ENROLLMENT ALREADY ON FILE<BR/>VPLBSAE21S : OFFER DEFINITION RECORD NOT FOUND<BR/>VPLBSAE22S : OFFER DEFINITION RECORD NOT FOUND<BR/>VPLBSAE23S : OFFER DEFINITION RECORD NOT FOUND<BR/>VPLBSAE24S : INVALID ENROLLMENT ACTION FOR ADD ON PRODUCT OFFER TYPE<BR/>VPLBSAE26S : START DATE IS NOT ALLOWED TO BE ADDED<BR/>VPLBSAE27S : INVALID EXPIRE DATE VALUE MUST BE WITHIN OFFER DATE RANGE<BR/>VPLBSAE28S : EXPIRE DATE MUST BE ZEROS<BR/>VPLBSAE29S : EFF DATE INVALID MUST BE GREATER THAN NEXT PROCESS DATE<BR/>VPLBSAE30S : EFF DATE SHOULD BE ZEROS FOR THIS OFFER TYPE<BR/>VPLBSAE31S : INVALID ORG OR ACCT OR EFF DATE INSURANCE ADD SVC FAIL<BR/>VPLBSAEZ1S : INVALID SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPLBSAEZ2S : SVC INPUT TO INTEREST TABLE ADD SVC IS AN INCORRECT LENGT<BR/>VPLBSAEZ4S : FILE-TABLE USED BY THIS SERVICE IS MISSING<BR/>VPLBSAEZ5S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINU
*/
type OfferAddV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *OfferAddV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerAdd][%d] offerAddV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *OfferAddV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferAddV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
