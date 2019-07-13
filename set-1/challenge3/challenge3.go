package challenge3

import (
	"encoding/hex"
	"strings"
)

// SolveChallenge returns the result of challenge 3
// To be called by the testing suite
func SolveChallenge() (res string, score float32, key rune) {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	return FindKeyAndDecodeString(in)
}

// FindKeyAndDecodeString takes in an encrypted string, tries out all lowercase and uppercase letters
// of the alphabet and returns the cleartext string with the highest score, it's actual score and the
// alphabet character used to encrypt it
func FindKeyAndDecodeString(in string) (bestString string, bestScore float32, bestCharacter rune) {
	// Generate all letters of the alphabet
	alphabet := alphabetMaker() + strings.ToUpper(alphabetMaker())

	for _, letter := range alphabet {
		// Try to decrypt the string with the letter
		decMsg := SingleByteXOR(in, string(letter))

		// Score the decrypted message
		score := ScoreString(decMsg)

		// Compare the score to the highscore
		if score >= bestScore {
			bestScore = score
			bestString = decMsg
			bestCharacter = letter
		}
	}

	return
}

// SingleByteXOR takes in an encrypted string
// and a character to decrypt and decrypts
// the string
func SingleByteXOR(enc, c string) string {
	key := byte(c[0])
	dec := HexToByte(enc)
	res := make([]byte, len(dec))
	for i := range dec {
		res[i] = dec[i] ^ key
	}
	return string(res)
}

// HexToByte converts a Hex String to a byte array
func HexToByte(input string) []byte {
	res, err := hex.DecodeString(input)

	if err != nil {
		panic(err)
	}

	return res
}

func alphabetMaker() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}

// ScoreString takes in a string and calculates its reality score
func ScoreString(in string) float32 {
	var score float32

	for _, c := range in {
		score += characterScore(c)
	}

	return score / float32(len(in))
}

func characterScore(c rune) float32 {
	// scores := [26]float32{8.17, 1.49, 2.78, 4.25, 12.70, 2.23, 2.02, 6.10, 6.97, 0.15, 0.77, 4.03, 2.41, 6.75, 7.51, 1.93, 0.1, 5.99, 6.32, 9.06, 2.76, 0.98, 2.36, 0.15, 1.97, 0.07}
	var s float32

	switch {
	case 'a' <= c && c <= 'z':
		s = 10
	case 'A' <= c && c <= 'Z':
		s = 5
	case '0' <= c && c <= '9':
		s = 3
	default:
		switch c {
		case ';', ':', '\'', '"', '\t', '\r', '\n':
			s = 1
		case '.', '!', '?':
			s = 2
		case ',':
			s = 3
		case ' ':
			s = 15
		default:
			s = -25
		}
	}
	return s
}
