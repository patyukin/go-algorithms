package main

import "fmt"

func main() {
	var a, temp, sum int
	var b []int
	fmt.Scanln(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&temp)
		b = append(b, temp)
	}

	for _, num := range b {
		if num < 100 && num > 10 && num%8 == 0 {
			sum += num
		}
	}

	fmt.Println(sum)
}
