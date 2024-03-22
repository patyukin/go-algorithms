package main

import (
	"fmt"
	"github.com/patyukin/go-algorithms/algo/combinatorics/lib"
)

func main() {
	result := lib.BinomialCoefficient(5, 1) * lib.BinomialCoefficient(4, 1) * lib.BinomialCoefficient(6, 1)
	fmt.Println(result)
}
