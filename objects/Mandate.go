package objects

import (
	"time"
)

type Mandate struct {
	Object     string
	Created    time.Time
	Identifier string
	Status     string // Active,Cancelled,Expired,Failed,Pending
}
