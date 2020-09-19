// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NaData2ForAccountDetailInquiry2 na data2 for account detail inquiry2
//
// swagger:model NaData2ForAccountDetailInquiry2
type NaData2ForAccountDetailInquiry2 struct {

	//  Max length = 40, Address Line 1 of the customer as defined on the Customer Name/Address record.
	Addr1 string `json:"addr1,omitempty"`

	//  Max length = 40, Address Line 2 of the customer as defined on the Customer Name/Address record.
	Addr2 string `json:"addr2,omitempty"`

	//  Max length = 40, Address Line 3 of the customer as defined on the Customer Name/Address record.
	Addr3 string `json:"addr3,omitempty"`

	//  Max length = 40, Address Line 4 of the customer as defined on the Customer Name/Address record.
	AddressLine4 string `json:"addressLine4,omitempty"`

	//  Max length = 40, City state: City portion of the customer address as defined on the Customer Name/Address record.
	CityState string `json:"cityState,omitempty"`

	//  Max length = 3, Country Code: Country portion of the mailing address.
	CntryCd string `json:"cntryCd,omitempty"`

	//  Max length = 64, Country Name.
	Country string `json:"country,omitempty"`

	//  Max length = 30, County of the Customer.
	County string `json:"county,omitempty"`

	// Format: YYYYMMDD. Customer's date of birth.
	Dob string `json:"dob,omitempty"`

	//  Max length = 60, Email: E-mail address of the cardholder.
	Email string `json:"email,omitempty"`

	//  Max length = 1, Email Flag: code that indicates whether you have permission from the cardholder to Email. Values are:  0 - Do not send  1 - May be sent.  2 - Prefer to be sent.
	EmailFlag string `json:"emailFlag,omitempty"`

	//  Max length = 20, Employee Phone
	EmpPhone string `json:"empPhone,omitempty"`

	//  Max length = 6, Employee Phone Number Extension.
	EmpPhoneExtn string `json:"empPhoneExtn,omitempty"`

	//  Max length = 40, Employer: Name of the employer of the cardholder.
	Employer string `json:"employer,omitempty"`

	//  Max length = 20, Fax number of the accountholder.
	FaxPhone string `json:"faxPhone,omitempty"`

	//  Max length = 1, Fax Flag: Code that specifies whether the customer grants permission to use this number. Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	FaxPhoneFlag string `json:"faxPhoneFlag,omitempty"`

	//  Max length = 40, First name: Field that identifies the first name of the cardholder.
	FirstName string `json:"firstName,omitempty"`

	//  Max length = 1, Gender code. Values are: 0 - Gender not specified 1 - Male 2 - Female
	GenderCode string `json:"genderCode,omitempty"`

	//  Max length = 20, Customer home phone number
	HomePhone string `json:"homePhone,omitempty"`

	//  Max length = 1, Home Phone Flag: Code that specifies whether the customer grants permission to  call this number.  Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	HomePhoneFlag string `json:"homePhoneFlag,omitempty"`

	//  Max length = 40, House name of the customer.
	HouseName string `json:"houseName,omitempty"`

	//  Max length = 20, House number of the customer.
	HouseNumber string `json:"houseNumber,omitempty"`

	//  Max length = 1, ID Fraud: Value to indicate whether the customer has been the target of identity fraud.  Values are: 0 - Customer has not been the target of identity fraud 1 - Customer has been the target of identity fraud
	IDFraud string `json:"idFraud,omitempty"`

	//  Max length = 40, Last name: field that identifies the last name of the cardholder.
	LastName string `json:"lastName,omitempty"`

	//  Max length = 40, Middle Name: Field that identifies the middle name of the cardholder.
	MidName string `json:"midName,omitempty"`

	//  Max length = 20, Mobile number of the accountholder.
	MobPhone string `json:"mobPhone,omitempty"`

	//  Max length = 1, Mobile Phone Flag: Code that specifies whether the customer grants permission to  call this number.  Values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	MobilePhoneFlag string `json:"mobilePhoneFlag,omitempty"`

	//  Max length = 40, Name line 1 of the customer as defined on the Customer Name/Address record.
	NameLin1 string `json:"nameLin1,omitempty"`

	//  Max length = 40, Name line 2 of the customer as defined on the Customer Name/Address record.
	NameLin2 string `json:"nameLin2,omitempty"`

	//  Max length = 40, Name line 3  of the customer as defined on the Customer Name/Address record.
	NameLin3 string `json:"nameLin3,omitempty"`

	//  Max length = 25, Social Security number (national Insurance number): field that identifies the SSN ,tax identification number or another user-defined identification number of the cardholder.
	NinNbr string `json:"ninNbr,omitempty"`

	//  Max length = 1, Phone Flag: Code that specifies whether the customer grants permission to call this number. The values are: 0 = No, do not contact the customer at this number 1 = Yes, you can contact the customer at this number (Default) 2 = Preferred contact method
	PhoneFlag string `json:"phoneFlag,omitempty"`

	//  Max length = 40, Relative Name: Name of relative of the cardholder.
	RelativeName string `json:"relativeName,omitempty"`

	//  Max length = 1, SMS Flag: code that indicates whether you have permission from the cardholder to SMS.  Values are:  0 - Do not send  1 - May be sent
	SmsFlag string `json:"smsFlag,omitempty"`

	//  Max length = 1, Social Security number (national Insurance number) Flag: code that indicates the type of identification number held in the AMNA-SSAN field.  Values are: 0 - Social Security Number 1 - Tax Id Number 2 - Other ID 3 - National ID 4 - Tabiya ID 5 - Iqama ID 6 - Passport ID 7 - Diplomat ID 8 - National insurance number ID
	SsanFlag string `json:"ssanFlag,omitempty"`

	//  Max length = 3, State: State or province portion of the mailing address.
	State string `json:"state,omitempty"`

	//  Max length = 20, Suffix: Field that identifies the suffix of the cardholder (ex. Jr., Sr.).
	Suffix string `json:"suffix,omitempty"`

	//  Max length = 15, Social Security number (national Insurance number): field that identifies the SSN ,tax identification number or another user-defined identification number of the cardholder.
	Taxid string `json:"taxid,omitempty"`

	//  Max length = 20, Title: Professional or honorary title associated with the cardholder name entered in the NAME LINE 1 field.
	Title string `json:"title,omitempty"`

	//  Max length = 20, User-defined demographic data or other information for the customer.
	UserDemo1 string `json:"userDemo1,omitempty"`

	//  Max length = 40, User-defined demographic data or other information for the customer.
	UserLine1 string `json:"userLine1,omitempty"`

	//  Max length = 40, User-defined fields for an owner.
	UserLine2 string `json:"userLine2,omitempty"`

	//  Max length = 40, User-defined demographic data or other information for the customer.
	UserLine3 string `json:"userLine3,omitempty"`

	//  Max length = 10, ZIP CODE
	ZipCode string `json:"zipCode,omitempty"`
}

// Validate validates this na data2 for account detail inquiry2
func (m *NaData2ForAccountDetailInquiry2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NaData2ForAccountDetailInquiry2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NaData2ForAccountDetailInquiry2) UnmarshalBinary(b []byte) error {
	var res NaData2ForAccountDetailInquiry2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
