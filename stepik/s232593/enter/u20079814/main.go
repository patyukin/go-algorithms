package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	m := array[0]
	for _, num := range array {
		if m < num {
			m = num
		}
	}

	fmt.Println(m)
}
