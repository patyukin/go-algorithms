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

	for i, num := range result {
		if i%2 == 0 {
			fmt.Print(num, " ")
		}
	}
}
