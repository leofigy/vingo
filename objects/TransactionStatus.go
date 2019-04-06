package objects

import "time"

type TransactionStatus struct {
	Object            string
	Status            string
	Created           time.Time
	PaymentMethodType string

	// spec does not provide mor details on in
	CreditCardStatus     string
	ECPstatus            string
	BoletoStatus         string
	HostedPageStatus     string
	PayPalStatus         string
	DirectDebitStatus    string
	CarrierBillingStatus string
	AmazonStatus         string
	SkrillStatus         string
	ApplePayStatus       string

	VinAVS               string
	FundingSourceBalance float64
}
