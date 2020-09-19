// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthorisationsInquiryResponse authorisations inquiry response
//
// swagger:model AuthorisationsInquiryResponse
type AuthorisationsInquiryResponse struct {

	//  Max length = 1, Token - Next account log
	AccountLogExtNext string `json:"accountLogExtNext,omitempty"`

	//  Max length = 1, Token - Previous Account log
	AccountLogExtPrev string `json:"accountLogExtPrev,omitempty"`

	//  Max length = 9, Token - Next Accounts Relative Byte Address on the Auth log file
	AccountRbaNext string `json:"accountRbaNext,omitempty"`

	//  Max length = 9, Token - Previous Account Relative Bytes Address on the Auth log file
	AccountRbaPrev string `json:"accountRbaPrev,omitempty"`

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	Acct string `json:"acct,omitempty"`

	// AUTHS TABLE
	AuthsTable []*AuthsTableForAuthorisationsInquiry1 `json:"authsTable"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Token - Next authorisation date
	EffDateNext string `json:"effDateNext,omitempty"`

	// Format: YYYYMMDD. Token - Previous authorisation date
	EffDatePrev string `json:"effDatePrev,omitempty"`

	//  Max length = 1, Foreign use indicator: Value indicates whether the incoming action is applied to the local or foreign account. The values are: 0 = Local account (default) 1 = Foreign account SPACE = defaults to 0
	ForeignUse string `json:"foreignUse,omitempty"`

	//  Max length = 1, More data indicator: Indicates that there are more records to display.
	MoreDataFlag string `json:"moreDataFlag,omitempty"`

	//  Max length = 3, Number of items returned by the service
	NbrReturnedItems string `json:"nbrReturnedItems,omitempty"`

	//  Max length = 3, Organisation currency ISO number
	OutCurrNbr string `json:"outCurrNbr,omitempty"`

	//  Max length = 1, Organisation currency nod
	OutCurrNod string `json:"outCurrNod,omitempty"`

	//  Max length = 7, Token - Next authorisation time
	TimeNext string `json:"timeNext,omitempty"`

	//  Max length = 7, Token - Previous authorisation time
	TimePrev string `json:"timePrev,omitempty"`

	//  Max length = 1, Token - Next transaction type
	TxnRecTypeNext string `json:"txnRecTypeNext,omitempty"`

	//  Max length = 1, Token - Previous Transaction type
	TxnRecTypePrev string `json:"txnRecTypePrev,omitempty"`
}

// Validate validates this authorisations inquiry response
func (m *AuthorisationsInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthsTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorisationsInquiryResponse) validateAuthsTable(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthsTable) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthsTable); i++ {
		if swag.IsZero(m.AuthsTable[i]) { // not required
			continue
		}

		if m.AuthsTable[i] != nil {
			if err := m.AuthsTable[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authsTable" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorisationsInquiryResponse) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.Common) { // not required
		return nil
	}

	if m.Common != nil {
		if err := m.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorisationsInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorisationsInquiryResponse) UnmarshalBinary(b []byte) error {
	var res AuthorisationsInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
