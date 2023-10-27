package main

import (
	"fmt"
	"strings"
)

func main() {
	var s strings.Builder

	s.Write([]byte("a"))
	s.Write([]byte("b"))
	s.Write([]byte("c"))

	fmt.Println(s.String())
}
