package objects

type CreditCard struct {
	Object         string
	VID            string
	Account        string
	Bin            string
	LastDigits     string
	AccountLength  string
	HashType       string
	AccountHash    string
	ExpirationDate string
	ECA            *ExtendedCardAttributes
	CVN            string
}
