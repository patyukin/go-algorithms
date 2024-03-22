package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	s := v[:2]
	s = append(s, 10)
	fmt.Println(v)

	v2 := []int{1, 2, 3}
	s2 := v2[:2:2]
	s2 = append(s2, 10)
	fmt.Println(v2)
}
