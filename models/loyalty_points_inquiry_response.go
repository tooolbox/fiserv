// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoyaltyPointsInquiryResponse loyalty points inquiry response
//
// swagger:model LoyaltyPointsInquiryResponse
type LoyaltyPointsInquiryResponse struct {

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	Account string `json:"account,omitempty"`

	//  Max length = 40, Mailing address line 1 of the accountholder.
	Address1 string `json:"address1,omitempty"`

	//  Max length = 40, Mailing address line 2 of the accountholder.
	Address2 string `json:"address2,omitempty"`

	//  Max length = 19, Card number as received in the input message.
	CardNumber string `json:"cardNumber,omitempty"`

	//  Max length = 40, City portion of the mailing address of the account holder. If  NAME ADDRESS SOURCE FLAG is LMS, CITY  can't be spaces.
	City string `json:"city,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Next Statement Date: field that specifies the next date on which statement information is produced.
	DateNextStmt string `json:"dateNextStmt,omitempty"`

	// Format: YYYYMMDD. Effective date of the transaction.
	EffectiveDate string `json:"effectiveDate,omitempty"`

	//  Max length = 17, Last points redeemed.
	LastPntsRdmd string `json:"lastPntsRdmd,omitempty"`

	//  Max length = 19, LMS account number that identifies the Points Account record.
	LmsAcctNbr string `json:"lmsAcctNbr,omitempty"`

	//  Max length = 40, Account name: Last name if the account holder is an individual, or business name if the account holder is a business.
	Name1 string `json:"name1,omitempty"`

	//  Max length = 17, Points available for redemption.
	OpenToRedeemPoints string `json:"openToRedeemPoints,omitempty"`

	//  Max length = 17, Total Cycle-To-Date Bonus Points: total Cycle to Date points for bonus transactions.
	PointsBonusCtd string `json:"pointsBonusCtd,omitempty"`

	//  Max length = 17, Total Cycle-To-Date Points Earned: total Cycle to Date points for earned transactions.
	PointsEarned string `json:"pointsEarned,omitempty"`

	//  Max length = 17, Total Cycle-To-Date Points Expired: total Cycle to Date points for expired transactions.
	PointsExpiredCtd string `json:"pointsExpiredCtd,omitempty"`

	//  Max length = 17, Total Cycle-To-Date Points redeemed: total Cycle to Date points for redeemed transactions.
	PointsRedeemed string `json:"pointsRedeemed,omitempty"`

	//  Max length = 10, Postal code portion of the mailing address of the account holder If  NAME ADDRESS SOURCE FLAG is LMS, POSTAL CODE can't be spaces.
	PostalCode string `json:"postalCode,omitempty"`

	//  Max length = 18, Points expiring on 1st month statement. Calculated based on the points accumulated - points used + points expired.
	PtsExpFirstMonth string `json:"ptsExpFirstMonth,omitempty"`

	//  Max length = 18, Points expiring on 2nd month statement. Calculated based on the points accumulated - points used + points expired.
	PtsExpSecondMonth string `json:"ptsExpSecondMonth,omitempty"`

	//  Max length = 18, Points expiring on 3rd month statement. Calculated based on the points accumulated - points used + points expired.
	PtsExpThirdMonth string `json:"ptsExpThirdMonth,omitempty"`

	//  Max length = 3, Redeem Transaction code. Values are 71 or 51.
	RedemTransCode string `json:"redemTransCode,omitempty"`

	//  Max length = 1, Return code status. Valid values are A, B, C, D, E, F, G, 0.
	ReturnStatusCode string `json:"returnStatusCode,omitempty"`
}

// Validate validates this loyalty points inquiry response
func (m *LoyaltyPointsInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyPointsInquiryResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *LoyaltyPointsInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyPointsInquiryResponse) UnmarshalBinary(b []byte) error {
	var res LoyaltyPointsInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
