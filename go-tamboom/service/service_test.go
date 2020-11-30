package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go-tamboom/http"
)

func TestChargerService_ProcessDonationRecords(t *testing.T) {
	records := make([][]string, 0)
	records = append(records, []string{"Name", "AmountSubunits", "CCNumber", "CVV", "ExpMonth", "ExpYear"})
	records = append(records, []string{"Mr. Grossman R Oldbuck", "2879410", "5375543637862918", "488", "11", "2021"})
	records = append(records, []string{"Mr. Ferdinand H Took-Brandybuck", "2253551", "5238569266360327", "052", "11", "2019"})
	records = append(records, []string{"Ms. Estella R Boffin", "1245821", "4556490107025705", "653", "11", "2020"})
	records = append(records, []string{"Mrs. Camelia G Took", "110051", "5320073142514321", "998", "4", "2022"})
	records = append(records, []string{"Ms. Daisy C Tûk", "2213937", "4556712499523363", "814", "5", "2022"})
	records = append(records, []string{"Ms. Primrose F Smallburrow", "4267320", "5426958001804693", "140", "6", "2023"})
	records = append(records, []string{"Mrs. Esmeralda M Button", "3069758", "5243405935045374", "767", "10", "2020"})

	api, err := http.NewOmise()
	require.NoError(t, err)

	c := ChargerService{
		fileName: "blah",
		decrypt:  nil,
		encrypt:  nil,
		api:      api,
	}

	summary := c.processDonationRecords(records)
	require.Equal(t, int64(16039848), summary.total)
	require.Equal(t, int64(5323309), summary.failures)
	require.Equal(t, 5, len(summary.successes))
}
