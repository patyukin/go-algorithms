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

	fmt.Println(result[3])
}
