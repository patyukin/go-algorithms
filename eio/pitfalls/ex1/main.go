package main

import (
	"errors"
	"fmt"
)

func main() {
	shadow()
}

func check1() (int, error) {
	return 1, nil //errors.New("SDSD")
}

func check2(a int) (int, error) {
	return a, errors.New("SDSD")
}

func shadow() (er error) {
	x, err := check1()
	if err != nil {
		return
	}
	if y, err := check2(x); err != nil {
		return
	} else {
		fmt.Println(y)
	}
	return
}
