package objects

import (
	"time"
)

//TimeInterval object represents the time-based credit
//detail associated with a particular grant of time credits.
type TimeInterval struct {
	Object             string
	VID                string
	Years              int
	Months             int
	Days               int
	Metadata           map[string]interface{}
	SortValue          int
	Description        string
	Reason             string
	GrantedByCashBox   bool
	Granted            time.Time
	AccountVID         string
	MerchantAccountID  string
	AutoBillID         string
	MerchantAutoBillID string
	Source             string
}
