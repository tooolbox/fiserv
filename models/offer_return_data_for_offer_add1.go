// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OfferReturnDataForOfferAdd1 offer return data for offer add1
//
// swagger:model OfferReturnDataForOfferAdd1
type OfferReturnDataForOfferAdd1 struct {

	//  Max length = 1, Status of the Offer Enrollment record. The values are: 0 = Add pending 1 = Add complete
	OfferAddResult string `json:"offerAddResult,omitempty"`

	//  Max length = 10, OFFER CODE
	OfferCode string `json:"offerCode,omitempty"`
}

// Validate validates this offer return data for offer add1
func (m *OfferReturnDataForOfferAdd1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OfferReturnDataForOfferAdd1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OfferReturnDataForOfferAdd1) UnmarshalBinary(b []byte) error {
	var res OfferReturnDataForOfferAdd1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
