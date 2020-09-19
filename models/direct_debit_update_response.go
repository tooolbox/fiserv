// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectDebitUpdateResponse direct debit update response
//
// swagger:model DirectDebitUpdateResponse
type DirectDebitUpdateResponse struct {

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	AcctNbr string `json:"acctNbr,omitempty"`

	//  Max length = 1, This field identifies the current payment  method of the customer's DD Setup (This is the updated payment method after update). 1      Minimum Due 2Nominated Amount 3Nominated Percentage 4Full Balance
	AchPmtOptn string `json:"achPmtOptn,omitempty"`

	//  Max length = 1, This field identifies previous payment method of the customer's DD Setup (This is the payment method  immediately prior to the update request). 1      Minimum Due 2Nominated Amount 3Nominated Percentage 4Full Balance
	AchPmtOptnPrev string `json:"achPmtOptnPrev,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 3, Currency Code at Org Level: ISO currency code that identifies the unit of currency for this account.
	CurrencyCode string `json:"currencyCode,omitempty"`

	//  Max length = 1, Currency NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in monetary amount fields.
	CurrencyNod string `json:"currencyNod,omitempty"`

	// Format: YYYYMMDD. Date on which direct debit processing expires for the account. A value of zeros means the DD does not expire.
	DateAchPmtExpr string `json:"dateAchPmtExpr,omitempty"`

	// Format: YYYYMMDD. Account's Last Cycle Date.
	DateLastCycle string `json:"dateLastCycle,omitempty"`

	// Format: YYYYMMDD. Account's Next Cycle Date.
	DateNextCycle string `json:"dateNextCycle,omitempty"`

	// Format: YYYYMMDD. Payment due date: Date when the next payment is due.
	DatePmtDue string `json:"datePmtDue,omitempty"`

	//  Max length = 17, This field identifies the current Direct Debit Account Number after the DD setup change
	DdAcctNbr string `json:"ddAcctNbr,omitempty"`

	//  Max length = 2, Direct Debit request days. Number of days prior to the payment due date or the day of the month to request direct debit payments for the account.
	DdAchReqDay string `json:"ddAchReqDay,omitempty"`

	//  Max length = 17, This field identifies last cycle actual Direct Debit payment
	DdActPmt string `json:"ddActPmt,omitempty"`

	//  Max length = 11, Bank identifier code (BIC) associated with the direct debit account. A SWIFT code is an international bank code that identifies particular banks worldwide.
	DdBic string `json:"ddBic,omitempty"`

	//  Max length = 1, Direct Debit Cancellation. Values are:  ' 0 '  - do not cancel DD ' 1 '  - Cancel DD
	DdCancel string `json:"ddCancel,omitempty"`

	//  Max length = 1, Direct debit credit balance. Values are: 0 = Direct credits are not generated for accounts with credit balances 1 = Direct credits are generated for accounts with   credit balances
	DdCreditBal string `json:"ddCreditBal,omitempty"`

	//  Max length = 1, Direct debit daily frequency. Values are: 0 = Account does not generate direct debits on a daily frequency 1 = Account generates direct debits daily 2 = Account generates direct debits every week 3 = Account generates direct debits biweekly or every 14 days
	DdDailyFreq string `json:"ddDailyFreq,omitempty"`

	//  Max length = 1, Flag indicating if the client is using the facility to request the direct debit payment to the account a number of days before the direct debit is due. Valid values are: 0-DD Days same as pmt due days 1-DD Days can be different to pmt due days  #.#
	DdDaysFctn string `json:"ddDaysFctn,omitempty"`

	//  Max length = 2, The number of days before the payment is due to post the direct debit payment to the account
	DdDaysNbr string `json:"ddDaysNbr,omitempty"`

	// Format: YYYYMMDD. This field identifies Direct Debit Effective Date <Date will be responded back as Gregorian date-CCYYMMDD format>
	DdEffDate string `json:"ddEffDate,omitempty"`

	//  Max length = 34, Required if the IBAN flag on ARML12 is equal to 1
	DdIban string `json:"ddIban,omitempty"`

	//  Max length = 1, Direct debit interim payments. Values are: 0 = Do not apply interim payments 1 = Apply interim payment to the calculated payment prior to generating a direct debit transaction
	DdInterimPmts string `json:"ddInterimPmts,omitempty"`

	//  Max length = 1, This field indicates the mandate status of the Direct Debit. Direct Debit will be generated only if this status is active and all other mandatory Direct DebitD parameters are set up. This field will be open to maintenance only if the MANDATE FLAG processing is set ON at the ORG level SPACES  =  DD Mandate is not set A =  DD Mandate is active  I=  DD Mandate is inactive C =  DD Mandate is cancelled  E=  DD Mandate is expired
	DdMandateStatus string `json:"ddMandateStatus,omitempty"`

	// Format: YYYYMMDD. Direct debit payment change date <Date will be accepted as Gregorian date-CCYYMMDD format>
	DdPmtChangeDate string `json:"ddPmtChangeDate,omitempty"`

	//  Max length = 1, Counter that CMS updates whenever a payment reversal is applied to the account using transaction logic module 031 or 032
	DdPmtRevCounter string `json:"ddPmtRevCounter,omitempty"`

	//  Max length = 1, Direct Debit Suspend    Valid values are  '0' - Inactive '1'- Active
	DdSuspend string `json:"ddSuspend,omitempty"`

	//  Max length = 1, Direct Debit Suspend  bypass  Valid values are  '0' -Do not bypass '1'-bypass
	DdSuspendBypass string `json:"ddSuspendBypass,omitempty"`

	//  Max length = 1, Foreign use indicator The values are: 0 = Local account 1 = Foreign account
	ForeignUse string `json:"foreignUse,omitempty"`

	//  Max length = 3, Logo
	Logo string `json:"logo,omitempty"`

	//  Max length = 17, Customer-nominated amount or percentage. Enter an amount in monetary units and subunits. For example, if working with U.S. dollars, enter $125.73 as 00000000000012573. Enter a percentage using only the last seven positions of the field. The location of the decimal is determined by the percentage NOD. For example, if the percentage NOD is 7, enter 15% as 00000000001500000.
	NewDdAchAmt string `json:"newDdAchAmt,omitempty"`

	//  Max length = 7, Customer-nominated amount or percentage. Enter an amount in monetary units and subunits. For example, if working with U.S. dollars, enter $125.73 as 00000000000012573. Enter a percentage using only the last seven positions of the field. The location of the decimal is determined by the percentage NOD. For example, if the percentage NOD is 7, enter 15% as 00000000001500000.
	NewDdAchPct string `json:"newDdAchPct,omitempty"`

	//  Max length = 1, Direct debit customer-nominated type payment 0 = Not used (where PMT ACH FLAG is 0 or 1) 1 = Fixed payment amount 2 = Percentage of account balance (cycle-ending) 3 = Full account balance (cycle-ending)
	NewNomAchAmtPctFlag string `json:"newNomAchAmtPctFlag,omitempty"`

	//  Max length = 1, Payment ACH flag: Values are: 0 = No ACH 1 = ACH active A minimum payment is in effect. This option includes past due amounts for the projected Direct Debit. 2 = Direct debit processing is active. A customer nominated payment is in effect. This option includes past due amounts for the projected DD. 3 = xx 4 = xx 5 = xx 6 = xx 7 = Direct debit processing is active. A customer nominated payment is in effect. This option does not include past due amounts for the projected DD.
	NewPmtAchFlag string `json:"newPmtAchFlag,omitempty"`

	// Format: YYYYMMDD. Customer-nominated payment date <Date will be accepted as Gregorian date-CCYYMMDD format>
	NomAchStartDate string `json:"nomAchStartDate,omitempty"`

	//  Max length = 1, Payment remittance method. Values are: 1 = Standing order 2 = Direct debit
	PayRemitMethod string `json:"payRemitMethod,omitempty"`

	//  Max length = 1, Percentage NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in percentage fields.
	PctRateNod string `json:"pctRateNod,omitempty"`

	//  Max length = 1, DD payment Type: Code that indicates the type of account to which the direct debit is sent.  Values are: D = Demand deposit account (checking) S = Savings account.
	PmtAchDbType string `json:"pmtAchDbType,omitempty"`

	//  Max length = 10, Payment Automated Clearing House routing number: Dual-use field that accepts a nine-digit routing/transit number (ACH format) or a ten-digit Bank ID (non-ACH format) to identify the financial institution to receive direct debit transactions for account payments.
	PmtAchRtNbr string `json:"pmtAchRtNbr,omitempty"`

	//  Max length = 17, This field identifies the previous DD Account Number prior to DD setup change prior to DD Amend
	PrevDdAcctNbr string `json:"prevDdAcctNbr,omitempty"`

	//  Max length = 17, This field identifies old DD ACH Amount prior to Direct Debit setup change
	PrevDdAchAmt string `json:"prevDdAchAmt,omitempty"`

	//  Max length = 7, This field identifies old DD ACH Percent prior to Diret Debit setup change
	PrevDdAchPct string `json:"prevDdAchPct,omitempty"`

	//  Max length = 10, This field identifies the current DD Account Sort Code prior to DD setup change prior to DD Amend
	PrevDdSortCode string `json:"prevDdSortCode,omitempty"`

	//  Max length = 1, This field identifies old DD NOM ACH Type prior to Direct Debit setup change
	PrevNomAchAmtPctFlag string `json:"prevNomAchAmtPctFlag,omitempty"`

	//  Max length = 1, Previous Payment remittance method The values are: 1 = Standing order 2 = Direct debit
	PrevPayRemitMethod string `json:"prevPayRemitMethod,omitempty"`

	//  Max length = 1, This field identifies old DD ACH Type prior to Direct Debit setup change
	PrevPmtAchFlag string `json:"prevPmtAchFlag,omitempty"`

	//  Max length = 17, Projected Automated Clearing House Payment
	ProjAchPmt string `json:"projAchPmt,omitempty"`

	//  Max length = 17, RECENT PROJ DD AMT
	RecentProjDdAmt string `json:"recentProjDdAmt,omitempty"`
}

// Validate validates this direct debit update response
func (m *DirectDebitUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitUpdateResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *DirectDebitUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitUpdateResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
