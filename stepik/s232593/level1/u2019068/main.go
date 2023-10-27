package main

import (
	"fmt"
)

func main() {
	var minNum, num int
	fmt.Scanln(&num)
	m := make(map[int]int)

	for i := 0; i < num; i++ {
		var a int
		fmt.Scan(&a)
		if i == 0 {
			minNum = a
		}

		m[a]++
		if a < minNum {
			minNum = a
		}
	}

	fmt.Println(m[minNum])
}
