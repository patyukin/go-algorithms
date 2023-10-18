package lib

// Factorial Функция для вычисления факториала
func Factorial(n int64) int64 {
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

// BinomialCoefficient Функция для вычисления комбинаций
func BinomialCoefficient(n int64, k int64) int64 {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
