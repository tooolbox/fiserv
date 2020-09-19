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

// TransactionDetailInquiryRequest2 transaction detail inquiry request2
//
// swagger:model TransactionDetailInquiryRequest2
type TransactionDetailInquiryRequest2 struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero. This can be either Account number or Card number. Enter an account or card number to search for a match on either the Embosser record or the Account Base Segment record
	// Required: true
	// Max Length: 19
	// Min Length: 0
	Account *string `json:"account"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Date From: Start date of period of returned transactions, based on effective date of the transaction. If provided, TXN_DETAIL flag should be C, R, B or space.  Date format is CCYY-MM-DD.
	DtFrom string `json:"dtFrom,omitempty"`

	// Format: YYYYMMDD. Date Through: End date of returned transactions, based on effective date of the transaction. If provided, TXN_DETAIL flag should be C, R, B, Space. Date should be greater than Start Date.  Date format is CCYY-MM-DD.
	DtThru string `json:"dtThru,omitempty"`

	//  Max length = 1, Dual Indicator: Dual Account Flag of account.Valid values are: L = Local account (default) F = Foreign account If dual currency is not being used or if the field is left blank, the value is L.
	// Max Length: 1
	// Min Length: 0
	DualInd *string `json:"dualInd,omitempty"`

	//  Max length = 3, Number of Transactions: Number of transactions to be returned with the output service. Maximum number allowed is 50.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	NbrTxns *string `json:"nbrTxns,omitempty"`

	//  Max length = 1, This field is used to enable the billing account processing. Valid values are <UL><LI><EM>0 -</EM> Process Normal Accounts </LI><LI><EM>1 -</EM> Process Sub Accounts </LI></UL>
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ProcBillAcct *string `json:"procBillAcct,omitempty"`

	//  Max length = 3, This field is used to return the index of last sub account processed. For Normal accounts index will be zero.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	SubAcctIdx *string `json:"subAcctIdx,omitempty"`

	//  Max length = 1, Transaction Detail settings: Parameter to define period of returned transactions.  Valid values are C, R, M, B  and Space: C = CURRENT- Cycle to Date R = RECENT- Statemented M = MONTH- Statemented transaction of specific month. B = BOTH- Cycle to date and Statemented,.  If space, default is 'B'
	// Max Length: 1
	// Min Length: 0
	TxnDtl *string `json:"txnDtl,omitempty"`

	//  Max length = 1, Transaction File Type. Populated from a previous service request-for internal use only Values are: C = Use current AMOS file S = Use small statement file L = Use large statement file
	// Max Length: 1
	// Min Length: 0
	TxnFileTyp *string `json:"txnFileTyp,omitempty"`

	//  Max length = 2, Transaction Month number:  Statement Month Number to be  provided if TXN-DETAIL is 'M'. Relative statement number can be used if the actual statement dates are not known, i.e. : 00 - Current statement01 - Previous months statement. 02 - Second previous months statement, etc.
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	TxnNbrMths *string `json:"txnNbrMths,omitempty"`

	// Format: YYYYMMDD. Transaction Start Date. Populated from a previous service request-for internal use only. Date format is CCYY-MM-DD.
	TxnStartDt string `json:"txnStartDt,omitempty"`

	//  Max length = 5, Transaction Start Number. Populated from a previous service request-for internal use only
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	TxnStartNbr *string `json:"txnStartNbr,omitempty"`

	//  Max length = 1, Transaction Details Suppress Flag. Valid values are: Y - Suppress transactions details that are marked as requiring suppression from being returned.  N - (default if not supplied) No need to Suppress
	// Max Length: 1
	// Min Length: 0
	TxnSuppFlg *string `json:"txnSuppFlg,omitempty"`
}

// Validate validates this transaction detail inquiry request2
func (m *TransactionDetailInquiryRequest2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNbrTxns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcBillAcct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAcctIdx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxnDtl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxnFileTyp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxnNbrMths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxnStartNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxnSuppFlg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionDetailInquiryRequest2) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	if err := validate.MinLength("account", "body", string(*m.Account), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("account", "body", string(*m.Account), 19); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateCommon(formats strfmt.Registry) error {

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

func (m *TransactionDetailInquiryRequest2) validateDualInd(formats strfmt.Registry) error {

	if swag.IsZero(m.DualInd) { // not required
		return nil
	}

	if err := validate.MinLength("dualInd", "body", string(*m.DualInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("dualInd", "body", string(*m.DualInd), 1); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateNbrTxns(formats strfmt.Registry) error {

	if swag.IsZero(m.NbrTxns) { // not required
		return nil
	}

	if err := validate.MinLength("nbrTxns", "body", string(*m.NbrTxns), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("nbrTxns", "body", string(*m.NbrTxns), 3); err != nil {
		return err
	}

	if err := validate.Pattern("nbrTxns", "body", string(*m.NbrTxns), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateProcBillAcct(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcBillAcct) { // not required
		return nil
	}

	if err := validate.MinLength("procBillAcct", "body", string(*m.ProcBillAcct), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("procBillAcct", "body", string(*m.ProcBillAcct), 1); err != nil {
		return err
	}

	if err := validate.Pattern("procBillAcct", "body", string(*m.ProcBillAcct), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateSubAcctIdx(formats strfmt.Registry) error {

	if swag.IsZero(m.SubAcctIdx) { // not required
		return nil
	}

	if err := validate.MinLength("subAcctIdx", "body", string(*m.SubAcctIdx), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("subAcctIdx", "body", string(*m.SubAcctIdx), 3); err != nil {
		return err
	}

	if err := validate.Pattern("subAcctIdx", "body", string(*m.SubAcctIdx), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateTxnDtl(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnDtl) { // not required
		return nil
	}

	if err := validate.MinLength("txnDtl", "body", string(*m.TxnDtl), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("txnDtl", "body", string(*m.TxnDtl), 1); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateTxnFileTyp(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnFileTyp) { // not required
		return nil
	}

	if err := validate.MinLength("txnFileTyp", "body", string(*m.TxnFileTyp), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("txnFileTyp", "body", string(*m.TxnFileTyp), 1); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateTxnNbrMths(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnNbrMths) { // not required
		return nil
	}

	if err := validate.MinLength("txnNbrMths", "body", string(*m.TxnNbrMths), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("txnNbrMths", "body", string(*m.TxnNbrMths), 2); err != nil {
		return err
	}

	if err := validate.Pattern("txnNbrMths", "body", string(*m.TxnNbrMths), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateTxnStartNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnStartNbr) { // not required
		return nil
	}

	if err := validate.MinLength("txnStartNbr", "body", string(*m.TxnStartNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("txnStartNbr", "body", string(*m.TxnStartNbr), 5); err != nil {
		return err
	}

	if err := validate.Pattern("txnStartNbr", "body", string(*m.TxnStartNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *TransactionDetailInquiryRequest2) validateTxnSuppFlg(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnSuppFlg) { // not required
		return nil
	}

	if err := validate.MinLength("txnSuppFlg", "body", string(*m.TxnSuppFlg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("txnSuppFlg", "body", string(*m.TxnSuppFlg), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionDetailInquiryRequest2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionDetailInquiryRequest2) UnmarshalBinary(b []byte) error {
	var res TransactionDetailInquiryRequest2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
