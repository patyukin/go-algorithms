package main

import "fmt"

func main() {
	var b int
	a := -1
	m := make(map[int]int)

	for a != 0 {
		fmt.Scan(&a)
		if a > b {
			b = a
		}

		m[a]++
		if a == 0 {
			break
		}
	}

	fmt.Println(m[b])
}
