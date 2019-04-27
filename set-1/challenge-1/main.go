package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	res := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(res)
}

// HexToBase64 converts hex encoded string to base64 encoded string
func HexToBase64(in string) string {
	hp, err := hex.DecodeString(in)

	if err != nil {
		panic("Give me a valid hex string!!")
	}
	return base64.StdEncoding.EncodeToString(hp)
}
