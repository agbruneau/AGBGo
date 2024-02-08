package main

import (
	"fmt"
	"math/big"
	"time"
)

// fibDoubling calculates the Nth Fibonacci number using the doubling method.
func fibDoubling(n int) *big.Int {
	if n < 2 || n > 100000001 {
		return big.NewInt(int64(n))
	}
	result := _fibDoubling(n) // No need to unpack here, since it returns a slice
	return result[0] // Access the first element of the slice, equivalent to F(n)
}

// _fibDoubling is a helper function that calculates the Nth Fibonacci number using the doubling method.
// It returns a slice of *big.Int representing (F(n), F(n+1)).
func _fibDoubling(n int) []*big.Int {
	if n == 0 {
		return []*big.Int{big.NewInt(0), big.NewInt(1)}
	}
	aB := _fibDoubling(n >> 1)
	a := aB[0]
	b := aB[1]
	c := new(big.Int).Sub(new(big.Int).Lsh(b, 1), a)
	c.Mul(c, a)
	d := new(big.Int).Add(new(big.Int).Mul(a, a), new(big.Int).Mul(b, b))
	if n&1 == 1 {
		return []*big.Int{d, new(big.Int).Add(c, d)}
	} else {
		return []*big.Int{c, d}
	}
}

func main() {
	// Example usage of fibDoubling with timing
	n1 := 10000000

	start := time.Now()
	fmt.Printf("fibDoubling(%d): %s\n", n1, fibDoubling(n1))
	fmt.Printf("Time taken for n1: %s\n", time.Since(start))
}
