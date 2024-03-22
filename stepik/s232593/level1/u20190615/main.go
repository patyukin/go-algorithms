package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	arr := getNumbers(a, b)
	for i := 0; i < len(arr)/2; i++ {
		left, right := i, len(arr)-1-i
		arr[left], arr[right] = arr[right], arr[left]
	}

	for _, num := range arr {
		fmt.Print(num)
	}
}

func getNumbers(num, sliced int) []int {
	var result []int

	if num < 10 {
		return append(result, num)
	}

	for {
		if num < 10 {
			if num != sliced {
				result = append(result, num)
			}

			break
		}

		a := num % 10
		if a != sliced {
			result = append(result, a)
		}

		num /= 10
	}

	return result
}
