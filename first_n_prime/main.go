package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Gather input from the user.
	// Repeats until the input is a valid integer.
	var (
		answer int
		err    error
	)
	for ok := true; ok; ok = err != nil {
		fmt.Println("? How many prime numbers you would like:")
		if _, err = fmt.Scan(&answer); err != nil {
			fmt.Println("Please provide a valid integer.")
		}
	}
	if answer == 0 {
		return
	}

	fmt.Println("Prime numbers in ascending order:")
	count := 0
	// Assumes range of natural numbers is 0-100.
	for i := 2; i < 101; i++ {
		if count == answer {
			return
		}

		// Gives a 100% accurate reading for n < 2^64 (int64).
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Println(i)
			count += 1
		}
	}
}
