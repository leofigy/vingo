package objects

type ExtendedCardAttributes struct {
	Object               string
	CommercialCard       int64
	ConsumerCard         int64
	CreditCard           int64
	SignatureDebitCard   int64
	PinlessDebitCard     int64
	GiftCard             int64
	PrepaidCard          int64
	PayrollCard          int64
	HealthcareCard       int64
	VirtualAccountNumber int64
	Reloadable           int64
	CountryOfIssuance    string
	DurbinRegulated      int64
	Affluent             int64
	MassAffluent         int64
	CardSubscription     string
}
