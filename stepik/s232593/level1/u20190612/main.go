package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scanln(&a)
	if a == 0 {
		return
	}

	var i float64
	for ; i <= a; i++ {
		pow := math.Pow(2, i)
		if pow > a {
			return
		}
		fmt.Printf("%.0f ", pow)
	}
}
