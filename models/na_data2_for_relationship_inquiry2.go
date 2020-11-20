// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NaData2ForRelationshipInquiry2 na data2 for relationship inquiry2
//
// swagger:model NaData2ForRelationshipInquiry2
type NaData2ForRelationshipInquiry2 struct {

	//  Max length = 40, Employer: Cardholder's employer name.
	Employer string `json:"employer,omitempty"`

	//  Max length = 40, Address Line 1 of the customer as defined on the Customer Name/Address record.
	NaAddr1 string `json:"naAddr1,omitempty"`

	//  Max length = 40, Address Line 2 of the customer as defined on the Customer Name/Address record.
	NaAddr2 string `json:"naAddr2,omitempty"`

	//  Max length = 3, Age of cardholder. Field is calculated based on Date of Birth.
	NaAge string `json:"naAge,omitempty"`

	//  Max length = 40, City state: City portion of the customer address as defined on the Customer Name/Address record.
	NaCityState string `json:"naCityState,omitempty"`

	//  Max length = 3, Country Code: Country portion of the mailing address.
	NaCntryCd string `json:"naCntryCd,omitempty"`

	//  Max length = 1, Email Flag: code that indicates whether you have permission from the cardholder to Email. Values are:  0 - Do not send  1 - May be sent.  2 - Prefer to be sent.
	NaEmailFlag string `json:"naEmailFlag,omitempty"`

	//  Max length = 20, Employee phone number of the accountholder/customer.
	NaEmpPhone string `json:"naEmpPhone,omitempty"`

	//  Max length = 20, Fax number of the accountholder/customer.
	NaFaxPhone string `json:"naFaxPhone,omitempty"`

	//  Max length = 1, Fax Flag: Code that specifies whether the customer grants permission to use this number. Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	NaFaxPhoneFlag string `json:"naFaxPhoneFlag,omitempty"`

	//  Max length = 1, Gender code. Values are: 0 - Gender not specified 1 - Male 2 - Female
	NaGenderCode string `json:"naGenderCode,omitempty"`

	//  Max length = 20, Home phone number of the accountholder/customer.
	NaHomePhone string `json:"naHomePhone,omitempty"`

	//  Max length = 1, Home Phone Flag: Code that specifies whether the customer grants permission to  call this number. Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method.
	NaHomePhoneFlag string `json:"naHomePhoneFlag,omitempty"`

	//  Max length = 20, Mobile number of the accountholder/customer.
	NaMobilePhone string `json:"naMobilePhone,omitempty"`

	//  Max length = 1, Mobile Phone Flag: Code that specifies whether the customer grants permission to  call this number.  Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	NaMobilePhoneFlag string `json:"naMobilePhoneFlag,omitempty"`

	//  Max length = 40, Name line 1 of the customer as defined on the Customer Name/Address record.
	NaNameLine1 string `json:"naNameLine1,omitempty"`

	//  Max length = 40, Name line 2 of the customer as defined on the Customer Name/Address record.
	NaNameLine2 string `json:"naNameLine2,omitempty"`

	//  Max length = 40, Name line 3 of the customer as defined on the Customer Name/Address record.
	NaNameLine3 string `json:"naNameLine3,omitempty"`

	//  Max length = 1, Employee Phone Flag: Code that specifies whether the customer grants permission to call this number. Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method.
	NaPhoneFlag string `json:"naPhoneFlag,omitempty"`

	//  Max length = 10, Postal Code: Postal code portion of the mailing address.
	NaPstlCd string `json:"naPstlCd,omitempty"`

	//  Max length = 1, SMS Flag: code that indicates whether you have permission from the cardholder to SMS.  Values are:  0 - Do not send  1 - May be sent
	NaSmsFlag string `json:"naSmsFlag,omitempty"`

	//  Max length = 3, City portion of the mailing address for cards and PIN mailers.
	NaState string `json:"naState,omitempty"`

	//  Max length = 25, Tax identification number of the customer.
	Taxid string `json:"taxid,omitempty"`
}

// Validate validates this na data2 for relationship inquiry2
func (m *NaData2ForRelationshipInquiry2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NaData2ForRelationshipInquiry2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NaData2ForRelationshipInquiry2) UnmarshalBinary(b []byte) error {
	var res NaData2ForRelationshipInquiry2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}