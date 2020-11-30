package http

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"go-tamboom/model"
)

const (
	// Read these from environment variables or configuration files!
	// Since this is a test and not meant for any commercial use, we'll just use these
	OmisePublicKey = "pkey_test_5m1gc72exlp41sox2ql"
	OmiseSecretKey = "skey_test_5m1gc72f5zdltcw8t68"
	JPY            = "jpy"
	Tokyo          = "Tokyo"
	ZipCode        = "1100001"
)

// API describes a series of api call functions
type API interface {
	//CreateCard creates a new Card on the Omise API, returns an error should something fail during the call
	CreateCard(model.Donation) (*omise.Card, error)
	//CreateCharge creates a new Charge on the Omise API, returns an error should something fail during the call
	CreateCharge(model.Donation, *omise.Card) (*omise.Charge, error)
}

// Omise is a struct, used to store the Omise Client
type Omise struct {
	client *omise.Client
}

// NewOmise returns a new Omise
func NewOmise() (API, error) {
	c, err := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if err != nil {
		return nil, err
	}

	return &Omise{
		client: c,
	}, nil
}

func (o *Omise) CreateCard(donation model.Donation) (*omise.Card, error) {
	card, createToken := &omise.Card{}, &operations.CreateToken{
		Name:            donation.Name,
		Number:          donation.CCNumber,
		ExpirationMonth: donation.ExpirationMonth,
		ExpirationYear:  donation.ExpirationYear,

		City:         Tokyo,
		PostalCode:   ZipCode,
		SecurityCode: donation.CVV,
	}

	if e := o.client.Do(card, createToken); e != nil {
		return nil, e
	}

	return card, nil
}

func (o *Omise) CreateCharge(donation model.Donation, card *omise.Card) (*omise.Charge, error) {
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   donation.AmountSubunits,
		Currency: JPY,
		Card:     card.ID,
	}
	if e := o.client.Do(charge, createCharge); e != nil {
		return nil, e
	}
	return charge, nil
}
