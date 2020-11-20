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

// OfferEnrolmentInquiryV1FvEmeaReader is a Reader for the OfferEnrolmentInquiryV1FvEmea structure.
type OfferEnrolmentInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OfferEnrolmentInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOfferEnrolmentInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOfferEnrolmentInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOfferEnrolmentInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOfferEnrolmentInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOfferEnrolmentInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOfferEnrolmentInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewOfferEnrolmentInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewOfferEnrolmentInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewOfferEnrolmentInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewOfferEnrolmentInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOfferEnrolmentInquiryV1FvEmeaOK creates a OfferEnrolmentInquiryV1FvEmeaOK with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaOK() *OfferEnrolmentInquiryV1FvEmeaOK {
	return &OfferEnrolmentInquiryV1FvEmeaOK{}
}

/*OfferEnrolmentInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type OfferEnrolmentInquiryV1FvEmeaOK struct {
	Payload *models.OfferEnrolmentInquiryResponse
}

func (o *OfferEnrolmentInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaOK) GetPayload() *models.OfferEnrolmentInquiryResponse {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OfferEnrolmentInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaBadRequest creates a OfferEnrolmentInquiryV1FvEmeaBadRequest with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaBadRequest() *OfferEnrolmentInquiryV1FvEmeaBadRequest {
	return &OfferEnrolmentInquiryV1FvEmeaBadRequest{}
}

/*OfferEnrolmentInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type OfferEnrolmentInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *OfferEnrolmentInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaUnauthorized creates a OfferEnrolmentInquiryV1FvEmeaUnauthorized with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaUnauthorized() *OfferEnrolmentInquiryV1FvEmeaUnauthorized {
	return &OfferEnrolmentInquiryV1FvEmeaUnauthorized{}
}

/*OfferEnrolmentInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type OfferEnrolmentInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *OfferEnrolmentInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaForbidden creates a OfferEnrolmentInquiryV1FvEmeaForbidden with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaForbidden() *OfferEnrolmentInquiryV1FvEmeaForbidden {
	return &OfferEnrolmentInquiryV1FvEmeaForbidden{}
}

/*OfferEnrolmentInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type OfferEnrolmentInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *OfferEnrolmentInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaNotFound creates a OfferEnrolmentInquiryV1FvEmeaNotFound with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaNotFound() *OfferEnrolmentInquiryV1FvEmeaNotFound {
	return &OfferEnrolmentInquiryV1FvEmeaNotFound{}
}

/*OfferEnrolmentInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type OfferEnrolmentInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *OfferEnrolmentInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaTooManyRequests creates a OfferEnrolmentInquiryV1FvEmeaTooManyRequests with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaTooManyRequests() *OfferEnrolmentInquiryV1FvEmeaTooManyRequests {
	return &OfferEnrolmentInquiryV1FvEmeaTooManyRequests{}
}

/*OfferEnrolmentInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type OfferEnrolmentInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *OfferEnrolmentInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaStatus452 creates a OfferEnrolmentInquiryV1FvEmeaStatus452 with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaStatus452() *OfferEnrolmentInquiryV1FvEmeaStatus452 {
	return &OfferEnrolmentInquiryV1FvEmeaStatus452{}
}

/*OfferEnrolmentInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type OfferEnrolmentInquiryV1FvEmeaStatus452 struct {
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaStatus452 ", 452)
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaStatus453 creates a OfferEnrolmentInquiryV1FvEmeaStatus453 with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaStatus453() *OfferEnrolmentInquiryV1FvEmeaStatus453 {
	return &OfferEnrolmentInquiryV1FvEmeaStatus453{}
}

/*OfferEnrolmentInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type OfferEnrolmentInquiryV1FvEmeaStatus453 struct {
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaStatus453 ", 453)
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaStatus455 creates a OfferEnrolmentInquiryV1FvEmeaStatus455 with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaStatus455() *OfferEnrolmentInquiryV1FvEmeaStatus455 {
	return &OfferEnrolmentInquiryV1FvEmeaStatus455{}
}

/*OfferEnrolmentInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type OfferEnrolmentInquiryV1FvEmeaStatus455 struct {
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaStatus455 ", 455)
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOfferEnrolmentInquiryV1FvEmeaStatus465 creates a OfferEnrolmentInquiryV1FvEmeaStatus465 with default headers values
func NewOfferEnrolmentInquiryV1FvEmeaStatus465() *OfferEnrolmentInquiryV1FvEmeaStatus465 {
	return &OfferEnrolmentInquiryV1FvEmeaStatus465{}
}

/*OfferEnrolmentInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPEBSQEZ1S : SERVICE COULD NOT OBTAIN STORAGE OUTPUT AREA<BR/>VPEBSQEZ2S : SERVICE INPUT TO QE1 SERVICE HAS AN INCORRECT LENGHTH<BR/>VPEBSQEZ3S : FILE-TABLE USED BY THIS SERVICE IS MISSING<BR/>VPEBSQEZ5S : REQUESTED ACCOUNT NUMBER IS NOT FOUND<BR/>VPEBSQEZ6S : ORGANIZATION MUST BE NUMERIC AND VALID VALUES ARE 000-998TES<BR/>VPEBSQE02S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPEBSQE03E : ORG NUMBER MUST BE NUMERIC VALID VALUES ARE 001-998<BR/>VPEBSQE04E : ACCOUNT NUMBER MUST BE NUMERIC AND GREATER THAN ZEROS<BR/>VPEBSQE05S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES<BR/>VPEBSQE05E : CARD CONVERT INDICATOR MUST BE 0 OR 1<BR/>VPEBSQE06S : ORG NUMBER NOT FOUND<BR/>VPEBSQE07S : ORG NUMBER NOT FOUND                                     TES<BR/>VPEBSQE08S : ORG NUMBER NOT FOUND<BR/>VPEBSQE09S : REQUESTED ACCOUNTorOFFER-ID NUMBER IS NOT FOUND<BR/>VPEBSQE10S : REQUESTED ACCOUNTorOFFER-ID NUMBER IS NOT FOUND<BR/>VPEBSQE11S : REQUESTED ACCOUNTorOFFER-ID NUMBER IS NOT FOUND<BR/>VPEBSQE12S : OFFER IS NOT FOUND<BR/>VPEBSQE13S : OFFER IS NOT FOUND<BR/>VPEBSQE14S : REQUESTED OFFER-ID IS NOT FOUN
*/
type OfferEnrolmentInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/offerEnrolmentInquiry][%d] offerEnrolmentInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *OfferEnrolmentInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}