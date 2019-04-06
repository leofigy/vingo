package objects

import (
	"time"
)

type PaymentMethod struct {
	Object                string
	ID                    string
	VID                   string
	Created               time.Time
	Type                  string // check enum
	CreditCard            *CreditCard
	ECP                   *ECP
	DirectDebit           *DirectDebit
	PayPal                *PayPal
	Boleto                *Boleto
	HostedPage            *HostedPage
	Token                 *Token
	MAP                   *MerchantAcceptedPayment
	CarrierBilling        *CarrierBilling
	ApplePay              *ApplePay
	Metadata              map[string]interface{}
	AccountHolder         string
	BillingAddress        *Address
	CustomerSpecifiedType string
	CustomerDescription   string
	Currency              string
	Primary               bool
	Active                bool
	ValidationTransaction *Transaction
	Policy                map[string]interface{}
}
