package main

import (
	"fmt"
	"github.com/patyukin/go-algorithms/algo/combinatorics/lib"
)

func main() {
	// Анализ различных возможных вариантов комитетов
	committeeWays := lib.BinomialCoefficient(10, 2) * lib.BinomialCoefficient(10, 3)
	committeeWays += lib.BinomialCoefficient(10, 3) * lib.BinomialCoefficient(10, 2)
	committeeWays += lib.BinomialCoefficient(10, 4) * lib.BinomialCoefficient(10, 1)
	committeeWays += lib.BinomialCoefficient(10, 5) * lib.BinomialCoefficient(10, 0)
	fmt.Printf("Количество способов сформировать комитет: %v\n", committeeWays)
}
