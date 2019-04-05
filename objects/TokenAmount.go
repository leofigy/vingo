package objects

/*TokenAmount object represents the token-based credit detail
associated with a particular grant of token credits.
*/
type TokenAmount struct {
	ID     string
	Token  *Token
	Amount int64
	source string
}
