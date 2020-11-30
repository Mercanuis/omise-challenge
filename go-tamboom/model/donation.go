// Package model contains structs used by the service to store and handle data
package model

import (
	"fmt"
	"strconv"
	"time"
)

// Donation represents a row in the CSV file, representing a donation
type Donation struct {
	Name            string
	AmountSubunits  int64
	CCNumber        string
	CVV             string
	ExpirationMonth time.Month
	ExpirationYear  int
}

// NewDonation creates and returns a new Donation
func NewDonation(name, amount, ccNumber, cvv, expMonth, expYear string) Donation {
	amt, _ := strconv.ParseInt(amount, 10, 64)
	m, _ := strconv.Atoi(expMonth)
	em := time.Month(m)
	ey, _ := strconv.Atoi(expYear)

	return Donation{
		Name:            name,
		AmountSubunits:  amt,
		CCNumber:        ccNumber,
		CVV:             cvv,
		ExpirationMonth: em,
		ExpirationYear:  ey,
	}
}

// ToString returns a debug string of the Donation's data
func (d Donation) ToString() string {
	res := fmt.Sprintf("Donation - [Donor = [%s], amount = [%d], CCNumber = [%s], CVV = [%s], expiration = [%d/%d]", d.Name, d.AmountSubunits, d.CCNumber, d.CVV, d.ExpirationMonth, d.ExpirationYear)
	return res
}
