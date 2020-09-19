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

// LettersHistoryResponse letters history response
//
// swagger:model LettersHistoryResponse
type LettersHistoryResponse struct {

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 4, File Name. Assocaited with page scroll and by the service in response message.
	FileName string `json:"fileName,omitempty"`

	// Historical Response Entry Group.
	HistRespEntryGroup []*HistRespEntryGroupForLettersHistory1 `json:"histRespEntryGroup"`

	//  Max length = 2, HIST RESP NBR OCCUR
	HistRespNbrOccur string `json:"histRespNbrOccur,omitempty"`

	// Format: CCYYMMDDC. Historical Response Start Date. Assocaited with page scroll and by the service in response message.
	HistRespStartDate string `json:"histRespStartDate,omitempty"`

	//  Max length = 9, Historical Response Start Timestamp.Assocaited with page scroll and by the service in response message.
	HistRespStartTimestamp string `json:"histRespStartTimestamp,omitempty"`
}

// Validate validates this letters history response
func (m *LettersHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistRespEntryGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LettersHistoryResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *LettersHistoryResponse) validateHistRespEntryGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.HistRespEntryGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.HistRespEntryGroup); i++ {
		if swag.IsZero(m.HistRespEntryGroup[i]) { // not required
			continue
		}

		if m.HistRespEntryGroup[i] != nil {
			if err := m.HistRespEntryGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("histRespEntryGroup" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LettersHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LettersHistoryResponse) UnmarshalBinary(b []byte) error {
	var res LettersHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
