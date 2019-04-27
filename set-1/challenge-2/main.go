package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	res := FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println(res)
}

// FixedXOR takes in two hex encoded strings and bytes XORs them
func FixedXOR(i1, i2 string) string {
	buf1, err := hex.DecodeString(i1)

	if err != nil {
		panic("Give me a valid first hex string!!")
	}

	buf2, err := hex.DecodeString(i2)

	if err != nil {
		panic("Give me a valid hex string!!")
	}

	res := make([]byte, len(buf1))

	for i := range buf1 {
		res[i] = buf1[i] ^ buf2[i]
	}

	return hex.EncodeToString(res)
}
