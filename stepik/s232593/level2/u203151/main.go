package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(text)
	if unicode.IsUpper(runes[0]) && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
		return
	}

	fmt.Println("Wrong")
}
