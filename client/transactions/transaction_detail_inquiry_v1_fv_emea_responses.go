// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tooolbox/firstvision/models"
)

// TransactionDetailInquiryV1FvEmeaReader is a Reader for the TransactionDetailInquiryV1FvEmea structure.
type TransactionDetailInquiryV1FvEmeaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransactionDetailInquiryV1FvEmeaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransactionDetailInquiryV1FvEmeaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTransactionDetailInquiryV1FvEmeaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTransactionDetailInquiryV1FvEmeaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTransactionDetailInquiryV1FvEmeaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTransactionDetailInquiryV1FvEmeaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTransactionDetailInquiryV1FvEmeaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 452:
		result := NewTransactionDetailInquiryV1FvEmeaStatus452()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 453:
		result := NewTransactionDetailInquiryV1FvEmeaStatus453()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 455:
		result := NewTransactionDetailInquiryV1FvEmeaStatus455()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 465:
		result := NewTransactionDetailInquiryV1FvEmeaStatus465()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTransactionDetailInquiryV1FvEmeaOK creates a TransactionDetailInquiryV1FvEmeaOK with default headers values
func NewTransactionDetailInquiryV1FvEmeaOK() *TransactionDetailInquiryV1FvEmeaOK {
	return &TransactionDetailInquiryV1FvEmeaOK{}
}

/*TransactionDetailInquiryV1FvEmeaOK handles this case with default header values.

successful operation
*/
type TransactionDetailInquiryV1FvEmeaOK struct {
	Payload *models.TransactionDetailInquiryResponse
}

func (o *TransactionDetailInquiryV1FvEmeaOK) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaOK  %+v", 200, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaOK) GetPayload() *models.TransactionDetailInquiryResponse {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionDetailInquiryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaBadRequest creates a TransactionDetailInquiryV1FvEmeaBadRequest with default headers values
func NewTransactionDetailInquiryV1FvEmeaBadRequest() *TransactionDetailInquiryV1FvEmeaBadRequest {
	return &TransactionDetailInquiryV1FvEmeaBadRequest{}
}

/*TransactionDetailInquiryV1FvEmeaBadRequest handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-003 - Invalid request. This could be because of invalid data in the request or header, among other reasons.<BR/>API-Sec-004 - Invalid request. The client_assertion JWT is invalid<BR/>API-Sec-013 - The JWT has an invalid expiry. Pleases end a request with a valid JWT
*/
type TransactionDetailInquiryV1FvEmeaBadRequest struct {
	Payload *models.FsErrorDetails
}

func (o *TransactionDetailInquiryV1FvEmeaBadRequest) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaBadRequest  %+v", 400, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaBadRequest) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaUnauthorized creates a TransactionDetailInquiryV1FvEmeaUnauthorized with default headers values
func NewTransactionDetailInquiryV1FvEmeaUnauthorized() *TransactionDetailInquiryV1FvEmeaUnauthorized {
	return &TransactionDetailInquiryV1FvEmeaUnauthorized{}
}

/*TransactionDetailInquiryV1FvEmeaUnauthorized handles this case with default header values.

Unauthorized: Authorization failed due to missing or invalid credentials. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-001 - The auth header has invalid values. Please try with valid values<BR/>API-Sec-002 - The client id provided is invalid. Please try with valid client id<BR/>API-Sec-007 - The access token is invalid. It may have expired, or maybe incorrect
*/
type TransactionDetailInquiryV1FvEmeaUnauthorized struct {
	Payload *models.FsErrorDetails
}

func (o *TransactionDetailInquiryV1FvEmeaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaUnauthorized  %+v", 401, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaUnauthorized) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaForbidden creates a TransactionDetailInquiryV1FvEmeaForbidden with default headers values
func NewTransactionDetailInquiryV1FvEmeaForbidden() *TransactionDetailInquiryV1FvEmeaForbidden {
	return &TransactionDetailInquiryV1FvEmeaForbidden{}
}

/*TransactionDetailInquiryV1FvEmeaForbidden handles this case with default header values.

Forbidden: Insufficient access for requested operation. <BR/>Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-006 - Replay attack detected. Please try again with valid data<BR/>API-Sec-008 - Mismatch in scope. The access token does not have the permission to access this resource
*/
type TransactionDetailInquiryV1FvEmeaForbidden struct {
	Payload *models.FsErrorDetails
}

func (o *TransactionDetailInquiryV1FvEmeaForbidden) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaForbidden  %+v", 403, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaForbidden) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaNotFound creates a TransactionDetailInquiryV1FvEmeaNotFound with default headers values
func NewTransactionDetailInquiryV1FvEmeaNotFound() *TransactionDetailInquiryV1FvEmeaNotFound {
	return &TransactionDetailInquiryV1FvEmeaNotFound{}
}

/*TransactionDetailInquiryV1FvEmeaNotFound handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-005 - Resource not found
*/
type TransactionDetailInquiryV1FvEmeaNotFound struct {
	Payload *models.FsErrorDetails
}

func (o *TransactionDetailInquiryV1FvEmeaNotFound) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaNotFound  %+v", 404, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaNotFound) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaTooManyRequests creates a TransactionDetailInquiryV1FvEmeaTooManyRequests with default headers values
func NewTransactionDetailInquiryV1FvEmeaTooManyRequests() *TransactionDetailInquiryV1FvEmeaTooManyRequests {
	return &TransactionDetailInquiryV1FvEmeaTooManyRequests{}
}

/*TransactionDetailInquiryV1FvEmeaTooManyRequests handles this case with default header values.

Below are the error codes you can expect when connecting over internet:<BR/><BR/>API-Sec-009 - You have exceeded the App level quota<BR/>API-Sec-010 - You have exceeded the Developer level quota<BR/>API-Sec-011 - You have exceeded the Product level quota<BR/>API-Sec-012 - Too many requests for this API. Please try after sometime
*/
type TransactionDetailInquiryV1FvEmeaTooManyRequests struct {
	Payload *models.FsErrorDetails
}

func (o *TransactionDetailInquiryV1FvEmeaTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaTooManyRequests  %+v", 429, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaTooManyRequests) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaStatus452 creates a TransactionDetailInquiryV1FvEmeaStatus452 with default headers values
func NewTransactionDetailInquiryV1FvEmeaStatus452() *TransactionDetailInquiryV1FvEmeaStatus452 {
	return &TransactionDetailInquiryV1FvEmeaStatus452{}
}

/*TransactionDetailInquiryV1FvEmeaStatus452 handles this case with default header values.

System Exception: Internal processes not related to client interaction with application are the cause of the failure
*/
type TransactionDetailInquiryV1FvEmeaStatus452 struct {
}

func (o *TransactionDetailInquiryV1FvEmeaStatus452) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaStatus452 ", 452)
}

func (o *TransactionDetailInquiryV1FvEmeaStatus452) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaStatus453 creates a TransactionDetailInquiryV1FvEmeaStatus453 with default headers values
func NewTransactionDetailInquiryV1FvEmeaStatus453() *TransactionDetailInquiryV1FvEmeaStatus453 {
	return &TransactionDetailInquiryV1FvEmeaStatus453{}
}

/*TransactionDetailInquiryV1FvEmeaStatus453 handles this case with default header values.

Validation Exception: The request failed validation, modify the request and resubmit
*/
type TransactionDetailInquiryV1FvEmeaStatus453 struct {
}

func (o *TransactionDetailInquiryV1FvEmeaStatus453) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaStatus453 ", 453)
}

func (o *TransactionDetailInquiryV1FvEmeaStatus453) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaStatus455 creates a TransactionDetailInquiryV1FvEmeaStatus455 with default headers values
func NewTransactionDetailInquiryV1FvEmeaStatus455() *TransactionDetailInquiryV1FvEmeaStatus455 {
	return &TransactionDetailInquiryV1FvEmeaStatus455{}
}

/*TransactionDetailInquiryV1FvEmeaStatus455 handles this case with default header values.

ODS Error Exception: ODS returned a message with transaction status ERROR
*/
type TransactionDetailInquiryV1FvEmeaStatus455 struct {
}

func (o *TransactionDetailInquiryV1FvEmeaStatus455) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaStatus455 ", 455)
}

func (o *TransactionDetailInquiryV1FvEmeaStatus455) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransactionDetailInquiryV1FvEmeaStatus465 creates a TransactionDetailInquiryV1FvEmeaStatus465 with default headers values
func NewTransactionDetailInquiryV1FvEmeaStatus465() *TransactionDetailInquiryV1FvEmeaStatus465 {
	return &TransactionDetailInquiryV1FvEmeaStatus465{}
}

/*TransactionDetailInquiryV1FvEmeaStatus465 handles this case with default header values.

Backend Response Code and message in odsMessages field: <BR/><BR/>VPL1S0100S : SERVICE REQUEST LENGTH ERROR                                     SERVICE INPUT FOR SERVICE IS AN INCORRECT LENGTH<BR/>VPL5S0003S : APPLICATION IN NO-PROCESSING STATUS RE-TRY IN A FEW MINUTES     CMS IS NO-PROCESSING STATE<BR/>VPL5S0005E : INPUT CARDorACCT NUMBER IS NOT NUMERIC                            INPUT ACCOUNT NUMBER IS NOT NUMERIC<BR/>VPL5S0006S : REQUESTED ORG NUMBER IS NOT FOUND                                USER IS NOT ALLOWED TO ACCESS THE ORG<BR/>VPL5S0008E : ORGANIZATION NOT DETERMINED<BR/>VPL5S0012S : ACCOUNT NOT FOUND ON FILE<BR/>VPL5S0015E : REQUESTED FUNCTION TYPE IS NOT VALID<BR/>VPL5S0041E : NO STATEMENTS FOUND                                              NO STATEMENTS FOUND<BR/>VPL5S0042E : INVALID TRANSACTION DETAIL REQUEST                               INVALID TRANSACTION DETAIL REQUEST - OPTION NEXT<BR/>VPL5S0043E : NUMBER TRANSACTIONS REQUESTED IS INVALID<BR/>VPL5S0056E : INVALID LOCALorFOREIGN INDICATOR SPECIFIED<BR/>VPL5S0059E : START JULIAN IS GREATER THAN END JULIAN<BR/>VPL5S0060S : TRANSACTION DETAIL BEGIN DATE FORMAT MUST BE CCYYMMDD<BR/>VPL5S0061S : TRANSACTION DETAIL END DATE FORMAT MUST BE CCYYMMDD<BR/>VPL5S0063S : TRANSACTION DETAIL BEGIN DATE IS NOT NUMERIC<BR/>VPL5S0064S : TRANSACTION DATE THRU SHOULD BE EQUAL TO 99999999<BR/>VPL5S0067E : TRANSACTION NBR MONTHS SHOULD BE VALID WHEN TRANSACTION DETAIL'M'STATEMENT OCCURENCE NUMBER MUST BE A NUMERIC VALUE WHEN THE SERVICE REQUEST IS FOR STATEMENT HISTORY BY MONTH.<BR/>VPL5S0068E : WHILE DOING START BROWSE LARGE FILE RECORD NOT FOUND             STATEMENT NUMBER REQUESTED IS GREATER THAN THE NUMBER OF STATEMENTS ON FILE FOR ACCOUNT.<BR/>VPL5S0069E : STATEMENT OCCURENCE NUMBER NOT ON FILE                           THE REQUESTED ACCOUNT DOES NOT HAVE ANY STATEMENT HISTORY ON FILE.<BR/>VPL5S0072E : STATEMENT OCCURENCE NUMBER IS INVALID WHEN REQUEST NOT FOR MONTH STATEMENT OCCURENCE NUMBER IS NOT VALID WHEN THE REQUEST IS HAS NOT BEEN SUBMITTED FOR MONTHSTATEMENT PERIOD.<BR/>VPL5S0073E : TRANSACTION DATE FROM SHOULD BE ZEROS WHEN TRANSACTION DETAIL 'M'TRANSACTION END DATE MUST NOT BE PRESENT WHEN THE REQUEST IS FOR MONTH STATEMENT PERIOD.<BR/>VPL5S0074E : TRANSACTION DATE THRU SHOULD BE ZEROS WHEN TRANSACTION DETAIL 'M'TRANSACTION START DATE WAS PROVIDED WITH A REQUEST TYPE FOR TRANSACTIO DETAIL BY MONTH.   START DATE IS INVALID FOR THIS REQUEST.<BR/>VPL5S9001S : ERROR ACCESSING TRANSACTION FILE<BR/>VPL5S9002S : ERROR ACCESSING ACCOUNT FILE<BR/>VPL5S9003S : ERROR ACCESSING XRF TRANSACTION FILE<BR/>VPL5S9004S : ERROR ACCESSING STATEMENT FILE<BR/>VPL5S9005S : FILE-TABLE RECORD NOT FOUND<BR/>VPL5S9006S : ORG RECORD NOT ON FILE<BR/>VPL5S9007S : CMS LOGO RECORD NOT ON FIL
*/
type TransactionDetailInquiryV1FvEmeaStatus465 struct {
	Payload *models.FsErrorDetails
}

func (o *TransactionDetailInquiryV1FvEmeaStatus465) Error() string {
	return fmt.Sprintf("[POST /fv_emea/v1/transactionDetailInquiry][%d] transactionDetailInquiryV1FvEmeaStatus465  %+v", 465, o.Payload)
}

func (o *TransactionDetailInquiryV1FvEmeaStatus465) GetPayload() *models.FsErrorDetails {
	return o.Payload
}

func (o *TransactionDetailInquiryV1FvEmeaStatus465) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
