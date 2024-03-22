package main

import "fmt"

func main() {
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
		var a uint8
		fmt.Scan(&a)
		workArray[i] = a
	}

	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	workArray[a], workArray[b] = workArray[b], workArray[a]
	workArray[c], workArray[d] = workArray[d], workArray[c]
	workArray[f], workArray[e] = workArray[e], workArray[f]

	for _, val := range workArray {
		fmt.Print(val, " ")
	}
}
