package main

import (
	"fmt"
	"math"
)

func main() {
	var p, y, i int
	var x float64

	fmt.Scanln(&x)
	fmt.Scanln(&p)
	fmt.Scanln(&y)

	for int(x) < y {
		x = x + (x*float64(p))/100
		x = math.Round(x*100) / 100
		i++
	}

	fmt.Println(i)
}
