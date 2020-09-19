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

// UniqueLocateInquiryV1FvEmeaReader is a Reader for the UniqueLocateInquiryV1FvEmea structure.
type UniqueLocateInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UniqueLocateInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUniqueLocateInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUniqueLocateInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUniqueLocateInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUniqueLocateInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUniqueLocateInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUniqueLocateInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewUniqueLocateInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewUniqueLocateInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewUniqueLocateInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewUniqueLocateInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUniqueLocateInquiryV1FvEmeaOK creates a UniqueLocateInquiryV1FvEmeaOK with default headers values
func NewUniqueLocateInquiryV1FvEmeaOK() *UniqueLocateInquiryV1FvEmeaOK {
	return &UniqueLocateInquiryV1FvEmeaOK{}
}

/*UniqueLocateInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type UniqueLocateInquiryV1FvEmeaOK struct {
	Payload *models.UniqueLocateInquiryResponse
}

func (o *UniqueLocateInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaOK) GetPayload() *models.UniqueLocateInquiryResponse {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UniqueLocateInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaBadRequest creates a UniqueLocateInquiryV1FvEmeaBadRequest with default headers values
func NewUniqueLocateInquiryV1FvEmeaBadRequest() *UniqueLocateInquiryV1FvEmeaBadRequest {
	return &UniqueLocateInquiryV1FvEmeaBadRequest{}
}

/*UniqueLocateInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type UniqueLocateInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *UniqueLocateInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaUnauthorized creates a UniqueLocateInquiryV1FvEmeaUnauthorized with default headers values
func NewUniqueLocateInquiryV1FvEmeaUnauthorized() *UniqueLocateInquiryV1FvEmeaUnauthorized {
	return &UniqueLocateInquiryV1FvEmeaUnauthorized{}
}

/*UniqueLocateInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type UniqueLocateInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *UniqueLocateInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaForbidden creates a UniqueLocateInquiryV1FvEmeaForbidden with default headers values
func NewUniqueLocateInquiryV1FvEmeaForbidden() *UniqueLocateInquiryV1FvEmeaForbidden {
	return &UniqueLocateInquiryV1FvEmeaForbidden{}
}

/*UniqueLocateInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type UniqueLocateInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *UniqueLocateInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaNotFound creates a UniqueLocateInquiryV1FvEmeaNotFound with default headers values
func NewUniqueLocateInquiryV1FvEmeaNotFound() *UniqueLocateInquiryV1FvEmeaNotFound {
	return &UniqueLocateInquiryV1FvEmeaNotFound{}
}

/*UniqueLocateInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type UniqueLocateInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *UniqueLocateInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaTooManyRequests creates a UniqueLocateInquiryV1FvEmeaTooManyRequests with default headers values
func NewUniqueLocateInquiryV1FvEmeaTooManyRequests() *UniqueLocateInquiryV1FvEmeaTooManyRequests {
	return &UniqueLocateInquiryV1FvEmeaTooManyRequests{}
}

/*UniqueLocateInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type UniqueLocateInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *UniqueLocateInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaStatus452 creates a UniqueLocateInquiryV1FvEmeaStatus452 with default headers values
func NewUniqueLocateInquiryV1FvEmeaStatus452() *UniqueLocateInquiryV1FvEmeaStatus452 {
	return &UniqueLocateInquiryV1FvEmeaStatus452{}
}

/*UniqueLocateInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type UniqueLocateInquiryV1FvEmeaStatus452 struct {
}

func (o *UniqueLocateInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaStatus452 ", 452)
}

func (o *UniqueLocateInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaStatus453 creates a UniqueLocateInquiryV1FvEmeaStatus453 with default headers values
func NewUniqueLocateInquiryV1FvEmeaStatus453() *UniqueLocateInquiryV1FvEmeaStatus453 {
	return &UniqueLocateInquiryV1FvEmeaStatus453{}
}

/*UniqueLocateInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type UniqueLocateInquiryV1FvEmeaStatus453 struct {
}

func (o *UniqueLocateInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaStatus453 ", 453)
}

func (o *UniqueLocateInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaStatus455 creates a UniqueLocateInquiryV1FvEmeaStatus455 with default headers values
func NewUniqueLocateInquiryV1FvEmeaStatus455() *UniqueLocateInquiryV1FvEmeaStatus455 {
	return &UniqueLocateInquiryV1FvEmeaStatus455{}
}

/*UniqueLocateInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type UniqueLocateInquiryV1FvEmeaStatus455 struct {
}

func (o *UniqueLocateInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaStatus455 ", 455)
}

func (o *UniqueLocateInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUniqueLocateInquiryV1FvEmeaStatus465 creates a UniqueLocateInquiryV1FvEmeaStatus465 with default headers values
func NewUniqueLocateInquiryV1FvEmeaStatus465() *UniqueLocateInquiryV1FvEmeaStatus465 {
	return &UniqueLocateInquiryV1FvEmeaStatus465{}
}

/*UniqueLocateInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPE5SUV01S : SERVICE REQUEST LENGTH ERROR<BR/>VPE5SUV02S : USER NOT AUTHORISED TO ACCESS THE SERVICE<BR/>VPE5SUV03S : CMS FILE TABLE UNAVAILABLE<BR/>VPE5SUV04S : ORGANIZATION RCEORD NOT FOUND<BR/>VPE5SUV05S : RELATIONSHIP CAN BE ZERO OR SPACE<BR/>VPE5SUV06S : VALID VALUES FOR SEARCH TYPE ARE IorCorAorE<BR/>VPE5SUV07S : VALID VALUES ARE 0 AND 1 FOR SEARH TYPE<BR/>VPE5SUV08S : THIS ORG DOES NOT SUPPORT UNIQUE ID FUNCTIONALITY<BR/>VPE5SUV10S : THIS CLIENT DOES NOT SUPPORT THE CUSTOMER VIEW FUNCTIONALITY<BR/>VPE5SUV12S : INVALID CARD NUMBER<BR/>VPE5SUV13S : CARD NOT ON FILE<BR/>VPE5SUV15S : INVALID ACCOUNT NUMBER<BR/>VPE5SUV16S : ACCOUNT NOT ON FILE<BR/>VPE5SUV18S : CUSTOMER ORG COULD NOT BE DETERMINED FROM CUSTOMER NUMBER<BR/>VPE5SUV19S : CUSTOMER NOT ON FILE<BR/>VPE5SUV20S : INVALID UNIQUE ID<BR/>VPE5SUV21S : UNIQUE ID NOT ON FILE<BR/>VPE5SUV24S : RECORD NOT FOUND IN CUSTOMER-CARD XREF FILE<BR/>VPE5SUV26S : RECORD NOT FOUND IN ACCOUNT-CARD XREF FILE<BR/>VPE5SUV27S : RECORD NOT FOUND IN CUSTOMER-ACCT XREF FILE<BR/>VPE5SUV30S : INVALID FUNCTIONAL CODE.VALID VALUES : 'BorTorForEor '<BR/>VPE5SUV31S : NO CUSTOMER FOUND FOR THE CARD IN THIS VIEW<BR/>VPE5SUV32S : NO CUSTOMER FOUND FOR THE CARD IN THIS VIEW<BR/>VPE5SUV33S : INVALID CUSTOMER NUMBER FOR THE ACCOUNT NUMBER<BR/>VPE5SUV37S : ORG SHOULD BE BETWEEN 0 AND 999<BR/>VPE5SUV38S : SEVERE ERROR END OF FILE REACHED FOR AZUCX<BR/>VPE5SUV39S : INVALID CUSTOMER NUMBER<BR/>VPE5SUV40S : NO DETAIL RECORDS FOUND<BR/>VPE5SUV41S : ORG IS NOT NUMERIC<BR/>VPE5SUV42S : UNIQUE ID NOT ON CUSTOMER RECOR
*/
type UniqueLocateInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *UniqueLocateInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/uniqueLocateInquiry][%d] uniqueLocateInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *UniqueLocateInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *UniqueLocateInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
