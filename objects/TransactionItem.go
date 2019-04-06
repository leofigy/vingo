package objects

import (
	"time"
)

type TransactionItem struct {
	Object              string
	SKU                 string
	IndexNumber         int
	ItemType            string
	Name                string
	SubscriptionItem    *SubscriptionItem
	Price               int64
	Metadata            map[string]interface{}
	Quantity            int64
	TaxClassifications  string
	Tokens              []TokenAmount
	CampaignCode        string
	CampaignID          string
	CampaignDescription string
	ServicePeriodStarts time.Time
	ServicePeriodEnds   time.Time
	TaxType             string
	Tax                 []TaxItem
	RelatedTransactions []string
	ItemRefunds         []TransactionItemRefundSummary
	Discount            float64
	SubTotal            float64
	Total               float64
}
