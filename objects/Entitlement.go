package objects

import (
	"time"
)

/*Entitlement object represents the services to which a customer is entitled.
Entitlement can reside at the Account (granted directly to the Customer)
or the Subscription level (granted due to purchase or presence of an active subscription).
*/
type Entitlement struct {
	Object           string
	ID               string
	Description      string
	Account          *Account
	Subscription     *Subscription
	BillingPlan      *BillingPlan
	SubscriptionItem *SubscriptionItem
	Product          *Product
	Starts           time.Time
	End              time.Time
	Active           bool
}
