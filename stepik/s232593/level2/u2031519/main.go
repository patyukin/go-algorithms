package main

import (
	"fmt"
	"strings"
)

func main() {
	var mainString, searchString string
	fmt.Scanln(&mainString)
	fmt.Scanln(&searchString)

	idx := strings.Index(mainString, searchString)
	fmt.Println(idx)
}
