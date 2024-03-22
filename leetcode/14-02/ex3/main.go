package main

import (
	"fmt"
)

type I interface {
	Foo()
}

type S struct{}

func (s *S) Foo() {
	fmt.Println("foo")
}

func (s S) Bar() {
	fmt.Println("bar")
}

/*

var s S

s.Foo()

s1 := &S{}
s1.Foo() OK
s1.Bar() OK

var s2 *S
s2.Foo() OK
s2.Bar() - panic (*s).Bar() с точки зрения го

var s3 S
s3.Foo() - не скомпилируется
s3.Bar() ОК


s4 := &s3
s4.Foo() - ОК

s5 := *s1 - за тебя неявно делает гошка, если вызывается метод с НЕ указательном ресивером на указателе
s5.Bar() - OK
s5.Foo() - не компилируется

Го неявно умеет разыменовывать звездочки и работать с ними, как с конкретными типами: *T -> T
Но НЕ НАОБОРОТ. Го НИКОГДА не берёт указатели неявно. T -> &T не бывает.


type A struct {
    val int
}

a1 := A{}
a1.val = 5

a2 := &A{}
a2.val = 5

В мире C++:
a1.val = 5 OK

a2 -> val = 5 OK
*/

//iface {
//*itab // тип и методы
//unsafe.Pointer // указатель на конкретный объект/кусок памяти
//}

/*
i = {
    *itab -> *S
    ptr -> nil
}

i == nil -> i.itab == nil

var i I = nil

var s *S

var i2 I = s

fmt.Println(i == nil) // true
fmt.Println(i2 == nil) // false


*/

func Build() I {
	var res *S
	if res == nil {
		fmt.Println("here is nil")
	}

	return res
}

func main() {
	i := Build()
	// i - не нулевая структура
	if i != nil { // nil (*S) != nil
		i.Foo() // foo
	} else {
		fmt.Println("nil")
	}
}
