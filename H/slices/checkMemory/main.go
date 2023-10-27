package main

import (
	"fmt"
)

func main() {
	var s []int
	for i := 0; i < 1024; i++ {
		s = append(s, i)
	}

	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 12)
	fmt.Println("-==- -==- -==- -==- -==-")
	fmt.Println(len(s))
	fmt.Println(cap(s))

	for i := 0; i < 512; i++ {
		s = append(s[:i], s[i+1:]...)
	}

	fmt.Println("-==- -==- -==- -==- -==-")
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
