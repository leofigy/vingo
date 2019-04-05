package objects_test

import (
	"log"
	"testing"

	vin "github.com/angel-figueroa/vingo/objects"
)

func TestAccountCreation(t *testing.T) {
	account := &vin.Account{}
	if account == nil {
		t.Error("cannot create account")
	}

	log.Println(account)

}
