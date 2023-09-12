package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("world!")

	fmt.Println(buffer.String()) // Выводит "Hello, world!"
}
