package main

//	func main.go() {
//		m := map[rune]int{'a': 1, 'b': 2, 'c': 3}
//
//		for a, b := range m {
//			fmt.Println(a, b)
//		}
//	}
//func main.go() {
//	values := []int{1, 2, 3, 4, 5}
//	for _, val := range values {
//		go func(i int) {
//			fmt.Println(i) // 5 5 5 4 5
//		}(val)
//	}
//
//	time.Sleep(100 * time.Millisecond)
//}

// что выведет код?

//type X struct {
//	V int // 123
//}
//
//func (x X) S() {
//	fmt.Println(x.V)
//}
//
//func main.go() {
//	x := X{123}
//	defer func() {
//		x.S()
//	}()
//
//	x.V = 456
//}

type List struct {
	Val    int
	Next   *List
	Random *List
}

func copyList(root *List) *List {
	copied := &List{}
	head := root

	s := []int{}

	for head.Next != nil {
		copied.Val = head.Val
		copied.Next = &List{}
		s = append(s, head.Random.Val)
	}

	i := 0
	for copied.Next != nil {
		copied.Random = &List{Val: s[i]}
		i++
	}

	return copied
}
