package cipher

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Encrypt describes a series of encryption functions
type Encrypt interface {
	// MakeDecipherText enciphers a file, using the 128 Caesar Cipher
	MakeDecipherText()
}

// Encipher is a struct that is used to handle encryption
type Encipher struct {
	fileName string
}

// NewEncipher returns a new Encipher
func NewEncipher(f string) Encrypt {
	return &Encipher{
		fileName: f,
	}
}

func (e Encipher) MakeDecipherText() {
	ciphered, err := os.Create(e.fileName)
	if err != nil {
		fmt.Printf("[FILE OPEN ERROR] - Couldn't open new file: %s\n", err)
	}
	defer ciphered.Close()

	cw, err := NewRot128Writer(ciphered)
	if err != nil {
		fmt.Printf("[INTIALIZATION ERROR] - Failed to load writer: %s\n", err)
	}

	writer := csv.NewWriter(cw)
	var data = [][]string{
		{"Donor", "Number", "Donation"},
		{"Yshtola Ruhn", "1111111111111111", "2200"},
		{"Thancred Waters", "1111111111111112", "4300"},
		{"Alisaie Leveilleur", "1111111111111113", "3900"},
	}

	err = writer.WriteAll(data)
	if err != nil {
		fmt.Println(e)
	}
}
