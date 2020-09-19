// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InsData2ForInsuranceInquiry2 ins data2 for insurance inquiry2
//
// swagger:model InsData2ForInsuranceInquiry2
type InsData2ForInsuranceInquiry2 struct {

	//  Max length = 1, Insurance(INS) Cancelation reason Blank = Not cancelled (Default) A= Reached expiration age B=Due to block code(BC) C=Charged off D=past due exceeds number of cycles F=fiche nr required but not assigned. Nr of cancellation days reached G=date of birth(DOB) not meeting enrol. age req. I=DOB not meeting min. age req. N=loan settlement P=credit line or account closed R=request of cust. T=max. AMT at risk exceeds max. coverage U=at request of underwriter X=other J=INS table not found M=BC cancels IN
	CanRsn string `json:"canRsn,omitempty"`

	//  Max length = 10, Insurance certificate number for credit insurance product.
	CertNbr string `json:"certNbr,omitempty"`

	//  Max length = 2, Insurance Company: User-defined code that identifies the underwriter of the insurance product.
	Company string `json:"company,omitempty"`

	//  Max length = 17, Cycle to date premium billed.
	CtdPremiumBilled string `json:"ctdPremiumBilled,omitempty"`

	//  Max length = 18, Balance as of the current claim date.
	CurClmBal string `json:"curClmBal,omitempty"`

	// Format: YYYYMMDD. Ending date on which the new insurance claim will expire.
	CurClmEddt string `json:"curClmEddt,omitempty"`

	//  Max length = 2, Current Insurance Claim Reason Code: Reason why current insurance claim was initiated.
	CurClmRsn string `json:"curClmRsn,omitempty"`

	//  Max length = 1, Status of Current Insurance Claim status.
	CurClmStat string `json:"curClmStat,omitempty"`

	// Format: YYYYMMDD. Current Insurance Claim Status Date: Effective date for the new insurance claim.
	CurClmStdt string `json:"curClmStdt,omitempty"`

	//  Max length = 20, Current Insurance Claim Reason Description: Description associated with the claim reason code. This description is available for reports, letters, and statement messages.
	CurRsnDesc string `json:"curRsnDesc,omitempty"`

	// Format: YYYYMMDD. Last date insurance was billed.  Date format is CCYY-MM-DD.
	DtLstBilled string `json:"dtLstBilled,omitempty"`

	// Format: YYYYMMDD. Insurance effective date.  Date format is CCYY-MM-DD.
	EffDt string `json:"effDt,omitempty"`

	//  Max length = 3, Insurance enrollment ID of the account holder: enrollment ID  which determines the Processing Control Table and the Insurance table for an account.
	EnrollID string `json:"enrollId,omitempty"`

	// Format: YYYYMMDD. Free Look Period End Date: Date on which the free look period ends for the insurance product associated with the account.
	FlpEndDate string `json:"flpEndDate,omitempty"`

	//  Max length = 1, System generated code that indicates if the insurance product is active for this account 0 = Insurance product is not active. (Default) 1 = Insurance product is active but CMS has not generated a statement for the account since the product became active. 2 = Insurance product is active and a statement for the account has been produced since the insurance product became active. CMS sets this value after generating the first statement with this insurance product.
	InsActive string `json:"insActive,omitempty"`

	//  Max length = 1, Insurance Billing Frequency: Code that indicates the frequency at which CMS assesses insurance premiums. Values are: 0 = Monthly on the cycle date (Default) 1 = Bimonthly 2 = Quarterly 3 = Semi-annually 4 = Annually.
	InsBlngFreq string `json:"insBlngFreq,omitempty"`

	// Format: YYYYMMDD. Insurance Cancellation Date: Date on which this insurance policy was cancelled, either manually or automatically.  Date format is CCYY-MM-DD.
	InsCanDt string `json:"insCanDt,omitempty"`

	//  Max length = 1, Insurance Channel: method of insurance application. Values are: P - Postal  S - Store  I - Internet T - Telephone
	InsChnl string `json:"insChnl,omitempty"`

	//  Max length = 9, Insurance group policy number: Group Policy Number for a Credit or free insurance product.
	InsChnlPlyNbr string `json:"insChnlPlyNbr,omitempty"`

	//  Max length = 19, Insurance Customer Number: Number of the Customer Name/Address record that identifies the insured customer for the insurance product.
	InsCustNbr string `json:"insCustNbr,omitempty"`

	// Format: YYYYMMDD. Insurance expiration date
	InsDateExp string `json:"insDateExp,omitempty"`

	// Format: YYYYMMDD. Date of last insurance claim.  Date format is CCYY-MM-DD.
	InsDtLstClaim string `json:"insDtLstClaim,omitempty"`

	//  Max length = 17, Fixed Insurance Premium: Amount that establishes a premium rate in addition to the PER-VALUE rate.
	InsFixPremium string `json:"insFixPremium,omitempty"`

	//  Max length = 18, Insurance Premium.
	InsInsPrem string `json:"insInsPrem,omitempty"`

	//  Max length = 40, Insured customer name.
	InsName string `json:"insName,omitempty"`

	//  Max length = 2, Insurance Per-Value Age: Maximum age of the insured party at which CMS will assess an insurance premium based on the percentage or amount in the corresponding (PER VALUE) field.
	InsPerValAge string `json:"insPerValAge,omitempty"`

	//  Max length = 17, Insurance Per-Value Amount: Amount assessed for insurance premiums based on the age of the insured party.
	InsPerValAmt string `json:"insPerValAmt,omitempty"`

	//  Max length = 1, Insurance Per-Value flag: Code that indicates whether PER-VALUE is a percentage or dollar amount.0 = Percentage value1 = Amount value
	InsPerValFlg string `json:"insPerValFlg,omitempty"`

	//  Max length = 7, Insurance Per-Value Percentage: Percentage of the account balance for insurance premiums based on the age of the insured party.
	InsPerValPercent string `json:"insPerValPercent,omitempty"`

	//  Max length = 2, Insurance Product: Insurance product code as defined on the Insurance table. An insurance product cannot be added more than once to an account at any one time. Similarly an insurance product cannot be re-added after initialization until the associated insurance history record has been removed from the Insurance History file.
	InsPrd string `json:"insPrd,omitempty"`

	//  Max length = 40, Insurance product description as defined on the Insurance table.
	InsPrdDesc string `json:"insPrdDesc,omitempty"`

	// Format: YYYYMMDD. Insurance Reinstate Date: Date on which the insurance product was reinstated, either manually or automatically.  Date format is CCYY-MM-DD.
	InsReinstateDt string `json:"insReinstateDt,omitempty"`

	//  Max length = 2, Insurance source code
	InsSourceCode string `json:"insSourceCode,omitempty"`

	//  Max length = 1, Status of Insurance.  Values are: C = Insurance product is cancelled. F = Insurance product is in force.  I = Insurance product is inactive due to a user defined condition.  S = Insurance product is suspended due to delinquency. P = Pending cancelled
	InsStatCd string `json:"insStatCd,omitempty"`

	// Format: YYYYMMDD. Insurance Status Change Date: Date when the insurance status was last changed.  Date format is CCYY-MM-DD.
	InsStatChgDate string `json:"insStatChgDate,omitempty"`

	//  Max length = 3, Insurance Table Number: Number that identifies the Insurance table to which this insurance product is assigned.
	InsTblNbr string `json:"insTblNbr,omitempty"`

	//  Max length = 1, Insurance Type: Code that indicates whether a product is an insurance product or non-insurance (additional account) product. Values are: 0 = Insurance product (Default) 1 = Non-insurance or additional account product
	InsTyp string `json:"insTyp,omitempty"`

	//  Max length = 17, Last insurance premium billed.
	LstPremiumBilled string `json:"lstPremiumBilled,omitempty"`

	//  Max length = 17, Life to date premium billed.
	LtdPremiumBilled string `json:"ltdPremiumBilled,omitempty"`

	//  Max length = 17, Month to date premium billed.
	MtdPremiumBilled string `json:"mtdPremiumBilled,omitempty"`

	//  Max length = 18, Insurance policy reference number.
	PlcChnl string `json:"plcChnl,omitempty"`

	//  Max length = 18, Balance as of the previous claim date.
	PreClmBal string `json:"preClmBal,omitempty"`

	// Format: YYYYMMDD. Ending date on which the previous insurance claim will expire.
	PreClmEddt string `json:"preClmEddt,omitempty"`

	//  Max length = 2, Previous Insurance claim reason: reason code that indicates the status of the previous insurance claim.
	PreClmRsn string `json:"preClmRsn,omitempty"`

	//  Max length = 1, Previous Insurance claim status: status of the previous insurance claim.
	PreClmStat string `json:"preClmStat,omitempty"`

	// Format: YYYYMMDD. Previous Insurance Claim Status Date: Effective date for the previous insurance claim.
	PreClmStdt string `json:"preClmStdt,omitempty"`

	//  Max length = 20, Previous Insurance Claim reason Description: Description associated with the claim reason code. This description is available for reports, letters, and statement messages.
	PreRsnDesc string `json:"preRsnDesc,omitempty"`

	//  Max length = 9, Insurance premium.
	PremiumRate string `json:"premiumRate,omitempty"`

	// Format: YYYYMMDD. Insurance sales date: date when the customer has contracted the insurance.
	SlsDate string `json:"slsDate,omitempty"`

	//  Max length = 17, year-to-date premium billed.
	YtdPremiumBilled string `json:"ytdPremiumBilled,omitempty"`
}

// Validate validates this ins data2 for insurance inquiry2
func (m *InsData2ForInsuranceInquiry2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InsData2ForInsuranceInquiry2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InsData2ForInsuranceInquiry2) UnmarshalBinary(b []byte) error {
	var res InsData2ForInsuranceInquiry2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
