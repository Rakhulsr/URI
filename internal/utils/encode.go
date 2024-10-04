package utils

import (
	"crypto/sha256"
	"os"

	"github.com/Rakhulsr/go-url-shortener/internal/helper"
	"github.com/itchyny/base58-go"
)

func Sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func Base56Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	endcoded, err := encoding.Encode(bytes)

	if err != nil {
		helper.PanicIfError("Error while encode base58", err)
		os.Exit(1)
	}

	return string(endcoded)
}
