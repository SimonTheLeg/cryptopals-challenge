package challenge5

import (
	"encoding/hex"
)

func SolveChallenge() (res string) {
	msg := []byte(`Burning 'em, if you ain't quick and nimble 
I go crazy when I hear a cymbal`)
	key := []byte("ICE")

	return hex.EncodeToString(RepeatingKeyXOR(msg, key))
}

// RepeatingKeyXOR applies a RepeatingKeyXOR function to the input
func RepeatingKeyXOR(msg []byte, key []byte) (res []byte) {
	var keyPos int
	keyLen := len(key)

	for _, b := range msg {
		if keyPos >= keyLen {
			keyPos = 0
		}

		resb := b ^ key[keyPos]
		res = append(res, resb)
		keyPos++
	}
	return
}
