// Code generated by go-swagger; DO NOT EDIT.

package account_maintenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// AccountUserFieldsUpdateV2FvEmeaReader is a Reader for the AccountUserFieldsUpdateV2FvEmea structure.
type AccountUserFieldsUpdateV2FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountUserFieldsUpdateV2FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountUserFieldsUpdateV2FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountUserFieldsUpdateV2FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountUserFieldsUpdateV2FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccountUserFieldsUpdateV2FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountUserFieldsUpdateV2FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAccountUserFieldsUpdateV2FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewAccountUserFieldsUpdateV2FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewAccountUserFieldsUpdateV2FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewAccountUserFieldsUpdateV2FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewAccountUserFieldsUpdateV2FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountUserFieldsUpdateV2FvEmeaOK creates a AccountUserFieldsUpdateV2FvEmeaOK with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaOK() *AccountUserFieldsUpdateV2FvEmeaOK {
	return &AccountUserFieldsUpdateV2FvEmeaOK{}
}

/*AccountUserFieldsUpdateV2FvEmeaOK handles this case with default header values.

successful operation
*/
type AccountUserFieldsUpdateV2FvEmeaOK struct {
	Payload *models.AccountUserFieldsUpdateResponse2
}

func (o *AccountUserFieldsUpdateV2FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaOK  %+v", 200, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaOK) GetPayload() *models.AccountUserFieldsUpdateResponse2 {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountUserFieldsUpdateResponse2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaBadRequest creates a AccountUserFieldsUpdateV2FvEmeaBadRequest with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaBadRequest() *AccountUserFieldsUpdateV2FvEmeaBadRequest {
	return &AccountUserFieldsUpdateV2FvEmeaBadRequest{}
}

/*AccountUserFieldsUpdateV2FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type AccountUserFieldsUpdateV2FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV2FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaUnauthorized creates a AccountUserFieldsUpdateV2FvEmeaUnauthorized with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaUnauthorized() *AccountUserFieldsUpdateV2FvEmeaUnauthorized {
	return &AccountUserFieldsUpdateV2FvEmeaUnauthorized{}
}

/*AccountUserFieldsUpdateV2FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type AccountUserFieldsUpdateV2FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV2FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaForbidden creates a AccountUserFieldsUpdateV2FvEmeaForbidden with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaForbidden() *AccountUserFieldsUpdateV2FvEmeaForbidden {
	return &AccountUserFieldsUpdateV2FvEmeaForbidden{}
}

/*AccountUserFieldsUpdateV2FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type AccountUserFieldsUpdateV2FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV2FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaNotFound creates a AccountUserFieldsUpdateV2FvEmeaNotFound with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaNotFound() *AccountUserFieldsUpdateV2FvEmeaNotFound {
	return &AccountUserFieldsUpdateV2FvEmeaNotFound{}
}

/*AccountUserFieldsUpdateV2FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type AccountUserFieldsUpdateV2FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV2FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaTooManyRequests creates a AccountUserFieldsUpdateV2FvEmeaTooManyRequests with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaTooManyRequests() *AccountUserFieldsUpdateV2FvEmeaTooManyRequests {
	return &AccountUserFieldsUpdateV2FvEmeaTooManyRequests{}
}

/*AccountUserFieldsUpdateV2FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type AccountUserFieldsUpdateV2FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV2FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaStatus452 creates a AccountUserFieldsUpdateV2FvEmeaStatus452 with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaStatus452() *AccountUserFieldsUpdateV2FvEmeaStatus452 {
	return &AccountUserFieldsUpdateV2FvEmeaStatus452{}
}

/*AccountUserFieldsUpdateV2FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type AccountUserFieldsUpdateV2FvEmeaStatus452 struct {
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaStatus452 ", 452)
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaStatus453 creates a AccountUserFieldsUpdateV2FvEmeaStatus453 with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaStatus453() *AccountUserFieldsUpdateV2FvEmeaStatus453 {
	return &AccountUserFieldsUpdateV2FvEmeaStatus453{}
}

/*AccountUserFieldsUpdateV2FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type AccountUserFieldsUpdateV2FvEmeaStatus453 struct {
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaStatus453 ", 453)
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaStatus455 creates a AccountUserFieldsUpdateV2FvEmeaStatus455 with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaStatus455() *AccountUserFieldsUpdateV2FvEmeaStatus455 {
	return &AccountUserFieldsUpdateV2FvEmeaStatus455{}
}

/*AccountUserFieldsUpdateV2FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type AccountUserFieldsUpdateV2FvEmeaStatus455 struct {
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaStatus455 ", 455)
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserFieldsUpdateV2FvEmeaStatus465 creates a AccountUserFieldsUpdateV2FvEmeaStatus465 with default headers values
func NewAccountUserFieldsUpdateV2FvEmeaStatus465() *AccountUserFieldsUpdateV2FvEmeaStatus465 {
	return &AccountUserFieldsUpdateV2FvEmeaStatus465{}
}

/*AccountUserFieldsUpdateV2FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL5SUD02S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SUD03E : ORGANIZATION IS NOT NUMERIC OR MUST BE VALUE BETWEEN 000 AND 998<BR/>VPL5SUD04E : REQUESTED CARDorACCT NUMBER IS NOT NUMERIC OR EQUAL SPACES<BR/>VPL5SUD05E : INVALID FOREIGN USE IND SPECIFIED - VALID VALUES SPACE 0 1 OR 2<BR/>VPL5SUD06S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPL5SUD07S : ORGANIZATION NOT FOUND ON FILE<BR/>VPL5SUD08S : FOREIGN ORGANIZATION NOT FOUND ON FILE<BR/>VPL5SUD09S : ACCOUNT NUMBER NOT FOUND<BR/>VPL5SUD10S : ACCOUNT NUMBER NOT FOUND ON FILE<BR/>VPL5SUD11S : LOGO NOT FOUND ON FILE<BR/>VPL5SUD13I : FOREIGN USE ORG UNAVAILABLE PROCESSED LOCAL ORG ACCOUNT DATA<BR/>VPL5SUD18S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPL5SUD19S : SERVICE INPUT TO USER DATES AMTS AND CODES IS INCORRECT LENGTH<BR/>VPL5SUD20S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SUD21S : REQUESTED ORGANIZATION COULD NOT BE DETERMINED<BR/>VPL5SUD22S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPL5SUD23S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPL5SUD50E : USER DATE 1 IS NOT NUMERIC<BR/>VPL5SUD51E : USER DATE 1 IS INVALID<BR/>VPL5SUD52E : USER DATE 2 IS NOT NUMERIC<BR/>VPL5SUD53E : USER DATE 2 IS INVALID<BR/>VPL5SUD54E : USER DATE 3 IS NOT NUMERIC<BR/>VPL5SUD55E : USER DATE 3 IS INVALID<BR/>VPL5SUD56E : USER DATE 4 IS NOT NUMERIC<BR/>VPL5SUD57E : USER DATE 4 IS INVALID<BR/>VPL5SUD58E : USER DATE 5 IS NOT NUMERIC<BR/>VPL5SUD59E : USER DATE 5 IS INVALID<BR/>VPL5SUD60E : USER DATE 6 IS NOT NUMERIC<BR/>VPL5SUD61E : USER DATE 6 IS INVALID<BR/>VPL5SUD62E : USER DATE 7 IS NOT NUMERIC<BR/>VPL5SUD63E : USER DATE 7 IS INVALID<BR/>VPL5SUD64E : USER DATE 8 IS NOT NUMERIC<BR/>VPL5SUD65E : USER DATE 8 IS INVALID<BR/>VPL5SUD66E : USER DATE 9 IS NOT NUMERIC<BR/>VPL5SUD67E : USER DATE 9 IS INVALID<BR/>VPL5SUD68E : USER DATE 10 IS NOT NUMERIC<BR/>VPL5SUD69E : USER DATE 10 IS INVALID<BR/>VPL5SUD70E : USER DATE 11 IS NOT NUMERIC<BR/>VPL5SUD71E : USER DATE 11 IS INVALID<BR/>VPL5SUD72E : USER DATE 12 IS NOT NUMERIC<BR/>VPL5SUD73E : USER DATE 12 IS INVALID<BR/>VPL5SUD74E : USER DATE 13 IS NOT NUMERIC<BR/>VPL5SUD75E : USER DATE 13 IS INVALID<BR/>VPL5SUD76E : USER DATE 14 IS NOT NUMERIC<BR/>VPL5SUD77E : USER DATE 14 IS INVALID<BR/>VPL5SUD81E : USER AMOUNT 1 IS NOT NUMERIC<BR/>VPL5SUD82E : USER AMOUNT 2 IS NOT NUMERIC<BR/>VPL5SUD83E : USER AMOUNT 3 IS NOT NUMERIC<BR/>VPL5SUD84E : USER AMOUNT 4 IS NOT NUMERIC<BR/>VPL5SUD85E : USER AMOUNT 5 IS NOT NUMERIC<BR/>VPL5SUD86E : USER AMOUNT 6 IS NOT NUMERIC<BR/>VPL5SUD87E : USER AMOUNT 7 IS NOT NUMERIC<BR/>VPL5SUD88E : USER AMOUNT 8 IS NOT NUMERIC<BR/>VPL5SUD89E : USER AMOUNT 9 IS NOT NUMERIC<BR/>VPL5SUD90E : USER AMOUNT 10 IS NOT NUMERIC<BR/>VPL5SUD91E : USER AMOUNT 11 IS NOT NUMERIC<BR/>VPL5SUD92E : USER AMOUNT 12 IS NOT NUMERIC<BR/>VPL5SUD93E : USER AMOUNT 13 IS NOT NUMERIC<BR/>VPL5SUD94E : USER AMOUNT 14 IS NOT NUMERIC<BR/>VPL5SUD95E : FI SUB STATUS UDPATE FLAG IS NOT NUMERIC OR NOT EQUAL TO 0 AND
*/
type AccountUserFieldsUpdateV2FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v2/accountUserFieldsUpdate][%d] accountUserFieldsUpdateV2FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *AccountUserFieldsUpdateV2FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
