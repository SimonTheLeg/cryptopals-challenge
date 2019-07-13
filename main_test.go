package main

import (
	"fmt"
	"testing"

	challenge3 "github.com/simontheleg/golang-crypto-challenge/set-1/challenge3"
	challenge4 "github.com/simontheleg/golang-crypto-challenge/set-1/challenge4"
	challenge5 "github.com/simontheleg/golang-crypto-challenge/set-1/challenge5"
)

func TestChallenge3(t *testing.T) {
	res, score, key := challenge3.SolveChallenge()
	fmt.Printf("Highest Score was string '%s' with a score of '%f', key was %q\n", res, score, key)
}

func TestChallenge4(t *testing.T) {
	res, score, key := challenge4.SolveChallenge()
	fmt.Printf("Highest Score was string '%s' with a score of '%f', key was %q\n", res, score, key)
}

func TestChallenge5(t *testing.T) {
	res := challenge5.SolveChallenge()
	fmt.Println(res)
}
