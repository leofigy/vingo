package objects

import (
	"time"
)

type Transaction struct {
	Object                        string
	ID                            string
	VID                           string
	Created                       time.Time
	Amount                        float64
	OriginalAmount                float64
	Currency                      string
	DivisionNumber                string
	PreviousID                    string
	Account                       *Account
	SourcePaymentMethod           *PaymentMethod
	DestinationPaymentMethod      *PaymentMethod
	ECPTransactionType            string
	StatusLog                     []TransactionStatus
	PaymentProcessor              string
	PaymentProcessorTransactionID string
	ShippingAddress               *Address
	Metadata                      map[string]interface{}
	Items                         []TransactionItem
	Affiliate                     string
	SubAffiliate                  string
	UserAgent                     string
	PreferredNotificationLanguage string
	SourceMacAddress              string
	SourceIP                      string
	BillingDescriptior            string
	SalesTaxAddress               Address
	VerificationCode              string
	SubscriptionSequence          int64
	BillingPlanSequence           int64
	Mandate                       *Mandate
	OriginalBillingData           time.Time
	BillingAttempt                int64
	Subscription                  *Subscription
	Score                         int64
	ScoreCodes                    *ScoreCode
	Policy                        map[string]interface{}
	ToBeCaptured                  bool
}
