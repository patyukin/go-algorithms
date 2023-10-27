package main

import "fmt"

func main() {
	var a float64
	fmt.Scanln(&a)

	if a <= 0 {
		fmt.Printf("число %reflectionSliceHeader.2f не подходит", a)
		return
	}

	res := a * a
	if res > 10_000 {
		fmt.Printf("%e", a)
		return
	}

	b := fmt.Sprintf("%.4f", res)
	fmt.Println(b)
}
