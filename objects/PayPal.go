package objects

type PayPal struct {
	Object             string
	PayPalEmail        string
	PayerID            string
	ReturnURL          string
	CancelURL          string
	RequestReferenceID bool
	ReferenceID        string
}
