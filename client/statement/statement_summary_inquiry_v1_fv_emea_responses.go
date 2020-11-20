// Code generated by go-swagger; DO NOT EDIT.

package statement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// StatementSummaryInquiryV1FvEmeaReader is a Reader for the StatementSummaryInquiryV1FvEmea structure.
type StatementSummaryInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatementSummaryInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatementSummaryInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatementSummaryInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatementSummaryInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatementSummaryInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatementSummaryInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewStatementSummaryInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewStatementSummaryInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewStatementSummaryInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewStatementSummaryInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewStatementSummaryInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStatementSummaryInquiryV1FvEmeaOK creates a StatementSummaryInquiryV1FvEmeaOK with default headers values
func NewStatementSummaryInquiryV1FvEmeaOK() *StatementSummaryInquiryV1FvEmeaOK {
	return &StatementSummaryInquiryV1FvEmeaOK{}
}

/*StatementSummaryInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type StatementSummaryInquiryV1FvEmeaOK struct {
	Payload *models.StatementSummaryInquiryResponse
}

func (o *StatementSummaryInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaOK) GetPayload() *models.StatementSummaryInquiryResponse {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatementSummaryInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaBadRequest creates a StatementSummaryInquiryV1FvEmeaBadRequest with default headers values
func NewStatementSummaryInquiryV1FvEmeaBadRequest() *StatementSummaryInquiryV1FvEmeaBadRequest {
	return &StatementSummaryInquiryV1FvEmeaBadRequest{}
}

/*StatementSummaryInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type StatementSummaryInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *StatementSummaryInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaUnauthorized creates a StatementSummaryInquiryV1FvEmeaUnauthorized with default headers values
func NewStatementSummaryInquiryV1FvEmeaUnauthorized() *StatementSummaryInquiryV1FvEmeaUnauthorized {
	return &StatementSummaryInquiryV1FvEmeaUnauthorized{}
}

/*StatementSummaryInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type StatementSummaryInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *StatementSummaryInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaForbidden creates a StatementSummaryInquiryV1FvEmeaForbidden with default headers values
func NewStatementSummaryInquiryV1FvEmeaForbidden() *StatementSummaryInquiryV1FvEmeaForbidden {
	return &StatementSummaryInquiryV1FvEmeaForbidden{}
}

/*StatementSummaryInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type StatementSummaryInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *StatementSummaryInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaNotFound creates a StatementSummaryInquiryV1FvEmeaNotFound with default headers values
func NewStatementSummaryInquiryV1FvEmeaNotFound() *StatementSummaryInquiryV1FvEmeaNotFound {
	return &StatementSummaryInquiryV1FvEmeaNotFound{}
}

/*StatementSummaryInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type StatementSummaryInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *StatementSummaryInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaTooManyRequests creates a StatementSummaryInquiryV1FvEmeaTooManyRequests with default headers values
func NewStatementSummaryInquiryV1FvEmeaTooManyRequests() *StatementSummaryInquiryV1FvEmeaTooManyRequests {
	return &StatementSummaryInquiryV1FvEmeaTooManyRequests{}
}

/*StatementSummaryInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type StatementSummaryInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *StatementSummaryInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaStatus452 creates a StatementSummaryInquiryV1FvEmeaStatus452 with default headers values
func NewStatementSummaryInquiryV1FvEmeaStatus452() *StatementSummaryInquiryV1FvEmeaStatus452 {
	return &StatementSummaryInquiryV1FvEmeaStatus452{}
}

/*StatementSummaryInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type StatementSummaryInquiryV1FvEmeaStatus452 struct {
}

func (o *StatementSummaryInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaStatus452 ", 452)
}

func (o *StatementSummaryInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaStatus453 creates a StatementSummaryInquiryV1FvEmeaStatus453 with default headers values
func NewStatementSummaryInquiryV1FvEmeaStatus453() *StatementSummaryInquiryV1FvEmeaStatus453 {
	return &StatementSummaryInquiryV1FvEmeaStatus453{}
}

/*StatementSummaryInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type StatementSummaryInquiryV1FvEmeaStatus453 struct {
}

func (o *StatementSummaryInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaStatus453 ", 453)
}

func (o *StatementSummaryInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaStatus455 creates a StatementSummaryInquiryV1FvEmeaStatus455 with default headers values
func NewStatementSummaryInquiryV1FvEmeaStatus455() *StatementSummaryInquiryV1FvEmeaStatus455 {
	return &StatementSummaryInquiryV1FvEmeaStatus455{}
}

/*StatementSummaryInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type StatementSummaryInquiryV1FvEmeaStatus455 struct {
}

func (o *StatementSummaryInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaStatus455 ", 455)
}

func (o *StatementSummaryInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatementSummaryInquiryV1FvEmeaStatus465 creates a StatementSummaryInquiryV1FvEmeaStatus465 with default headers values
func NewStatementSummaryInquiryV1FvEmeaStatus465() *StatementSummaryInquiryV1FvEmeaStatus465 {
	return &StatementSummaryInquiryV1FvEmeaStatus465{}
}

/*StatementSummaryInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPE1SSD00S : SERVICE REQUEST LENGTH ERROR<BR/>VPE5SSD02S : FILE-TABLE USED BY THIS SERVICE IS NOT FOUND<BR/>VPE5SSD04S : ORGANIZATION NUMBER MUST BE NUMERIC AND VALID VALUES ARE 001-998<BR/>VPE5SSD03E : INPUT CARDorACCT NUMBER IS NOT NUMERIC<BR/>VPE5SSD05E : REQUESTED FUNCTION TYPE IS NOT VALID<BR/>VPE5SSD06E : INVALID LOCALorFOREIGN INDICATOR SPECIFIED<BR/>VPE5SSD07E : DATE MUST BE VALID<BR/>VPE5SSD08E : STATEMENT DATE MUST BE NUMERIC AND GREATER THAN ZERO<BR/>VPE5SSD09E : STATEMENT MUST BE VALID FOR OPTION MONTH ONLY<BR/>VPE5SSD10E : STATEMENT MUST BE VALID FOR OPTION DETAIL ONLY<BR/>VPE5SSD11E : STATEMENT NBR OF MONTHS SHOULD BE NUMERIC<BR/>VPE5SSD12E : ORGANIZATION NUMBER NOT DETERMINED<BR/>VPE5SSD13S : USER IS NOT ALLOWED TO ACCESS THE ORG<BR/>VPE5SSD14S : CMS FILE-TABLE RECORD NOT FOUND<BR/>VPE5SSD15S : CMS IS NO-PROCESSING STATE TRY LATER<BR/>VPE5SSD16S : CMS ORG RECORD NOT ON FILE<BR/>VPE5SSD17S : FOREIGN ORGANIZATION NUMBER NOT FOUND<BR/>VPE5SSD18S : ACCOUNT NOT FOUND ON FILE<BR/>VPE5SSD19S : CMS LOGO RECORD NOT ON FILE<BR/>VPE5SSD20E : SERVICE RETURN CODE FAILED<BR/>VPE5SSD21E : NO STATEMENTS FOUND<BR/>VPE5SSD22E : ONLINE BASE SEGMENT RECORDS NOT FOUND<BR/>VPE5SSD23E : NBR OF STATEMENTS LESSER THAN REQUESTED<BR/>VPE5SSD24E : NO STATEMENTS FOUND FOR MONTHS REQUESTE
*/
type StatementSummaryInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *StatementSummaryInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/statementSummaryInquiry][%d] statementSummaryInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *StatementSummaryInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *StatementSummaryInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}