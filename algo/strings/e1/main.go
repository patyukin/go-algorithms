package main

import (
	"fmt"
)

func main() {
	s := `⓴`

	for _, v := range s {
		fmt.Printf("%[1]d, %+[1]q\n", v)
	}
	for _, v := range []rune(s) {
		fmt.Printf("%[1]d, %+[1]q\n", v)
	}

	fmt.Println()
	fmt.Println([]byte(s))

	for _, v := range []byte(s) {
		fmt.Printf("%[1]d, %+[1]q\n", v)
	}
}
