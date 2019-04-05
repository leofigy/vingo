package objects

/*Credit object represents a service credit granted to an account or subscription,
taking the form of either currency, token, or time amounts.*/
type Credit struct {
	Object          string
	TokenAmounts    []TokenAmount
	TimeIntervals   []TimeInterval
	CurrencyAmounts []CurrencyAmount
}
