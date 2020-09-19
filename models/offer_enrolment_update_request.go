// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OfferEnrolmentUpdateRequest offer enrolment update request
//
// swagger:model OfferEnrolmentUpdateRequest
type OfferEnrolmentUpdateRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero. This can be either Account number or Card number.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr"`

	// card convert
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardConvert *string `json:"cardConvert,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Effective date of the add-on product. This date can be a date in the future. CMS assesses the first premium in the cycle that this date occurs.
	EffDate string `json:"effDate,omitempty"`

	// enrl cmws sw
	// Max Length: 1
	// Min Length: 0
	EnrlCmwsSw *string `json:"enrlCmwsSw,omitempty"`

	//  Max length = 1, Enrollment Action: Values are: 0 = No action required (Default) 1 = Waiting for customer enrollment 2 = Customer accepted the offer 3 = Customer rejected the offer 4 = Offer cancelled by customer representative.
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	EnrollmentAction *string `json:"enrollmentAction,omitempty"`

	//  Max length = 10, Offer id for this request: Identification number of the product being offered.
	// Required: true
	// Max Length: 10
	// Min Length: 0
	OfferID *string `json:"offerId"`

	//  Max length = 2, Reject Reason: field to identify the reason for which the offer was rejected. A value greater than 00 may be entered when the enroll action field is set to a value of 3 - reject. Valid values is 00 - No reason ( Default)  01 - Not interested 02 - Too expensive 03 to 99 - User defined
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	RejectReason *string `json:"rejectReason,omitempty"`

	//  Max length = 2, RESPONSE CHANNEL
	// Max Length: 2
	// Min Length: 0
	ResponseChannel *string `json:"responseChannel,omitempty"`

	//  Max length = 10, Offer Enrollment User field (1).
	// Max Length: 10
	// Min Length: 0
	User1 *string `json:"user1,omitempty"`

	//  Max length = 10, Offer Enrollment User field (2).
	// Max Length: 10
	// Min Length: 0
	User2 *string `json:"user2,omitempty"`
}

// Validate validates this offer enrolment update request
func (m *OfferEnrolmentUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardConvert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrlCmwsSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrollmentAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateAcctNbr(formats strfmt.Registry) error {

	if err := validate.Required("acctNbr", "body", m.AcctNbr); err != nil {
		return err
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateCardConvert(formats strfmt.Registry) error {

	if swag.IsZero(m.CardConvert) { // not required
		return nil
	}

	if err := validate.MinLength("cardConvert", "body", string(*m.CardConvert), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardConvert", "body", string(*m.CardConvert), 1); err != nil {
		return err
	}

	if err := validate.Pattern("cardConvert", "body", string(*m.CardConvert), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *OfferEnrolmentUpdateRequest) validateEnrlCmwsSw(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrlCmwsSw) { // not required
		return nil
	}

	if err := validate.MinLength("enrlCmwsSw", "body", string(*m.EnrlCmwsSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("enrlCmwsSw", "body", string(*m.EnrlCmwsSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateEnrollmentAction(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrollmentAction) { // not required
		return nil
	}

	if err := validate.MinLength("enrollmentAction", "body", string(*m.EnrollmentAction), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("enrollmentAction", "body", string(*m.EnrollmentAction), 1); err != nil {
		return err
	}

	if err := validate.Pattern("enrollmentAction", "body", string(*m.EnrollmentAction), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateOfferID(formats strfmt.Registry) error {

	if err := validate.Required("offerId", "body", m.OfferID); err != nil {
		return err
	}

	if err := validate.MinLength("offerId", "body", string(*m.OfferID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("offerId", "body", string(*m.OfferID), 10); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateRejectReason(formats strfmt.Registry) error {

	if swag.IsZero(m.RejectReason) { // not required
		return nil
	}

	if err := validate.MinLength("rejectReason", "body", string(*m.RejectReason), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("rejectReason", "body", string(*m.RejectReason), 2); err != nil {
		return err
	}

	if err := validate.Pattern("rejectReason", "body", string(*m.RejectReason), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateResponseChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseChannel) { // not required
		return nil
	}

	if err := validate.MinLength("responseChannel", "body", string(*m.ResponseChannel), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("responseChannel", "body", string(*m.ResponseChannel), 2); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateUser1(formats strfmt.Registry) error {

	if swag.IsZero(m.User1) { // not required
		return nil
	}

	if err := validate.MinLength("user1", "body", string(*m.User1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("user1", "body", string(*m.User1), 10); err != nil {
		return err
	}

	return nil
}

func (m *OfferEnrolmentUpdateRequest) validateUser2(formats strfmt.Registry) error {

	if swag.IsZero(m.User2) { // not required
		return nil
	}

	if err := validate.MinLength("user2", "body", string(*m.User2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("user2", "body", string(*m.User2), 10); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OfferEnrolmentUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OfferEnrolmentUpdateRequest) UnmarshalBinary(b []byte) error {
	var res OfferEnrolmentUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
