package main

import (
	"fmt"
)

// Factorial Функция для вычисления факториала
func Factorial(n int64) int64 {
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func BinomialCoefficient(n int64, k int64) int64 {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}

func main() {
	// Анализ различных возможных вариантов комитетов
	committeeWays := BinomialCoefficient(10, 2) * BinomialCoefficient(10, 3)
	committeeWays += BinomialCoefficient(10, 3) * BinomialCoefficient(10, 2)
	committeeWays += BinomialCoefficient(10, 4) * BinomialCoefficient(10, 1)
	committeeWays += BinomialCoefficient(10, 5) * BinomialCoefficient(10, 0)
	fmt.Printf("Количество способов сформировать комитет: %v\n", committeeWays)
}
