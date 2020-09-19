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

// MarketingChannelsDataForCustomerAccountEmbosserAdd1 marketing channels data for customer account embosser add1
//
// swagger:model MarketingChannelsDataForCustomerAccountEmbosserAdd1
type MarketingChannelsDataForCustomerAccountEmbosserAdd1 struct {

	//  Max length = 1, Branch Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Branch *string `json:"branch,omitempty"`

	//  Max length = 1, Cards Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Cards *string `json:"cards,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag1 *string `json:"channelUserFlag1,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag2 *string `json:"channelUserFlag2,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag3 *string `json:"channelUserFlag3,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag4 *string `json:"channelUserFlag4,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag5 *string `json:"channelUserFlag5,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag6 *string `json:"channelUserFlag6,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag7 *string `json:"channelUserFlag7,omitempty"`

	//  Max length = 1, If no valid value is provided, processing will be as follows: In ADD mode the field will be set to the default value indicated in the ORG record In UPDATE mode the field will be ignored and no update to this indicator will occur Valid values: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	ChannelUserFlag8 *string `json:"channelUserFlag8,omitempty"`

	//  Max length = 1, Direct Mail Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	DirectMail *string `json:"directMail,omitempty"`

	//  Max length = 1, Email Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Email *string `json:"email,omitempty"`

	//  Max length = 1, IDTV Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Idtv *string `json:"idtv,omitempty"`

	//  Max length = 1, Internet Ebanking Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	InternetEbank *string `json:"internetEbank,omitempty"`

	//  Max length = 1, Letter Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Letters *string `json:"letters,omitempty"`

	//  Max length = 1, Mobile Phone Ebanking Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	MobilePhoneEbank *string `json:"mobilePhoneEbank,omitempty"`

	//  Max length = 1, Outbound Call Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	OutboundCall *string `json:"outboundCall,omitempty"`

	//  Max length = 1, Pin Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Pin *string `json:"pin,omitempty"`

	//  Max length = 1, SMS Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Sms *string `json:"sms,omitempty"`

	//  Max length = 1, Statement Marketing Channel Option. Values are: 1 = Marketing to the customer is permitted via this channel, 0 = Marketing to the customer is not permitted via this channel
	// Max Length: 1
	// Min Length: 0
	Statements *string `json:"statements,omitempty"`
}

// Validate validates this marketing channels data for customer account embosser add1
func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag7(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelUserFlag8(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectMail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdtv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternetEbank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobilePhoneEbank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundCall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateBranch(formats strfmt.Registry) error {

	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	if err := validate.MinLength("branch", "body", string(*m.Branch), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("branch", "body", string(*m.Branch), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateCards(formats strfmt.Registry) error {

	if swag.IsZero(m.Cards) { // not required
		return nil
	}

	if err := validate.MinLength("cards", "body", string(*m.Cards), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cards", "body", string(*m.Cards), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag1(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag1) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag1", "body", string(*m.ChannelUserFlag1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag1", "body", string(*m.ChannelUserFlag1), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag2(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag2) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag2", "body", string(*m.ChannelUserFlag2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag2", "body", string(*m.ChannelUserFlag2), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag3(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag3) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag3", "body", string(*m.ChannelUserFlag3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag3", "body", string(*m.ChannelUserFlag3), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag4(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag4) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag4", "body", string(*m.ChannelUserFlag4), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag4", "body", string(*m.ChannelUserFlag4), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag5(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag5) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag5", "body", string(*m.ChannelUserFlag5), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag5", "body", string(*m.ChannelUserFlag5), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag6(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag6) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag6", "body", string(*m.ChannelUserFlag6), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag6", "body", string(*m.ChannelUserFlag6), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag7(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag7) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag7", "body", string(*m.ChannelUserFlag7), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag7", "body", string(*m.ChannelUserFlag7), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateChannelUserFlag8(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelUserFlag8) { // not required
		return nil
	}

	if err := validate.MinLength("channelUserFlag8", "body", string(*m.ChannelUserFlag8), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("channelUserFlag8", "body", string(*m.ChannelUserFlag8), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateDirectMail(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectMail) { // not required
		return nil
	}

	if err := validate.MinLength("directMail", "body", string(*m.DirectMail), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("directMail", "body", string(*m.DirectMail), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MinLength("email", "body", string(*m.Email), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "body", string(*m.Email), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateIdtv(formats strfmt.Registry) error {

	if swag.IsZero(m.Idtv) { // not required
		return nil
	}

	if err := validate.MinLength("idtv", "body", string(*m.Idtv), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("idtv", "body", string(*m.Idtv), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateInternetEbank(formats strfmt.Registry) error {

	if swag.IsZero(m.InternetEbank) { // not required
		return nil
	}

	if err := validate.MinLength("internetEbank", "body", string(*m.InternetEbank), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("internetEbank", "body", string(*m.InternetEbank), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateLetters(formats strfmt.Registry) error {

	if swag.IsZero(m.Letters) { // not required
		return nil
	}

	if err := validate.MinLength("letters", "body", string(*m.Letters), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("letters", "body", string(*m.Letters), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateMobilePhoneEbank(formats strfmt.Registry) error {

	if swag.IsZero(m.MobilePhoneEbank) { // not required
		return nil
	}

	if err := validate.MinLength("mobilePhoneEbank", "body", string(*m.MobilePhoneEbank), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("mobilePhoneEbank", "body", string(*m.MobilePhoneEbank), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateOutboundCall(formats strfmt.Registry) error {

	if swag.IsZero(m.OutboundCall) { // not required
		return nil
	}

	if err := validate.MinLength("outboundCall", "body", string(*m.OutboundCall), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("outboundCall", "body", string(*m.OutboundCall), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validatePin(formats strfmt.Registry) error {

	if swag.IsZero(m.Pin) { // not required
		return nil
	}

	if err := validate.MinLength("pin", "body", string(*m.Pin), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("pin", "body", string(*m.Pin), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateSms(formats strfmt.Registry) error {

	if swag.IsZero(m.Sms) { // not required
		return nil
	}

	if err := validate.MinLength("sms", "body", string(*m.Sms), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("sms", "body", string(*m.Sms), 1); err != nil {
		return err
	}

	return nil
}

func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) validateStatements(formats strfmt.Registry) error {

	if swag.IsZero(m.Statements) { // not required
		return nil
	}

	if err := validate.MinLength("statements", "body", string(*m.Statements), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("statements", "body", string(*m.Statements), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketingChannelsDataForCustomerAccountEmbosserAdd1) UnmarshalBinary(b []byte) error {
	var res MarketingChannelsDataForCustomerAccountEmbosserAdd1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
