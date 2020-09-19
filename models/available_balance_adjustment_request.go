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

// AvailableBalanceAdjustmentRequest available balance adjustment request
//
// swagger:model AvailableBalanceAdjustmentRequest
type AvailableBalanceAdjustmentRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero.
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr,omitempty"`

	//  Max length = 3, Organization Number: Three digit Identification number of the organization.  Valid values are 001 - 998. Organization number must be on file.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	AcctOrg *string `json:"acctOrg,omitempty"`

	//  Max length = 1, Action flag indicating whether the transaction is memo credit or debit The values are: C = Credit D = Debit.
	// Max Length: 1
	// Min Length: 0
	Action *string `json:"action,omitempty"`

	//  Max length = 6, Authorization code
	// Max Length: 6
	// Min Length: 0
	AuthCde *string `json:"authCde,omitempty"`

	//  Max length = 4, Card Sequence Number: Identification number assigned to Embosser record to distinguish between multiple cards issued with the same card. 1. Must be numeric and greater than 0 if provided 2. Must be between 0 and 99 if smart card 3. If not provided, and not a smart card, default to value of 0001
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSeq *string `json:"cardSeq,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 5, Credit Plan Number: Identifies the plan number of the Credit Plan Master record associated with the Credit Plan Segment record.
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CrdPlan *string `json:"crdPlan,omitempty"`

	// Format: YYYYMMDD. Effective Date. Date format is CCYYMMDD format.
	EffDate string `json:"effDate,omitempty"`

	//  Max length = 9, Prepaid card fee amount
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	FeeAmt *string `json:"feeAmt,omitempty"`

	//  Max length = 1, Prepaid card fee calculation type
	// Max Length: 1
	// Min Length: 0
	FeeCalcType *string `json:"feeCalcType,omitempty"`

	//  Max length = 15, Prepaid card fee description
	// Max Length: 15
	// Min Length: 0
	FeeDesc *string `json:"feeDesc,omitempty"`

	//  Max length = 2, Prepaid card fee occurrence number
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	FeeOccur *string `json:"feeOccur,omitempty"`

	//  Max length = 1, Prepaid card fee waive reason
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	FeeWaiveRsn *string `json:"feeWaiveRsn,omitempty"`

	//  Max length = 2, Insurance product
	// Max Length: 2
	// Min Length: 0
	InsPrde *string `json:"insPrde,omitempty"`

	//  Max length = 5, Merchant category code
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	Mcc *string `json:"mcc,omitempty"`

	//  Max length = 11, Stock Keeping Unit number associated with a particular transaction.
	// Max Length: 11
	// Min Length: 0
	// Pattern: ^[0-9]*$
	SkuNbr *string `json:"skuNbr,omitempty"`

	//  Max length = 1, Source of the outstanding record
	// Max Length: 1
	// Min Length: 0
	Source *string `json:"source,omitempty"`

	//  Max length = 9, Merchant Store Number: Identification number of the Mechant Store. This field is used when the OFXRSI-SERVICE-FLAG is set to 'A' i.e. if Reversal is being triggered for a transaction.
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	StoreNbr *string `json:"storeNbr,omitempty"`

	//  Max length = 3, Merchant Organization Number: Three digit Identification number of the organization. Valid values are 001 - 998. This field is used when the OFXRSI-SERVICE-FLAG is set to 'A' i.e. if Reversal is being triggered for a transaction.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	StoreOrg *string `json:"storeOrg,omitempty"`

	//  Max length = 17, Transaction amount. without the commas and decimal places.    Example: if an account has available credit of EUR52.50 and the system's NOD set at 2, this field will be entered as 00000000000005250.
	// Max Length: 17
	// Min Length: 0
	// Pattern: ^[0-9]*$
	TranAmt *string `json:"tranAmt,omitempty"`

	//  Max length = 1, Prepaid card transaction type indicator
	// Max Length: 1
	// Min Length: 0
	TranTypeInd *string `json:"tranTypeInd,omitempty"`

	//  Max length = 5, Transaction code
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	TrnCde *string `json:"trnCde,omitempty"`
}

// Validate validates this available balance adjustment request
func (m *AvailableBalanceAdjustmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcctOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthCde(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrdPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeCalcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeOccur(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeWaiveRsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInsPrde(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkuNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranTypeInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrnCde(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateAcctNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctNbr) { // not required
		return nil
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateAcctOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctOrg) { // not required
		return nil
	}

	if err := validate.MinLength("acctOrg", "body", string(*m.AcctOrg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctOrg", "body", string(*m.AcctOrg), 3); err != nil {
		return err
	}

	if err := validate.Pattern("acctOrg", "body", string(*m.AcctOrg), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := validate.MinLength("action", "body", string(*m.Action), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("action", "body", string(*m.Action), 1); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateAuthCde(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthCde) { // not required
		return nil
	}

	if err := validate.MinLength("authCde", "body", string(*m.AuthCde), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("authCde", "body", string(*m.AuthCde), 6); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateCardSeq(formats strfmt.Registry) error {

	if swag.IsZero(m.CardSeq) { // not required
		return nil
	}

	if err := validate.MinLength("cardSeq", "body", string(*m.CardSeq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardSeq", "body", string(*m.CardSeq), 4); err != nil {
		return err
	}

	if err := validate.Pattern("cardSeq", "body", string(*m.CardSeq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *AvailableBalanceAdjustmentRequest) validateCrdPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.CrdPlan) { // not required
		return nil
	}

	if err := validate.MinLength("crdPlan", "body", string(*m.CrdPlan), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("crdPlan", "body", string(*m.CrdPlan), 5); err != nil {
		return err
	}

	if err := validate.Pattern("crdPlan", "body", string(*m.CrdPlan), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateFeeAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeAmt) { // not required
		return nil
	}

	if err := validate.MinLength("feeAmt", "body", string(*m.FeeAmt), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("feeAmt", "body", string(*m.FeeAmt), 9); err != nil {
		return err
	}

	if err := validate.Pattern("feeAmt", "body", string(*m.FeeAmt), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateFeeCalcType(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeCalcType) { // not required
		return nil
	}

	if err := validate.MinLength("feeCalcType", "body", string(*m.FeeCalcType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("feeCalcType", "body", string(*m.FeeCalcType), 1); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateFeeDesc(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeDesc) { // not required
		return nil
	}

	if err := validate.MinLength("feeDesc", "body", string(*m.FeeDesc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("feeDesc", "body", string(*m.FeeDesc), 15); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateFeeOccur(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeOccur) { // not required
		return nil
	}

	if err := validate.MinLength("feeOccur", "body", string(*m.FeeOccur), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("feeOccur", "body", string(*m.FeeOccur), 2); err != nil {
		return err
	}

	if err := validate.Pattern("feeOccur", "body", string(*m.FeeOccur), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateFeeWaiveRsn(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeWaiveRsn) { // not required
		return nil
	}

	if err := validate.MinLength("feeWaiveRsn", "body", string(*m.FeeWaiveRsn), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("feeWaiveRsn", "body", string(*m.FeeWaiveRsn), 1); err != nil {
		return err
	}

	if err := validate.Pattern("feeWaiveRsn", "body", string(*m.FeeWaiveRsn), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateInsPrde(formats strfmt.Registry) error {

	if swag.IsZero(m.InsPrde) { // not required
		return nil
	}

	if err := validate.MinLength("insPrde", "body", string(*m.InsPrde), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("insPrde", "body", string(*m.InsPrde), 2); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateMcc(formats strfmt.Registry) error {

	if swag.IsZero(m.Mcc) { // not required
		return nil
	}

	if err := validate.MinLength("mcc", "body", string(*m.Mcc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("mcc", "body", string(*m.Mcc), 5); err != nil {
		return err
	}

	if err := validate.Pattern("mcc", "body", string(*m.Mcc), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateSkuNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.SkuNbr) { // not required
		return nil
	}

	if err := validate.MinLength("skuNbr", "body", string(*m.SkuNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("skuNbr", "body", string(*m.SkuNbr), 11); err != nil {
		return err
	}

	if err := validate.Pattern("skuNbr", "body", string(*m.SkuNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if err := validate.MinLength("source", "body", string(*m.Source), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("source", "body", string(*m.Source), 1); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateStoreNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.StoreNbr) { // not required
		return nil
	}

	if err := validate.MinLength("storeNbr", "body", string(*m.StoreNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("storeNbr", "body", string(*m.StoreNbr), 9); err != nil {
		return err
	}

	if err := validate.Pattern("storeNbr", "body", string(*m.StoreNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateStoreOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.StoreOrg) { // not required
		return nil
	}

	if err := validate.MinLength("storeOrg", "body", string(*m.StoreOrg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("storeOrg", "body", string(*m.StoreOrg), 3); err != nil {
		return err
	}

	if err := validate.Pattern("storeOrg", "body", string(*m.StoreOrg), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateTranAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.TranAmt) { // not required
		return nil
	}

	if err := validate.MinLength("tranAmt", "body", string(*m.TranAmt), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("tranAmt", "body", string(*m.TranAmt), 17); err != nil {
		return err
	}

	if err := validate.Pattern("tranAmt", "body", string(*m.TranAmt), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateTranTypeInd(formats strfmt.Registry) error {

	if swag.IsZero(m.TranTypeInd) { // not required
		return nil
	}

	if err := validate.MinLength("tranTypeInd", "body", string(*m.TranTypeInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("tranTypeInd", "body", string(*m.TranTypeInd), 1); err != nil {
		return err
	}

	return nil
}

func (m *AvailableBalanceAdjustmentRequest) validateTrnCde(formats strfmt.Registry) error {

	if swag.IsZero(m.TrnCde) { // not required
		return nil
	}

	if err := validate.MinLength("trnCde", "body", string(*m.TrnCde), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("trnCde", "body", string(*m.TrnCde), 5); err != nil {
		return err
	}

	if err := validate.Pattern("trnCde", "body", string(*m.TrnCde), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailableBalanceAdjustmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableBalanceAdjustmentRequest) UnmarshalBinary(b []byte) error {
	var res AvailableBalanceAdjustmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
