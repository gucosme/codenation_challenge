package data

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	NumberOfCases   int    `json:"numero_casas"`
	Token           string `json:"token"`
	Encrypted       string `json:"cifrado"`
	Decripted       string `json:"decifrado"`
	EncryptedResume string `json:"resumo_criptografico"`
}

func (d Data) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d Data) String() string {
	return fmt.Sprintf(
		"numero_casas: %d\ntoken: %s\ncifrado: %s\ndecifrado: %s\nresumo_criptografado: %s",
		d.NumberOfCases,
		d.Token,
		d.Encrypted,
		d.Decripted,
		d.EncryptedResume,
	)
}
