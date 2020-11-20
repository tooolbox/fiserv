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

// MonetaryActionRequest monetary action request
//
// swagger:model MonetaryActionRequest
type MonetaryActionRequest struct {

	//  Max length = 19, ACCT NBR
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr,omitempty"`

	//  Max length = 3, ACCT ORG
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	AcctOrg *string `json:"acctOrg,omitempty"`

	//  Max length = 4, ACTION CODE
	// Max Length: 4
	// Min Length: 0
	ActionCode *string `json:"actionCode,omitempty"`

	//  Max length = 3, ASM ORG
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	AsmOrg *string `json:"asmOrg,omitempty"`

	//  Max length = 3, ASM REP
	// Max Length: 3
	// Min Length: 0
	AsmRep *string `json:"asmRep,omitempty"`

	//  Max length = 6, AUTH CODE
	// Max Length: 6
	// Min Length: 0
	AuthCode *string `json:"authCode,omitempty"`

	//  Max length = 19, CARD NBR
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr,omitempty"`

	//  Max length = 4, CARD SEQ
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSeq *string `json:"cardSeq,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 3, CURR CODE
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CurrCode *string `json:"currCode,omitempty"`

	//  Max length = 4, DEPT CODE
	// Max Length: 4
	// Min Length: 0
	DeptCode *string `json:"deptCode,omitempty"`

	// Format: YYYYMMDD. EFF DATE
	EffDate string `json:"effDate,omitempty"`

	//  Max length = 1, FOREIGN USE
	// Max Length: 1
	// Min Length: 0
	ForeignUse *string `json:"foreignUse,omitempty"`

	//  Max length = 2, INSURANCE CODE
	// Max Length: 2
	// Min Length: 0
	InsuranceCode *string `json:"insuranceCode,omitempty"`

	//  Max length = 60, LINE 1
	// Max Length: 60
	// Min Length: 0
	Line1 *string `json:"line1,omitempty"`

	//  Max length = 60, LINE 2
	// Max Length: 60
	// Min Length: 0
	Line2 *string `json:"line2,omitempty"`

	//  Max length = 60, LINE 3
	// Max Length: 60
	// Min Length: 0
	Line3 *string `json:"line3,omitempty"`

	//  Max length = 60, LINE 4
	// Max Length: 60
	// Min Length: 0
	Line4 *string `json:"line4,omitempty"`

	//  Max length = 60, LINE 5
	// Max Length: 60
	// Min Length: 0
	Line5 *string `json:"line5,omitempty"`

	//  Max length = 3, LTR CODE
	// Max Length: 3
	// Min Length: 0
	LtrCode *string `json:"ltrCode,omitempty"`

	//  Max length = 3, LTR ORG
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	LtrOrg *string `json:"ltrOrg,omitempty"`

	//  Max length = 4, MCC
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	Mcc *string `json:"mcc,omitempty"`

	//  Max length = 5, PLAN NBR
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	PlanNbr *string `json:"planNbr,omitempty"`

	//  Max length = 2, PLAN SEQ
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	PlanSeq *string `json:"planSeq,omitempty"`

	//  Max length = 1, Payment Notify Alert Flag Valid values are: <UL><LI> <EM>Y -</EM> Yes </LI> <LI><EM>N - </EM> No</LI></UL>
	// Max Length: 1
	// Min Length: 0
	PmtNtfyAlertSrc *string `json:"pmtNtfyAlertSrc,omitempty"`

	//  Max length = 16, PURCHASE ORDER NBR
	// Max Length: 16
	// Min Length: 0
	PurchaseOrderNbr *string `json:"purchaseOrderNbr,omitempty"`

	//  Max length = 14, REFERENCE NBR
	// Max Length: 14
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ReferenceNbr *string `json:"referenceNbr,omitempty"`

	//  Max length = 12, SALES CLERK
	// Max Length: 12
	// Min Length: 0
	SalesClerk *string `json:"salesClerk,omitempty"`

	//  Max length = 9, SKU NBR
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	SkuNbr *string `json:"skuNbr,omitempty"`

	//  Max length = 9, STORE NBR
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	StoreNbr *string `json:"storeNbr,omitempty"`

	//  Max length = 15, TICKET NBR
	// Max Length: 15
	// Min Length: 0
	TicketNbr *string `json:"ticketNbr,omitempty"`

	//  Max length = 40, TRAN DESC
	// Max Length: 40
	// Min Length: 0
	TranDesc *string `json:"tranDesc,omitempty"`

	//  Max length = 11, TXN AMOUNT
	// Max Length: 11
	// Min Length: 0
	// Pattern: ^[0-9]*$
	TxnAmount *string `json:"txnAmount,omitempty"`
}

// Validate validates this monetary action request
func (m *MonetaryActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcctOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsmOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsmRep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeptCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForeignUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInsuranceCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLtrCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLtrOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmtNtfyAlertSrc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalesClerk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkuNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxnAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonetaryActionRequest) validateAcctNbr(formats strfmt.Registry) error {

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

func (m *MonetaryActionRequest) validateAcctOrg(formats strfmt.Registry) error {

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

func (m *MonetaryActionRequest) validateActionCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionCode) { // not required
		return nil
	}

	if err := validate.MinLength("actionCode", "body", string(*m.ActionCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("actionCode", "body", string(*m.ActionCode), 4); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateAsmOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.AsmOrg) { // not required
		return nil
	}

	if err := validate.MinLength("asmOrg", "body", string(*m.AsmOrg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("asmOrg", "body", string(*m.AsmOrg), 3); err != nil {
		return err
	}

	if err := validate.Pattern("asmOrg", "body", string(*m.AsmOrg), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateAsmRep(formats strfmt.Registry) error {

	if swag.IsZero(m.AsmRep) { // not required
		return nil
	}

	if err := validate.MinLength("asmRep", "body", string(*m.AsmRep), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("asmRep", "body", string(*m.AsmRep), 3); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateAuthCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthCode) { // not required
		return nil
	}

	if err := validate.MinLength("authCode", "body", string(*m.AuthCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("authCode", "body", string(*m.AuthCode), 6); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateCardNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CardNbr) { // not required
		return nil
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateCardSeq(formats strfmt.Registry) error {

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

func (m *MonetaryActionRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *MonetaryActionRequest) validateCurrCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrCode) { // not required
		return nil
	}

	if err := validate.MinLength("currCode", "body", string(*m.CurrCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("currCode", "body", string(*m.CurrCode), 3); err != nil {
		return err
	}

	if err := validate.Pattern("currCode", "body", string(*m.CurrCode), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateDeptCode(formats strfmt.Registry) error {

	if swag.IsZero(m.DeptCode) { // not required
		return nil
	}

	if err := validate.MinLength("deptCode", "body", string(*m.DeptCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("deptCode", "body", string(*m.DeptCode), 4); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateForeignUse(formats strfmt.Registry) error {

	if swag.IsZero(m.ForeignUse) { // not required
		return nil
	}

	if err := validate.MinLength("foreignUse", "body", string(*m.ForeignUse), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("foreignUse", "body", string(*m.ForeignUse), 1); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateInsuranceCode(formats strfmt.Registry) error {

	if swag.IsZero(m.InsuranceCode) { // not required
		return nil
	}

	if err := validate.MinLength("insuranceCode", "body", string(*m.InsuranceCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("insuranceCode", "body", string(*m.InsuranceCode), 2); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLine1(formats strfmt.Registry) error {

	if swag.IsZero(m.Line1) { // not required
		return nil
	}

	if err := validate.MinLength("line1", "body", string(*m.Line1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line1", "body", string(*m.Line1), 60); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLine2(formats strfmt.Registry) error {

	if swag.IsZero(m.Line2) { // not required
		return nil
	}

	if err := validate.MinLength("line2", "body", string(*m.Line2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line2", "body", string(*m.Line2), 60); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLine3(formats strfmt.Registry) error {

	if swag.IsZero(m.Line3) { // not required
		return nil
	}

	if err := validate.MinLength("line3", "body", string(*m.Line3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line3", "body", string(*m.Line3), 60); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLine4(formats strfmt.Registry) error {

	if swag.IsZero(m.Line4) { // not required
		return nil
	}

	if err := validate.MinLength("line4", "body", string(*m.Line4), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line4", "body", string(*m.Line4), 60); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLine5(formats strfmt.Registry) error {

	if swag.IsZero(m.Line5) { // not required
		return nil
	}

	if err := validate.MinLength("line5", "body", string(*m.Line5), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line5", "body", string(*m.Line5), 60); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLtrCode(formats strfmt.Registry) error {

	if swag.IsZero(m.LtrCode) { // not required
		return nil
	}

	if err := validate.MinLength("ltrCode", "body", string(*m.LtrCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("ltrCode", "body", string(*m.LtrCode), 3); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateLtrOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.LtrOrg) { // not required
		return nil
	}

	if err := validate.MinLength("ltrOrg", "body", string(*m.LtrOrg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("ltrOrg", "body", string(*m.LtrOrg), 3); err != nil {
		return err
	}

	if err := validate.Pattern("ltrOrg", "body", string(*m.LtrOrg), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateMcc(formats strfmt.Registry) error {

	if swag.IsZero(m.Mcc) { // not required
		return nil
	}

	if err := validate.MinLength("mcc", "body", string(*m.Mcc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("mcc", "body", string(*m.Mcc), 4); err != nil {
		return err
	}

	if err := validate.Pattern("mcc", "body", string(*m.Mcc), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validatePlanNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanNbr) { // not required
		return nil
	}

	if err := validate.MinLength("planNbr", "body", string(*m.PlanNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("planNbr", "body", string(*m.PlanNbr), 5); err != nil {
		return err
	}

	if err := validate.Pattern("planNbr", "body", string(*m.PlanNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validatePlanSeq(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanSeq) { // not required
		return nil
	}

	if err := validate.MinLength("planSeq", "body", string(*m.PlanSeq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("planSeq", "body", string(*m.PlanSeq), 2); err != nil {
		return err
	}

	if err := validate.Pattern("planSeq", "body", string(*m.PlanSeq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validatePmtNtfyAlertSrc(formats strfmt.Registry) error {

	if swag.IsZero(m.PmtNtfyAlertSrc) { // not required
		return nil
	}

	if err := validate.MinLength("pmtNtfyAlertSrc", "body", string(*m.PmtNtfyAlertSrc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("pmtNtfyAlertSrc", "body", string(*m.PmtNtfyAlertSrc), 1); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validatePurchaseOrderNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.PurchaseOrderNbr) { // not required
		return nil
	}

	if err := validate.MinLength("purchaseOrderNbr", "body", string(*m.PurchaseOrderNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("purchaseOrderNbr", "body", string(*m.PurchaseOrderNbr), 16); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateReferenceNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceNbr) { // not required
		return nil
	}

	if err := validate.MinLength("referenceNbr", "body", string(*m.ReferenceNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("referenceNbr", "body", string(*m.ReferenceNbr), 14); err != nil {
		return err
	}

	if err := validate.Pattern("referenceNbr", "body", string(*m.ReferenceNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateSalesClerk(formats strfmt.Registry) error {

	if swag.IsZero(m.SalesClerk) { // not required
		return nil
	}

	if err := validate.MinLength("salesClerk", "body", string(*m.SalesClerk), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("salesClerk", "body", string(*m.SalesClerk), 12); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateSkuNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.SkuNbr) { // not required
		return nil
	}

	if err := validate.MinLength("skuNbr", "body", string(*m.SkuNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("skuNbr", "body", string(*m.SkuNbr), 9); err != nil {
		return err
	}

	if err := validate.Pattern("skuNbr", "body", string(*m.SkuNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateStoreNbr(formats strfmt.Registry) error {

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

func (m *MonetaryActionRequest) validateTicketNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.TicketNbr) { // not required
		return nil
	}

	if err := validate.MinLength("ticketNbr", "body", string(*m.TicketNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("ticketNbr", "body", string(*m.TicketNbr), 15); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateTranDesc(formats strfmt.Registry) error {

	if swag.IsZero(m.TranDesc) { // not required
		return nil
	}

	if err := validate.MinLength("tranDesc", "body", string(*m.TranDesc), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("tranDesc", "body", string(*m.TranDesc), 40); err != nil {
		return err
	}

	return nil
}

func (m *MonetaryActionRequest) validateTxnAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.TxnAmount) { // not required
		return nil
	}

	if err := validate.MinLength("txnAmount", "body", string(*m.TxnAmount), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("txnAmount", "body", string(*m.TxnAmount), 11); err != nil {
		return err
	}

	if err := validate.Pattern("txnAmount", "body", string(*m.TxnAmount), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MonetaryActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonetaryActionRequest) UnmarshalBinary(b []byte) error {
	var res MonetaryActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}