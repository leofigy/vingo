package objects

import (
	"time"
)

/*Account is a client for a given merchant */
type Account struct {
	Object              string
	ID                  string
	VID                 string
	Created             time.Time // Time.RFC3339
	Parent              *Account
	DefaultCurrency     string
	Email               string
	EmailType           string
	Lenguage            string
	NotifyBeforeBilling bool
	Company             string
	Name                string
	ShippingAddress     *Address
	PaymentMethods      []PaymentMethod
	Metadata            map[string]interface{}
	TaxExemptions       []TaxExemption
	tokens              []TokenAmount
	Credit              *Credit
	Entitlements        []Entitlement
	TaxUseCode          string
}
