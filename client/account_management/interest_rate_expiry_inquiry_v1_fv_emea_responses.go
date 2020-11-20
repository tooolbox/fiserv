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

// InterestRateExpiryInquiryV1FvEmeaReader is a Reader for the InterestRateExpiryInquiryV1FvEmea structure.
type InterestRateExpiryInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterestRateExpiryInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterestRateExpiryInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInterestRateExpiryInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInterestRateExpiryInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInterestRateExpiryInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInterestRateExpiryInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInterestRateExpiryInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewInterestRateExpiryInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewInterestRateExpiryInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewInterestRateExpiryInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewInterestRateExpiryInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterestRateExpiryInquiryV1FvEmeaOK creates a InterestRateExpiryInquiryV1FvEmeaOK with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaOK() *InterestRateExpiryInquiryV1FvEmeaOK {
	return &InterestRateExpiryInquiryV1FvEmeaOK{}
}

/*InterestRateExpiryInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type InterestRateExpiryInquiryV1FvEmeaOK struct {
	Payload *models.InterestRateExpiryInquiryResponse
}

func (o *InterestRateExpiryInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaOK) GetPayload() *models.InterestRateExpiryInquiryResponse {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterestRateExpiryInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaBadRequest creates a InterestRateExpiryInquiryV1FvEmeaBadRequest with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaBadRequest() *InterestRateExpiryInquiryV1FvEmeaBadRequest {
	return &InterestRateExpiryInquiryV1FvEmeaBadRequest{}
}

/*InterestRateExpiryInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type InterestRateExpiryInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *InterestRateExpiryInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaUnauthorized creates a InterestRateExpiryInquiryV1FvEmeaUnauthorized with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaUnauthorized() *InterestRateExpiryInquiryV1FvEmeaUnauthorized {
	return &InterestRateExpiryInquiryV1FvEmeaUnauthorized{}
}

/*InterestRateExpiryInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type InterestRateExpiryInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *InterestRateExpiryInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaForbidden creates a InterestRateExpiryInquiryV1FvEmeaForbidden with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaForbidden() *InterestRateExpiryInquiryV1FvEmeaForbidden {
	return &InterestRateExpiryInquiryV1FvEmeaForbidden{}
}

/*InterestRateExpiryInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type InterestRateExpiryInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *InterestRateExpiryInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaNotFound creates a InterestRateExpiryInquiryV1FvEmeaNotFound with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaNotFound() *InterestRateExpiryInquiryV1FvEmeaNotFound {
	return &InterestRateExpiryInquiryV1FvEmeaNotFound{}
}

/*InterestRateExpiryInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type InterestRateExpiryInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *InterestRateExpiryInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaTooManyRequests creates a InterestRateExpiryInquiryV1FvEmeaTooManyRequests with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaTooManyRequests() *InterestRateExpiryInquiryV1FvEmeaTooManyRequests {
	return &InterestRateExpiryInquiryV1FvEmeaTooManyRequests{}
}

/*InterestRateExpiryInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type InterestRateExpiryInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *InterestRateExpiryInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaStatus452 creates a InterestRateExpiryInquiryV1FvEmeaStatus452 with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaStatus452() *InterestRateExpiryInquiryV1FvEmeaStatus452 {
	return &InterestRateExpiryInquiryV1FvEmeaStatus452{}
}

/*InterestRateExpiryInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type InterestRateExpiryInquiryV1FvEmeaStatus452 struct {
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaStatus452 ", 452)
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaStatus453 creates a InterestRateExpiryInquiryV1FvEmeaStatus453 with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaStatus453() *InterestRateExpiryInquiryV1FvEmeaStatus453 {
	return &InterestRateExpiryInquiryV1FvEmeaStatus453{}
}

/*InterestRateExpiryInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type InterestRateExpiryInquiryV1FvEmeaStatus453 struct {
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaStatus453 ", 453)
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaStatus455 creates a InterestRateExpiryInquiryV1FvEmeaStatus455 with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaStatus455() *InterestRateExpiryInquiryV1FvEmeaStatus455 {
	return &InterestRateExpiryInquiryV1FvEmeaStatus455{}
}

/*InterestRateExpiryInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type InterestRateExpiryInquiryV1FvEmeaStatus455 struct {
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaStatus455 ", 455)
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInterestRateExpiryInquiryV1FvEmeaStatus465 creates a InterestRateExpiryInquiryV1FvEmeaStatus465 with default headers values
func NewInterestRateExpiryInquiryV1FvEmeaStatus465() *InterestRateExpiryInquiryV1FvEmeaStatus465 {
	return &InterestRateExpiryInquiryV1FvEmeaStatus465{}
}

/*InterestRateExpiryInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPE1SRE01S : SERVICE INPUT TO RorE SERVICE IS AN INCORRECT LENGTH<BR/>VPE5SRE26S : LOGO FILE NOT OPEN<BR/>VPE5SRE27S : LOGO NUMBER NOT FOUND OR ADD PENDING<BR/>VPE1SRE00S : SERVICE REQUEST LENGTH ERROR<BR/>VPE5SRE05S : CMS FILE-TABLE RECORD NOT FOUND<BR/>VPE5SRE03S : CMS IS NO-PROCESSING STATE TRY LATER<BR/>VPE5SRE06S : ORG NBR MUST BE NUMERIC AND VALID VALUES ARE 001-998<BR/>VPE5SRE07S : ACCOUNT IS NOT NUMERIC OR EQUALS SPACES<BR/>VPE5SRE56S : INVALID LOCALorFOREIGN INDICATOR SPECIFIED<BR/>VPE5SRE08S : ORGANIZATION NUMBER NOT DETERMINED<BR/>VPE5SRE10S : REQUESTED ORGANIZATION NUMBER IS NOT FOUND<BR/>VPE5SRE11S : CMS FILE-TABLE RECORD NOT FOUND<BR/>VPE5SRE19S : CMS IS NO-PROCESSING STATE TRY LATER<BR/>VPE5SRE13S : ORGANIZATION FILE NOT OPEN<BR/>VPE5SRE14S : ORGANIZATION RECORD NOT FOUND OR ADD PENDING<BR/>VPE5SRE13I : DUAL CURRENCY ORG IS ZERO<BR/>VPE5SRE15S : ACCOUNT FILE NOT OPEN<BR/>VPE5SRE02S : BASE SEGMENT FILE READ ERROR<BR/>VPE5SRE22S : PLAN HAS EXPIRED IN PLAN MASTER FILE<BR/>VPE5SRE23S : PLAN BEGINING AND END DATES ARE INVALID<BR/>VPE5SRE24E : RATE TYPE NOT F Z OR V<BR/>VPE5SRE28S : PLAN BEGINING AND END DATES ARE INVALID<BR/>VPE5SRE29E : RATE TYPE NOT F Z OR V<BR/>VPE5SRE30E : RATE TYPE IS NOT F<BR/>VPE5SRE31S : PLAN MASTER FILE NOT OPEN<BR/>VPE5SRE32S : PLAN MASTER NOT FOUND<BR/>VPE5SRE52S : PLAN SEGMENT FILE READ ERROR<BR/>VPE5SRE33S : RATE TABLE FILE NOT OPEN<BR/>VPE5SRE34S : RATE TABLE NOT FOUND<BR/>VPE5SRE35S : RATE TABLE FILE NOT OPEN<BR/>VPE5SRE36S : INTREST TABLE NOT FOUN
*/
type InterestRateExpiryInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/interestRateExpiryInquiry][%d] interestRateExpiryInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *InterestRateExpiryInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}