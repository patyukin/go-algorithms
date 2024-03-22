package main

import (
	"fmt"
	"github.com/patyukin/go-algorithms/algo/combinatorics/lib"
)

func main() {
	result := lib.BinomialCoefficient(10, 2) * lib.BinomialCoefficient(10, 3)
	result += lib.BinomialCoefficient(10, 3) * lib.BinomialCoefficient(10, 2)
	result += lib.BinomialCoefficient(10, 4) * lib.BinomialCoefficient(10, 1)
	result += lib.BinomialCoefficient(10, 5)
	fmt.Println(result)
}
