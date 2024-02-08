package main

import (
	"fmt"
	"math/big"
)

// fibonacci génère une liste de nombres de Fibonacci jusqu'à n inclusivement.
func fibonacci(n int) []*big.Int {
	// Initialiser les deux premiers nombres de Fibonacci
	fibNumbers := make([]*big.Int, n+1)
	fibNumbers[0] = big.NewInt(0)
	if n > 0 {
		fibNumbers[1] = big.NewInt(1)
	}

	// Calculer les nombres de Fibonacci pour n > 1
	for i := 2; i <= n; i++ {
		// fibNumbers[i] = fibNumbers[i-1] + fibNumbers[i-2]
		nextFib := big.NewInt(0)
		nextFib.Add(fibNumbers[i-1], fibNumbers[i-2])
		fibNumbers[i] = nextFib
	}

	return fibNumbers
}

func main() {
	// Modifier cette valeur pour générer plus ou moins de nombres de Fibonacci
	n := 20000 // Par exemple, les 100 premiers nombres de Fibonacci

	// Générer et afficher la liste de Fibonacci
	fibList := fibonacci(n)
	for i, fibNum := range fibList {
		fmt.Printf("Fibonacci(%d): %s\n", i, fibNum.String())
	}
}
