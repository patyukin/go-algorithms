package main

import "fmt"

type A struct {
	Val int
}

type I interface {
	GetValue()
}

func (f A) GetValue() {
	// f == nil
	// return f.Val
	fmt.Println("panic")
}

func getA() I {
	var result *A

	return result
}

//type iface struct {
//	ptr1 // методы (I) => структура (A)
//	ptr2 // unsafe.Pointer (A)
//}

func main() {
	i := getA() // iface{ptr2: &A, ptr1: (not nil)}
	println(i != nil)

	i.GetValue()
}
