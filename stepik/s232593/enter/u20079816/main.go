package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	var result []int
	for i := 0; i < a; i++ {
		var b int
		fmt.Scan(&b)
		result = append(result, b)
	}

	var sum int
	for _, num := range result {
		if num > 0 {
			sum++
		}
	}

	fmt.Println(sum)
}
