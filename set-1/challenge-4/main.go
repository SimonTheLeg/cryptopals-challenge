package main

import (
	"bufio"
	"fmt"
	"os"

	challenge3 "github.com/simontheleg/golang-crypto-challenge/set-1/challenge-3"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var highscore float32
	var highscorestring string
	for scanner.Scan() {
		enc := scanner.Text()
		scorestring, score := challenge3.SingleByteDecoder(enc)

		if score > highscore {
			highscore = score
			highscorestring = scorestring
		}
	}
	fmt.Println("==============================================")
	fmt.Printf("Highest score was '%s' (with a score of %f)", highscorestring, highscore)

}
