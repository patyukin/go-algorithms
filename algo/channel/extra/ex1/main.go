package main

import (
	"fmt"
	"regexp"
	"strings"
)

func clearWhiteSpaces(cs <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		r := regexp.MustCompile("\\s+")
		for v := range cs {
			out <- r.ReplaceAllString(v, " ")
		}
		close(out)
	}()

	return out
}

func deleteVowelLetter(s <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		r := regexp.MustCompile("[AEIOUY]")
		for v := range s {
			out <- r.ReplaceAllString(v, "#")
		}
		close(out)
	}()

	return out
}

func uppercase(s []string) <-chan string {
	out := make(chan string)

	go func() {
		for _, v := range s {
			out <- strings.ToUpper(v)
		}
		close(out)
	}()

	return out
}

func consume(s <-chan string) chan string {
	out := make(chan string)

	go func() {
		for v := range s {
			out <- v
		}
		close(out)
	}()

	return out
}

func main() {
	s := uppercase([]string{"hello", "world", "go", "channels"})
	s = clearWhiteSpaces(s)
	s = deleteVowelLetter(s)
	out := consume(s)

	var result []string
	for v := range out {
		result = append(result, v)
	}

	fmt.Println(result)
}
