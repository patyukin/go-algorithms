package main

import (
	"fmt"
)

type III interface {
}

func main() {
	m := make(map[III]int, 100)
	m[4] = 2
	_ = m[1]
	for i := 0; i < 1000; i++ {
		if i > 31 {
			m[i] = i
		}
		m[i] = i
	}
	fmt.Println(len(m))
	_ = m[1]
}
