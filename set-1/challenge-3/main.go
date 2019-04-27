package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

var alphabet = alphabetMaker()

func main() {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	SingleByteDecoder(in)
}

func SingleByteDecoder(in string) string {
	var highscore float32
	var highscorestring string

	hexChars := alphabet + strings.ToUpper(alphabet)

	for _, item := range hexChars {
		res := SingleByteXOR(in, string(item))
		score := ScoreString(res)
		fmt.Printf("%s (with score %f)\n", res, score)

		if score > highscore {
			highscore = score
			highscorestring = res
		} else if score == highscore {
			highscorestring = highscorestring + " and also string " + res
		}
	}

	fmt.Println("------------------------------------------------")
	fmt.Printf("Highscore was %s with score %f\n", highscorestring, highscore)

	return highscorestring
}

func SingleByteXOR(enc, c string) string {
	dec := HexToByte(enc)
	res := make([]byte, len(dec))
	for i := range dec {
		res[i] = dec[i] ^ []byte(c)[0]
	}
	return string(res)
}

func HexToByte(enc string) []byte {
	res, err := hex.DecodeString(enc)

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

func ScoreString(in string) float32 {
	var score float32

	in = strings.ToLower(in)

	for _, c := range in {
		score += characterScore(c)
	}

	return score / float32(len(in))
}

func characterScore(c rune) float32 {
	scores := [26]float32{8.17, 1.49, 2.78, 4.25, 12.70, 2.23, 2.02, 6.10, 6.97, 0.15, 0.77, 4.03, 2.41, 6.75, 7.51, 1.93, 0.1, 5.99, 6.32, 9.06, 2.76, 0.98, 2.36, 0.15, 1.97, 0.07}

	if c == ' ' {
		return 5
	}

	for i, ac := range alphabet {
		if ac == c {
			return scores[i]
		}
	}
	return 0
}
