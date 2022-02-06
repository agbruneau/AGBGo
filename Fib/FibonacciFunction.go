// Référence : https://levelup.gitconnected.com/coding-problem-nth-fibonacci-number-in-golang-68db104e318
package main

import (
	"fmt"
	"time"
)

func fibRecursive(n int) int {
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibDynamic(n int) int {
	baseCases := map[int]int{
		1: 0,
		2: 1,
	}
	return computeCache(n, baseCases)
}

func computeCache(n int, cache map[int]int) int {
	if val, found := cache[n]; found {
		return val
	}
	cache[n] = computeCache(n-1, cache) + computeCache(n-2, cache)
	return cache[n]
}

func fibFly(n int) uint64 {
	prevFibs := []uint64{0, 1}
	for i := 3; i <= n; i++ {
		nextFib := prevFibs[0] + prevFibs[1]
		prevFibs = []uint64{prevFibs[1], nextFib}
	}
	if n > 1 {
		return prevFibs[1]
	}
	return prevFibs[0]
}

func main() {
	start := time.Now()
	fmt.Printf("FibRecursive(45) is %d \n", fibRecursive(45))
	elapsed := time.Since(start)
	fmt.Printf("time eclibaced is %s \n\n", elapsed)

	start = time.Now()
	fmt.Printf("FibDynamic(4 500 000) is %d \n", fibDynamic(4500000))
	elapsed = time.Since(start)
	fmt.Printf("time eclibaced is %s \n\n", elapsed)

	// Comparaison avec FibonacciFunction.py
	start = time.Now()
	fmt.Printf("FibFly(4 500) is %d \n", fibFly(4500))
	elapsed = time.Since(start)
	fmt.Printf("time eclibaced is %s \n", elapsed)

	start = time.Now()
	fmt.Printf("FibFly(4 500 000) is %d \n", fibFly(4500000))
	elapsed = time.Since(start)
	fmt.Printf("time eclibaced is %s \n", elapsed)

	start = time.Now()
	fmt.Printf("FibFly(45 000 000) is %d \n", fibFly(45000000))
	elapsed = time.Since(start)
	fmt.Printf("time eclibaced is %s \n", elapsed)

	start = time.Now()
	fmt.Printf("FibFly(450 000 000) is %d \n", fibFly(450000000))
	elapsed = time.Since(start)
	fmt.Printf("time eclibaced is %s \n", elapsed)

}
