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

// LoyaltySchemeListV1FvEmeaReader is a Reader for the LoyaltySchemeListV1FvEmea structure.
type LoyaltySchemeListV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoyaltySchemeListV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoyaltySchemeListV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoyaltySchemeListV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLoyaltySchemeListV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoyaltySchemeListV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoyaltySchemeListV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLoyaltySchemeListV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewLoyaltySchemeListV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewLoyaltySchemeListV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewLoyaltySchemeListV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewLoyaltySchemeListV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoyaltySchemeListV1FvEmeaOK creates a LoyaltySchemeListV1FvEmeaOK with default headers values
func NewLoyaltySchemeListV1FvEmeaOK() *LoyaltySchemeListV1FvEmeaOK {
	return &LoyaltySchemeListV1FvEmeaOK{}
}

/*LoyaltySchemeListV1FvEmeaOK handles this case with default header values.

successful operation
*/
type LoyaltySchemeListV1FvEmeaOK struct {
	Payload *models.LoyaltySchemeListResponse
}

func (o *LoyaltySchemeListV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaOK) GetPayload() *models.LoyaltySchemeListResponse {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoyaltySchemeListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltySchemeListV1FvEmeaBadRequest creates a LoyaltySchemeListV1FvEmeaBadRequest with default headers values
func NewLoyaltySchemeListV1FvEmeaBadRequest() *LoyaltySchemeListV1FvEmeaBadRequest {
	return &LoyaltySchemeListV1FvEmeaBadRequest{}
}

/*LoyaltySchemeListV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type LoyaltySchemeListV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltySchemeListV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltySchemeListV1FvEmeaUnauthorized creates a LoyaltySchemeListV1FvEmeaUnauthorized with default headers values
func NewLoyaltySchemeListV1FvEmeaUnauthorized() *LoyaltySchemeListV1FvEmeaUnauthorized {
	return &LoyaltySchemeListV1FvEmeaUnauthorized{}
}

/*LoyaltySchemeListV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type LoyaltySchemeListV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltySchemeListV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltySchemeListV1FvEmeaForbidden creates a LoyaltySchemeListV1FvEmeaForbidden with default headers values
func NewLoyaltySchemeListV1FvEmeaForbidden() *LoyaltySchemeListV1FvEmeaForbidden {
	return &LoyaltySchemeListV1FvEmeaForbidden{}
}

/*LoyaltySchemeListV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type LoyaltySchemeListV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltySchemeListV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltySchemeListV1FvEmeaNotFound creates a LoyaltySchemeListV1FvEmeaNotFound with default headers values
func NewLoyaltySchemeListV1FvEmeaNotFound() *LoyaltySchemeListV1FvEmeaNotFound {
	return &LoyaltySchemeListV1FvEmeaNotFound{}
}

/*LoyaltySchemeListV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type LoyaltySchemeListV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltySchemeListV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltySchemeListV1FvEmeaTooManyRequests creates a LoyaltySchemeListV1FvEmeaTooManyRequests with default headers values
func NewLoyaltySchemeListV1FvEmeaTooManyRequests() *LoyaltySchemeListV1FvEmeaTooManyRequests {
	return &LoyaltySchemeListV1FvEmeaTooManyRequests{}
}

/*LoyaltySchemeListV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type LoyaltySchemeListV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltySchemeListV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoyaltySchemeListV1FvEmeaStatus452 creates a LoyaltySchemeListV1FvEmeaStatus452 with default headers values
func NewLoyaltySchemeListV1FvEmeaStatus452() *LoyaltySchemeListV1FvEmeaStatus452 {
	return &LoyaltySchemeListV1FvEmeaStatus452{}
}

/*LoyaltySchemeListV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type LoyaltySchemeListV1FvEmeaStatus452 struct {
}

func (o *LoyaltySchemeListV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaStatus452 ", 452)
}

func (o *LoyaltySchemeListV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltySchemeListV1FvEmeaStatus453 creates a LoyaltySchemeListV1FvEmeaStatus453 with default headers values
func NewLoyaltySchemeListV1FvEmeaStatus453() *LoyaltySchemeListV1FvEmeaStatus453 {
	return &LoyaltySchemeListV1FvEmeaStatus453{}
}

/*LoyaltySchemeListV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type LoyaltySchemeListV1FvEmeaStatus453 struct {
}

func (o *LoyaltySchemeListV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaStatus453 ", 453)
}

func (o *LoyaltySchemeListV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltySchemeListV1FvEmeaStatus455 creates a LoyaltySchemeListV1FvEmeaStatus455 with default headers values
func NewLoyaltySchemeListV1FvEmeaStatus455() *LoyaltySchemeListV1FvEmeaStatus455 {
	return &LoyaltySchemeListV1FvEmeaStatus455{}
}

/*LoyaltySchemeListV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type LoyaltySchemeListV1FvEmeaStatus455 struct {
}

func (o *LoyaltySchemeListV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaStatus455 ", 455)
}

func (o *LoyaltySchemeListV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoyaltySchemeListV1FvEmeaStatus465 creates a LoyaltySchemeListV1FvEmeaStatus465 with default headers values
func NewLoyaltySchemeListV1FvEmeaStatus465() *LoyaltySchemeListV1FvEmeaStatus465 {
	return &LoyaltySchemeListV1FvEmeaStatus465{}
}

/*LoyaltySchemeListV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPLKSCI01S : SERVICE COULD NOT OBTAIN STORAGE AREA<BR/>VPLKSCI02S : SERVICE INPUT TO SCHEME LIST-INQ SERVICE IS AN INCORRECT LENGTH<BR/>VPLKSCI04S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPLKSCI05S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPLKSCI06S : ORG NUMBER MUST BE NUMERIC AND VALID VALUES ARE 001-997<BR/>VPLKSCI07S : INVALID ACCOUNT NUMBER<BR/>VPLKSCI08S : REQUESTED ORGANIZATION NOT FOUND<BR/>VPLKSCI09S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPLKSCI10S : APPLICATION IN NO-PROCESSING STATUS RETRY IN A FEW MINUTES<BR/>VPLKSCI11S : ORGANIZATION RECORD IS IN PENDING ADD STATUS<BR/>VPLKSCI12S : ORGANIZATION RECORD STATUS IS INCOMPLETEorPURGEDorCLOSED<BR/>VPLKSCI13S : ORGANIZATION RECORD NOT FOUND<BR/>VPLKSCI14S : POINTS ACCOUNT RECORD NOT FOUND<BR/>VPLKSCI15E : POINTS ACCOUNT RECORD IS IN PENDING ADD STATUS<BR/>VPLKSCI16S : ACCOUNT DEMOGRAPHIC RECORD NOT FOUND<BR/>VPLKSCI17E : ACCOUNT DEMOGRAPHIC RECORD STATUS IS INCOMPLETEorCLOSEDorPURGED<BR/>VPLKSCI18S : POINTS SCHEME RECORD NOT FOUND<BR/>VPLKSCI19E : POINTS SCHEME RECORD IS INCOMPLETEorCLOSEDorPURGED<BR/>VPLKSCI20E : POINTS ACCOUNT RECORD IS INCOMPLETEorCLOSEDorPURGEDorPENDING PURGED<BR/>VPLKSCI21E : ACCOUNT DEMOGRAPHIC RECORD IS IN PENDING ADD STATUS<BR/>VPLKSCI22S : LAST SCHEME SHOULD BE PROVIDED FOR SCROLLING<BR/>VPLKSCI23S : FUNCTION CODE OTHER THAN PorNorSPACES ENTERED<BR/>VPLKSCI24E : NUMBER OF ITEMS REQUESTED SHOULD BE NUMERIC AND LESS THAN 2
*/
type LoyaltySchemeListV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *LoyaltySchemeListV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/loyaltySchemeList][%d] loyaltySchemeListV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *LoyaltySchemeListV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *LoyaltySchemeListV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}