package http

import (
	"fmt"
	"testing"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/stretchr/testify/require"
	"go-tamboom/model"
)

const CVV = "123"

func TestOmise_CreateCard(t *testing.T) {
	api, err := NewOmise()
	require.NoError(t, err)

	cases := map[string]struct {
		donation    model.Donation
		isErrorCase bool
	}{
		"Success": {
			donation: model.Donation{
				Name:            "Penny Morebucks",
				AmountSubunits:  10000,
				CCNumber:        "4242424242424242",
				CVV:             CVV,
				ExpirationMonth: time.Month(10),
				ExpirationYear:  2022,
			},
			isErrorCase: false,
		},
		"Expired Card": {
			donation: model.Donation{
				Name:            "Penny Morebucks",
				AmountSubunits:  10000,
				CCNumber:        "4242424242424242",
				CVV:             CVV,
				ExpirationMonth: time.Month(10),
				ExpirationYear:  2020,
			},
			isErrorCase: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := api.CreateCard(tc.donation)
			if tc.isErrorCase {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestOmise_CreateCharge(t *testing.T) {
	api, err := NewOmise()
	require.NoError(t, err)

	cases := map[string]struct {
		donation            model.Donation
		expectedStatus      omise.ChargeStatus
		expectedFailureCode string
	}{
		"Success": {
			donation: model.Donation{
				Name:            "Penny Morebucks",
				AmountSubunits:  10000,
				CCNumber:        "4242424242424242",
				CVV:             CVV,
				ExpirationMonth: time.Month(10),
				ExpirationYear:  2022,
			},
			expectedStatus: omise.ChargeStatus("successful"),
		},
		"Failed": {
			donation: model.Donation{
				Name:            "Penny Morebucks",
				AmountSubunits:  10000,
				CCNumber:        "4111111111110014",
				CVV:             CVV,
				ExpirationMonth: time.Month(10),
				ExpirationYear:  2022,
			},
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "payment_rejected",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			card, err := api.CreateCard(tc.donation)
			require.NoError(t, err)

			charge, err := api.CreateCharge(tc.donation, card)
			require.NoError(t, err)

			require.Equal(t, tc.expectedStatus, charge.Status, fmt.Sprintf("[%s] Expected: %s, Got: %s", name, tc.expectedStatus, charge.Status))
			if tc.expectedStatus != ("successful") {
				require.Equal(t, &tc.expectedFailureCode, charge.FailureCode, fmt.Sprintf("[%s] Expected: %p, Got: %p", name, &tc.expectedFailureCode, charge.FailureCode))
			}
		})
	}
}

func TestOmiseClient_CreateCharge(t *testing.T) {
	c, err := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	require.NoError(t, err)

	cases := map[string]struct {
		ccNumber            string
		expectedStatus      omise.ChargeStatus
		expectedFailureCode string
	}{
		"Success Case": {
			ccNumber:       "4242424242424242",
			expectedStatus: omise.ChargeStatus("successful"),
		},
		"insufficient_fund": {
			ccNumber:            "4111111111140011",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "insufficient_fund",
		},
		"failed_processing": {
			ccNumber:            "4111111111120013",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "failed_processing",
		},
		"payment_rejected": {
			ccNumber:            "4111111111110014",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "payment_rejected",
		},
		"failed_fraud_check": {
			ccNumber:            "4111111111190016",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "failed_fraud_check",
		},
		"stolen_or_lost_card": {
			ccNumber:            "5555551111100012",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "stolen_or_lost_card",
		},
		"invalid_account_number": {
			ccNumber:            "5555551111150017",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "invalid_account_number",
		},
		"invalid_security_code": {
			ccNumber:            "5555551111130001",
			expectedStatus:      omise.ChargeStatus("failed"),
			expectedFailureCode: "invalid_security_code",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			card, createToken := &omise.Card{}, &operations.CreateToken{
				Name:            "Somchai Prasert",
				Number:          tc.ccNumber,
				ExpirationMonth: time.Month(10),
				ExpirationYear:  2022,

				City:         Tokyo,
				PostalCode:   ZipCode,
				SecurityCode: CVV,
			}

			e := c.Do(card, createToken)
			require.Nil(t, e, fmt.Sprintf("[%s] Failed to create card", name))

			charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
				Amount:   10000,
				Currency: JPY,
				Card:     card.ID,
			}

			e = c.Do(charge, createCharge)
			require.Nil(t, e, fmt.Sprintf("[%s] Failed to create charge", name))

			require.Equal(t, tc.expectedStatus, charge.Status, fmt.Sprintf("[%s] Expected: %s, Got: %s", name, tc.expectedStatus, charge.Status))
			if tc.expectedStatus != ("successful") {
				require.Equal(t, &tc.expectedFailureCode, charge.FailureCode, fmt.Sprintf("[%s] Expected: %p, Got: %p", name, &tc.expectedFailureCode, charge.FailureCode))
			}
		})
	}
}
