// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TxnDataForTransactionDetailInquiry1 txn data for transaction detail inquiry1
//
// swagger:model TxnDataForTransactionDetailInquiry1
type TxnDataForTransactionDetailInquiry1 struct {

	//  Max length = 19, Account Number: Identification Number of Customer's account.
	Account string `json:"account,omitempty"`

	//  Max length = 17, Amount of the transaction
	Amt string `json:"amt,omitempty"`

	//  Max length = 6, Authorization code assigned to the transaction.
	AuthCd string `json:"authCd,omitempty"`

	// Format: YYYYMMDD. Batch date of the transaction.  Date format is CCYY-MM-DD.
	BatchDt string `json:"batchDt,omitempty"`

	//  Max length = 5, Transaction category code: ISO standard category code of the merchant where frequent shopper points were awarded.
	CategoryCd string `json:"categoryCd,omitempty"`

	//  Max length = 16, Card acceptor identification: Identification number associated with a recurring payment transaction record
	CrdAcceptorID string `json:"crdAcceptorId,omitempty"`

	//  Max length = 1, Card Block Code: Block code assigned to the card number
	CrdBlkCd string `json:"crdBlkCd,omitempty"`

	//  Max length = 19, Card number: Number embossed on the card that initiated the transaction
	CrdNbr string `json:"crdNbr,omitempty"`

	//  Max length = 5, Card Sequence: Sequence number of the card that initiated the transaction
	CrdSeq string `json:"crdSeq,omitempty"`

	//  Max length = 1, Direct Debit or Savings Account Indicator:  Flag that indicates which bank account, DDA or savings account applies to the transaction.  The  field displays for debit card transactions only and is applicable only when Single Message Support is installed and activated Valid values : 'D', 'S'
	DdaSavInd string `json:"ddaSavInd,omitempty"`

	//  Max length = 4, Department number associated with a particular transaction
	Dept string `json:"dept,omitempty"`

	//  Max length = 40, Transaction Description: Description of the item for statements
	Desc string `json:"desc,omitempty"`

	// Format: YYYYMMDD. Effective date of the transaction.  Date format is CCYY-MM-DD.
	EffDt string `json:"effDt,omitempty"`

	//  Max length = 3, ISO Foreign currency code: Valid only for transactions originated in foreign currencies; Derived from the MEMO description  in a foreign transaction
	FgnCurrCd string `json:"fgnCurrCd,omitempty"`

	//  Max length = 17, Foreign Transaction Amount: Original Transaction  Amount in foreign currency. Valid only for transactions originated in foreign countries Derived from the MEMO description in a foreign transaction
	FgnTxnAmt string `json:"fgnTxnAmt,omitempty"`

	//  Max length = 14, Foreign Exchange Rate: The exchange rate used to convert between the original transaction currency and the cardholder billing currency.
	FgnXchgRt string `json:"fgnXchgRt,omitempty"`

	//  Max length = 2, Insurance Product: Flag that indicates insurance product for which insurance premium is generated.
	InsProd string `json:"insProd,omitempty"`

	//  Max length = 9, Transaction interchange fee.
	InterchgFee string `json:"interchgFee,omitempty"`

	//  Max length = 5, Item Number: System-generated number assigned to all debit transactions
	ItmNbr string `json:"itmNbr,omitempty"`

	//  Max length = 3, Identification number of the logo
	Logo string `json:"logo,omitempty"`

	//  Max length = 3, Merchant country code
	MerchCountryCd string `json:"merchCountryCd,omitempty"`

	//  Max length = 17, Original payment amount
	OrigPmtAmt string `json:"origPmtAmt,omitempty"`

	//  Max length = 1, Frequent shopper program number
	PgmNbr string `json:"pgmNbr,omitempty"`

	//  Max length = 5, Plan Number: Number that identifies the Credit Plan Master record associated with the credit plan to which the transaction was posted
	Plan string `json:"plan,omitempty"`

	// Format: YYYYMMDD. Plan Open Date: Date plan segment was opened.  Date format is CCYY-MM-DD.
	PlanOpenDt string `json:"planOpenDt,omitempty"`

	//  Max length = 2, Plan Sequence: Sequence number (also called a reference number) that identifies the credit plan on the account to which the transaction was posted
	PlanSeq string `json:"planSeq,omitempty"`

	//  Max length = 15, Payment Reference Number: Resolution reference number generated during Posting for credits received
	PmtRefNbr string `json:"pmtRefNbr,omitempty"`

	// Format: YYYYMMDD. Posting date of transaction.  Date format is CCYY-MM-DD.
	PostDt string `json:"postDt,omitempty"`

	//  Max length = 9, Points: Field that indicates the number of frequent shopper points earned by the transaction
	Pts string `json:"pts,omitempty"`

	//  Max length = 16, Purchase order number
	PurchOrder string `json:"purchOrder,omitempty"`

	//  Max length = 3, Quantity for the transaction.
	Qty string `json:"qty,omitempty"`

	//  Max length = 1, Monetary qualification indicator. Values are:1 = ATPT-MT-INCOMING-INTRAPROCESSR2 = ATPT-MT-NON-INTERCHANGE3 = ATPT-MT-INCOMING-DOMESTIC6 = ATPT-MT-INCOMING-INTERNATIONAL3 = ATPT-MT-EUROPEAN6 = ATPT-MT-NON-EUROPEAN9 = ATPT-MT-NON-INTERCHANGE-DIFF
	QualInd string `json:"qualInd,omitempty"`

	//  Max length = 5, Identification number of the record
	RecNbr string `json:"recNbr,omitempty"`

	//  Max length = 23, Reference number: Cross Reference for a particular transaction to TRAMS CardHolder Transaction History file where the transaction details (including the scheme 23-digit reference number) are stored.
	RefNbr string `json:"refNbr,omitempty"`

	//  Max length = 17, Relationship payment amount: Payment amount linked to relationship number
	RelPmtAmt string `json:"relPmtAmt,omitempty"`

	//  Max length = 12, Salesperson: Code that identifies the salesclerk associated with the transaction
	Salesperson string `json:"salesperson,omitempty"`

	//  Max length = 9, Stock Keeping Unit number associated with a particular transaction
	SkuNbr string `json:"skuNbr,omitempty"`

	//  Max length = 1, General ledger source code: Code to identify the source of a transaction. Values are:0 = ATPT-MT-GL-LOCAL1 = ATPT-MT-GL-GENERATED2 THRU 6 = ATPT-MT-GL-USER7 = ATPT-MT-GL-WAREHOUSE8 = ATPT-MT-GL-REJECT9 = ATPT-MT-GL-XFR-HISTORY
	Source string `json:"source,omitempty"`

	//  Max length = 15, Ticket Number: Invoice or ticket number associated with the transaction
	TktNbr string `json:"tktNbr,omitempty"`

	//  Max length = 4, Transaction Code: Transaction code associated with the transaction.
	TxnCd string `json:"txnCd,omitempty"`

	//  Max length = 1, TYP
	Typ string `json:"typ,omitempty"`

	//  Max length = 17, Unit price of the transaction
	UnitPrice string `json:"unitPrice,omitempty"`

	//  Max length = 15, Visa transaction ID
	VisaTxnID string `json:"visaTxnId,omitempty"`
}

// Validate validates this txn data for transaction detail inquiry1
func (m *TxnDataForTransactionDetailInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TxnDataForTransactionDetailInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TxnDataForTransactionDetailInquiry1) UnmarshalBinary(b []byte) error {
	var res TxnDataForTransactionDetailInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
