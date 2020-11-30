package cipher

import (
	"errors"
	"io/ioutil"
	"os"
)

const DecipherExtension = ".decipher"

// Decrypt describes a series of decryption functions
type Decrypt interface {
	// GetDecipherText deciphers a file, returning the output or an error if one occurs
	GetDecipherText() (string, error)
}

// Decipher is a struct that is used to handle decryption
type Decipher struct {
	fileName string
}

// NewDecipher returns a new Decipher
func NewDecipher(f string) Decrypt {
	return &Decipher{
		fileName: f,
	}
}

func (d Decipher) GetDecipherText() (string, error) {
	cipheredFile, err := os.Open(d.fileName)
	if err != nil {
		return "", err
		//fmt.Printf("[FILE READ ERROR] - Couldn't read encrypted file: %s\n", err)
	}
	defer cipheredFile.Close()

	reader, err := NewRot128Reader(cipheredFile)
	if err != nil {
		return "", errors.New("[DECIPHER INITIALIZATION ERROR] Failed to load reader")
	}

	output, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", errors.New("[READER ERROR] Failed to load reader")
	}

	if err := ioutil.WriteFile(d.fileName+DecipherExtension, output, os.ModePerm); err != nil {
		return "", errors.New("[WRITE ERROR] Failed to load reader")
		//fmt.Printf("[WRITE ERROR] - failed to write deciphered file: %s\n", err)
	}

	return d.fileName + DecipherExtension, nil
}
