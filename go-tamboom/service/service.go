package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"go-tamboom/cipher"
	"go-tamboom/http"
	"go-tamboom/model"
)

// Service defines a series of functions for business logic
type Service interface {
	// MakeDecipherText creates a 128 ciphered text file, based on the file name that is passed to the Service
	MakeDecipherText()
	// ProcessDonations processes a file of donations, creating Donation structs and then handling the Summary
	ProcessDonations()
}

// ChargerService represents the main service of the application
type ChargerService struct {
	fileName string
	decrypt  cipher.Decrypt
	encrypt  cipher.Encrypt
	api      http.API
}

// NewChargerService creates a new ChargerService, with a file to be processed
// Should an error occur during the creation, no error will be produced
// This is meant to play like a batch job where errors are recorded but do not stop the service
func NewChargerService(fileName string) Service {
	api, err := http.NewOmise()
	if err != nil {
		log.Print(err)
	}

	return &ChargerService{
		fileName: fileName,
		decrypt:  cipher.NewDecipher(fileName),
		encrypt:  cipher.NewEncipher(fileName),
		api:      api,
	}
}

// ProcessDonations reads the file from the Service, and then creates calls with the API to generate a Summary
func (c *ChargerService) ProcessDonations() {
	fileName, err := c.decrypt.GetDecipherText()
	if err != nil {
		log.Fatal(err)
	}

	records := c.getRecordsFromFile(fileName)
	fmt.Println(records)
	log.Printf("[Service] Processing donations...\n")
	summary := c.processDonationRecords(records)
	log.Printf("[Service] Done.\n")
	summary.PrintSummary()
}

func (c *ChargerService) getRecordsFromFile(fn string) [][]string {
	data, err := os.Open(fn)
	if err != nil {
		log.Printf("[READER ERROR] - Couldn't read from reader: %s\n", err)
	}
	defer os.Remove(fn)

	reader := csv.NewReader(data)
	records, _ := reader.ReadAll()
	return records
}

func (c *ChargerService) processDonationRecords(records [][]string) Summary {
	summary := NewSummary()
	for i, data := range records {
		if i == 0 {
			//this is the CSV header, so we don't need this, skip
			continue
		}

		d := model.NewDonation(data[0], data[1], data[2], data[3], data[4], data[5])

		card, err := c.api.CreateCard(d)
		if err != nil {
			log.Printf("[OMISE] Error creating card: %s\n", err)
			summary.AddToFailure(d.AmountSubunits)
			continue
		}

		_, err = c.api.CreateCharge(d, card)
		if err != nil {
			log.Printf("[OMISE] Error creating charge: %s\n", err)
			summary.AddToFailure(d.AmountSubunits)
			continue
		}

		summary.AddSuccess(d.Name, d.AmountSubunits)
	}

	return summary
}

// MakeDecipherText enciphers the file from the Service
func (c *ChargerService) MakeDecipherText() {
	c.encrypt.MakeDecipherText()
}
