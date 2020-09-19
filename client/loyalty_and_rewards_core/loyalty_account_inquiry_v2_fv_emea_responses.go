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

// LoyaltyAccountInquiryV2FvEmeaReader is a Reader for the LoyaltyAccountInquiryV2FvEmea structure.
type LoyaltyAccountInquiryV2FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoyaltyAccountInquiryV2FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoyaltyAccountInquiryV2FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoyaltyAccountInquiryV2FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLoyaltyAccountInquiryV2FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoyaltyAccountInquiryV2FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoyaltyAccountInquiryV2FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoyaltyAccountInquiryV2FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLoyaltyAccountInquiryV2FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLoyaltyAccountInquiryV2FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLoyaltyAccountInquiryV2FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewLoyaltyAccountInquiryV2FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoyaltyAccountInquiryV2FvEmeaOK creates a LoyaltyAccountInquiryV2FvEmeaOK with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaOK() *LoyaltyAccountInquiryV2FvEmeaOK {
	return &LoyaltyAccountInquiryV2FvEmeaOK{}
}

/*LoyaltyAccountInquiryV2FvEmeaOK handles this case with default header values.

successful operation
*/
type LoyaltyAccountInquiryV2FvEmeaOK struct {
	Payload *models.LoyaltyAccountInquiryResponse2
}

func (o *LoyaltyAccountInquiryV2FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaOK) GetPayload() *models.LoyaltyAccountInquiryResponse2 {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoyaltyAccountInquiryResponse2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaBadRequest creates a LoyaltyAccountInquiryV2FvEmeaBadRequest with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaBadRequest() *LoyaltyAccountInquiryV2FvEmeaBadRequest {
	return &LoyaltyAccountInquiryV2FvEmeaBadRequest{}
}

/*LoyaltyAccountInquiryV2FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LoyaltyAccountInquiryV2FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyAccountInquiryV2FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaUnauthorized creates a LoyaltyAccountInquiryV2FvEmeaUnauthorized with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaUnauthorized() *LoyaltyAccountInquiryV2FvEmeaUnauthorized {
	return &LoyaltyAccountInquiryV2FvEmeaUnauthorized{}
}

/*LoyaltyAccountInquiryV2FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LoyaltyAccountInquiryV2FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyAccountInquiryV2FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaForbidden creates a LoyaltyAccountInquiryV2FvEmeaForbidden with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaForbidden() *LoyaltyAccountInquiryV2FvEmeaForbidden {
	return &LoyaltyAccountInquiryV2FvEmeaForbidden{}
}

/*LoyaltyAccountInquiryV2FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LoyaltyAccountInquiryV2FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyAccountInquiryV2FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaNotFound creates a LoyaltyAccountInquiryV2FvEmeaNotFound with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaNotFound() *LoyaltyAccountInquiryV2FvEmeaNotFound {
	return &LoyaltyAccountInquiryV2FvEmeaNotFound{}
}

/*LoyaltyAccountInquiryV2FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LoyaltyAccountInquiryV2FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyAccountInquiryV2FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaTooManyRequests creates a LoyaltyAccountInquiryV2FvEmeaTooManyRequests with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaTooManyRequests() *LoyaltyAccountInquiryV2FvEmeaTooManyRequests {
	return &LoyaltyAccountInquiryV2FvEmeaTooManyRequests{}
}

/*LoyaltyAccountInquiryV2FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LoyaltyAccountInquiryV2FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyAccountInquiryV2FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaStatus452 creates a LoyaltyAccountInquiryV2FvEmeaStatus452 with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaStatus452() *LoyaltyAccountInquiryV2FvEmeaStatus452 {
	return &LoyaltyAccountInquiryV2FvEmeaStatus452{}
}

/*LoyaltyAccountInquiryV2FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LoyaltyAccountInquiryV2FvEmeaStatus452 struct {
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaStatus452 ", 452)
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaStatus453 creates a LoyaltyAccountInquiryV2FvEmeaStatus453 with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaStatus453() *LoyaltyAccountInquiryV2FvEmeaStatus453 {
	return &LoyaltyAccountInquiryV2FvEmeaStatus453{}
}

/*LoyaltyAccountInquiryV2FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LoyaltyAccountInquiryV2FvEmeaStatus453 struct {
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaStatus453 ", 453)
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaStatus455 creates a LoyaltyAccountInquiryV2FvEmeaStatus455 with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaStatus455() *LoyaltyAccountInquiryV2FvEmeaStatus455 {
	return &LoyaltyAccountInquiryV2FvEmeaStatus455{}
}

/*LoyaltyAccountInquiryV2FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LoyaltyAccountInquiryV2FvEmeaStatus455 struct {
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaStatus455 ", 455)
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyAccountInquiryV2FvEmeaStatus465 creates a LoyaltyAccountInquiryV2FvEmeaStatus465 with default headers values
func NewLoyaltyAccountInquiryV2FvEmeaStatus465() *LoyaltyAccountInquiryV2FvEmeaStatus465 {
	return &LoyaltyAccountInquiryV2FvEmeaStatus465{}
}

/*LoyaltyAccountInquiryV2FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPLKSAI01S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPLKSAI02S : SERVICE INPUT TO ACCOUNT INQUIRY SERVICE IS AN INCORRECT LENGTH<BR/>VPLKSAI04S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPLKSAI05S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPLKSAI06S : ORG NUMBER MUST BE NUMERIC AND VALID VALUES ARE 001-997<BR/>VPLKSAI07S : INVALID ACCOUNT NUMBER<BR/>VPLKSAI08S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPLKSAI09S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPLKSAI10S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPLKSAI11S : ORGANIZATION RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAI12S : ORGANIZATION RECORD STATUS IS INCOMPLETE OR PURGED OR CLOSED<BR/>VPLKSAI13S : ORGANIZATION RECORD NOT FOUND<BR/>VPLKSAI14S : POINTS ACCOUNT RECORD NOT FOUND<BR/>VPLKSAI15E : POINT ACCOUNT RECORD IS INCOMPLETEorCLOSEDorPENDING PURGEDorPURGED<BR/>VPLKSAI16S : ACCOUNT DEMOGRAPHIC RECORD NOT FOUND<BR/>VPLKSAI17E : ACCOUNT DEMOGRAPHIC RECORD STATUS IS INCOMPLETE<BR/>VPLKSAI18S : POINTS SCHEME RECORD NOT FOUND<BR/>VPLKSAI19E : POINTS SCHEME RECORD IS INCOMPLETEorCLOSEDorPURGED<BR/>VPLKSAI20E : POINTS SCHEME RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAI21E : ACCOUNT DEMOGRAPHIC RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAI22E : POINTS ACCOUNT RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAI23E : INVALID SCHEME DATE RANG
*/
type LoyaltyAccountInquiryV2FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/loyaltyAccountInquiry][%d] loyaltyAccountInquiryV2FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyAccountInquiryV2FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
