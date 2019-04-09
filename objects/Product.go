package objects

import (
	"time"
)

type Product struct {
	Object             string
	ID                 string
	VID                string
	Created            time.Time
	Descriptions       []ProductDescription
	Status             string
	TaxClassification  string
	DefaultBillingPlan *BillingPlan
	DefaultRatePlan    *RatePlan
	EndOfLife          time.Time
	Metadata           map[string]interface{}
	Entitlements       []Entitlement
	BillingDescriptor  string
	CreditGranted      *Credit
	prices             []ProductPrice
	BundledProducts    []Product
}
