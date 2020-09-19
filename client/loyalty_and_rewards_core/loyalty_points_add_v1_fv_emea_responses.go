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

// LoyaltyPointsAddV1FvEmeaReader is a Reader for the LoyaltyPointsAddV1FvEmea structure.
type LoyaltyPointsAddV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoyaltyPointsAddV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoyaltyPointsAddV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoyaltyPointsAddV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLoyaltyPointsAddV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoyaltyPointsAddV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoyaltyPointsAddV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoyaltyPointsAddV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLoyaltyPointsAddV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLoyaltyPointsAddV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLoyaltyPointsAddV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewLoyaltyPointsAddV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoyaltyPointsAddV1FvEmeaOK creates a LoyaltyPointsAddV1FvEmeaOK with default headers values
func NewLoyaltyPointsAddV1FvEmeaOK() *LoyaltyPointsAddV1FvEmeaOK {
	return &LoyaltyPointsAddV1FvEmeaOK{}
}

/*LoyaltyPointsAddV1FvEmeaOK handles this case with default header values.

successful operation
*/
type LoyaltyPointsAddV1FvEmeaOK struct {
	Payload *models.LoyaltyPointsAddResponse
}

func (o *LoyaltyPointsAddV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaOK) GetPayload() *models.LoyaltyPointsAddResponse {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoyaltyPointsAddResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaBadRequest creates a LoyaltyPointsAddV1FvEmeaBadRequest with default headers values
func NewLoyaltyPointsAddV1FvEmeaBadRequest() *LoyaltyPointsAddV1FvEmeaBadRequest {
	return &LoyaltyPointsAddV1FvEmeaBadRequest{}
}

/*LoyaltyPointsAddV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LoyaltyPointsAddV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsAddV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaUnauthorized creates a LoyaltyPointsAddV1FvEmeaUnauthorized with default headers values
func NewLoyaltyPointsAddV1FvEmeaUnauthorized() *LoyaltyPointsAddV1FvEmeaUnauthorized {
	return &LoyaltyPointsAddV1FvEmeaUnauthorized{}
}

/*LoyaltyPointsAddV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LoyaltyPointsAddV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsAddV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaForbidden creates a LoyaltyPointsAddV1FvEmeaForbidden with default headers values
func NewLoyaltyPointsAddV1FvEmeaForbidden() *LoyaltyPointsAddV1FvEmeaForbidden {
	return &LoyaltyPointsAddV1FvEmeaForbidden{}
}

/*LoyaltyPointsAddV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LoyaltyPointsAddV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsAddV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaNotFound creates a LoyaltyPointsAddV1FvEmeaNotFound with default headers values
func NewLoyaltyPointsAddV1FvEmeaNotFound() *LoyaltyPointsAddV1FvEmeaNotFound {
	return &LoyaltyPointsAddV1FvEmeaNotFound{}
}

/*LoyaltyPointsAddV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LoyaltyPointsAddV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsAddV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaTooManyRequests creates a LoyaltyPointsAddV1FvEmeaTooManyRequests with default headers values
func NewLoyaltyPointsAddV1FvEmeaTooManyRequests() *LoyaltyPointsAddV1FvEmeaTooManyRequests {
	return &LoyaltyPointsAddV1FvEmeaTooManyRequests{}
}

/*LoyaltyPointsAddV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LoyaltyPointsAddV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsAddV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaStatus452 creates a LoyaltyPointsAddV1FvEmeaStatus452 with default headers values
func NewLoyaltyPointsAddV1FvEmeaStatus452() *LoyaltyPointsAddV1FvEmeaStatus452 {
	return &LoyaltyPointsAddV1FvEmeaStatus452{}
}

/*LoyaltyPointsAddV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LoyaltyPointsAddV1FvEmeaStatus452 struct {
}

func (o *LoyaltyPointsAddV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaStatus452 ", 452)
}

func (o *LoyaltyPointsAddV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaStatus453 creates a LoyaltyPointsAddV1FvEmeaStatus453 with default headers values
func NewLoyaltyPointsAddV1FvEmeaStatus453() *LoyaltyPointsAddV1FvEmeaStatus453 {
	return &LoyaltyPointsAddV1FvEmeaStatus453{}
}

/*LoyaltyPointsAddV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LoyaltyPointsAddV1FvEmeaStatus453 struct {
}

func (o *LoyaltyPointsAddV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaStatus453 ", 453)
}

func (o *LoyaltyPointsAddV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaStatus455 creates a LoyaltyPointsAddV1FvEmeaStatus455 with default headers values
func NewLoyaltyPointsAddV1FvEmeaStatus455() *LoyaltyPointsAddV1FvEmeaStatus455 {
	return &LoyaltyPointsAddV1FvEmeaStatus455{}
}

/*LoyaltyPointsAddV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LoyaltyPointsAddV1FvEmeaStatus455 struct {
}

func (o *LoyaltyPointsAddV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaStatus455 ", 455)
}

func (o *LoyaltyPointsAddV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltyPointsAddV1FvEmeaStatus465 creates a LoyaltyPointsAddV1FvEmeaStatus465 with default headers values
func NewLoyaltyPointsAddV1FvEmeaStatus465() *LoyaltyPointsAddV1FvEmeaStatus465 {
	return &LoyaltyPointsAddV1FvEmeaStatus465{}
}

/*LoyaltyPointsAddV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPLKSAA01S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPLKSAA02S : SERVICE INPUT TO POINT ACCOUNT ADD SERVICE IS AN INCORRECT LENGTH<BR/>VPLKSAA04S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPLKSAA05S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPLKSAA06S : ORG NUMBER MUST BE NUMERIC AND VALID VALUES ARE 001-997<BR/>VPLKSAA07S : INVALID ACCOUNT NUMBER<BR/>VPLKSAA08S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPLKSAA09S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPLKSAA10S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPLKSAA11S : ORGANIZATION RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAA12S : ORGANIZATION RECORD STATUS IS INCOMP OR PURGED OR CLOSED<BR/>VPLKSAA13S : ORGANIZATION RECORD NOT FOUND<BR/>VPLKSAA14S : ENTRY REQUIRED FOR SCHEME<BR/>VPLKSAA15S : POINT SCHEME RECORD NOT FOUND<BR/>VPLKSAA16S : POINT SCHEME RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAA17S : POINT SCHEME RECORD STATUS IS IN PURGED STATUS<BR/>VPLKSAA18S : INVALID SCHEME DATE RANGE<BR/>VPLKSAA19S : INVALID POINTS SCHEME RECORD<BR/>VPLKSAA20S : POINTS ACCOUNT RECORD IS ALREADY IN COMPLETE STATUS<BR/>VPLKSAA21E : INVALID SERVICE FLAG VALUE VALID VALUES ARE 'B' 'C' 'I'<BR/>VPLKSAA22S : ACCOUNT DEMOGRAPHICS RECORD NOT FOUND<BR/>VPLKSAA23S : DEMOGRAPHICS RECORD IS IN PENDING ADD STATUS<BR/>VPLKSAA24S : DEMOGRAPHICS RECORD STATUS IS IN PURGED STATUS<BR/>VPLKSAA25S : INVALID DEMOGRAPHICS DATE RANGE<BR/>VPLKSAA26S : DEMOGRAPHICS RECORD IS NOT IN ACTIVE STATUS<BR/>VPLKSAA27S : PRIMARY ACCOUNT SOURCE IS NOT 'CMS'<BR/>VPLKSAA28E : INVALID ACTION REQUESTED FOR POINTS RECORD<BR/>VPLKSAA29E : INVALID STATEMENT MONTH FREQ VALUE VALID VALUES ARE 01 THRU 99<BR/>VPLKSAA30E : INVALID STMT DAY OF MONTH VALID VALUES ARE 01 THRU 31 AND 99<BR/>VPLKSAA31E : NEXT STATEMENT DATE IS INVALID<BR/>VPLKSAA32E : NEXT STATEMENT DATE IS NOT A VALID FUTURE DAT
*/
type LoyaltyPointsAddV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltyPointsAddV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltyPointsAdd][%d] loyaltyPointsAddV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *LoyaltyPointsAddV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltyPointsAddV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
