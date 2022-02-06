package main

import (
	"fmt"
	"math/big"
	"time"
)

func fib(n int) *big.Int {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Loop while a is smaller than 1e100.
	for i := 0; i <= n-1; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}

func main() {
	start := time.Now()
	fmt.Printf("Fib(1 000 000) is %d \n", fib(1000000))
	elapsed := time.Since(start)
	fmt.Printf("time eclibaced is %s \n", elapsed)
}
