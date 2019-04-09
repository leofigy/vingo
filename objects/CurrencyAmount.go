package objects

import (
	"time"
)

// CurrencyAmount field
type CurrencyAmount struct {
	Object           string
	VID              string
	Currency         string
	Amount           float64
	Metadata         map[string]interface{}
	SortValue        int64
	Description      string
	Reason           string
	GrantedByCashBox bool
	Granted          time.Time
	Account          *Account
	Subscription     *Subscription
}
