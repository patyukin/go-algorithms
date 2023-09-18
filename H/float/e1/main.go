package main

import (
	"fmt"
	"math"
)

func main() {
	q := 1.23
	w := 1.229999

	q2 := int(math.Round(q * 100))
	w2 := int(math.Round(w * 100))

	fmt.Println(q2 == w2)
}
