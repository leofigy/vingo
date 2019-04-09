package objects

type RatePlan struct {
	Object               string
	ID                   string
	VID                  string
	Description          string
	MultiplyRatedUnitsBy interface{} //*RatedUnit TBD in spec
	Tier                 interface{} //[]RatePlanTier TBD in spec
	RoudingDecimals      int64
	IncludedUnits        float64
	MinimumFee           interface{} //*RatePlanPrice TBD in spec
	MaximumFee           interface{} //*RatePlanPrice TBD in spec
	RatePlanModel        string      //LicenseBased,UsageBased
	Status               string      // Active,Suspended
	Metadata             map[string]interface{}
}
