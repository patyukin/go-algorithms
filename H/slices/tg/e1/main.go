package main

import (
	"fmt"
)

func main() {
	a := []int{11, 22, 33}
	b := a[0:1:3]
	fmt.Println(b[2:3])
}
