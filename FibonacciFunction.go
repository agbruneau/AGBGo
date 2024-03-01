package main

import (
    "fmt"
    "time"
)

// fibDoubling calculates the Nth Fibonacci number using the doubling method.
func fibDoubling(n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("n must be non-negative")
    }
    if n < 2 {
        return n, nil
    }
    f, _ := fibDoublingHelper(n)
    return f, nil
}

// fibDoublingHelper is a helper function that calculates the Nth and (N+1)th
// Fibonacci numbers using the doubling method. It returns a tuple (F(n), F(n+1)).
func fibDoublingHelper(n int) (int, int) {
    if n == 0 {
        return 0, 1
    }
    a, b := fibDoublingHelper(n >> 1)
    c := a * ((b << 1) - a)
    d := a*a + b*b
    if n&1 == 1 {
        return d, c + d
    }
    return c, d
}

func main() {
    // Measure the time it takes to calculate large Fibonacci numbers
    start := time.Now()
    result, err := fibDoubling(4500)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("fibDoubling(4500): %d\n", result)
    }
    fmt.Printf("Time taken: %v\n", time.Since(start))

    start = time.Now()
    result, err = fibDoubling(4500000) // This will likely overflow int in Go
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("fibDoubling(4500000): %d\n", result)
    }
    fmt.Printf("Time taken: %v\n", time.Since(start))

    // Note: Calculation of very large Fibonacci numbers (e.g., 45 millionth)
    // may not be practical due to time and integer overflow considerations.
}
