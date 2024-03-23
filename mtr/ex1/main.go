package main

func main() {
	Update("hello", "a", 1)
}

func Update(s, a string, id int) string {
	sl := []rune(s)
	sl[id] = []rune(a)[0]
	result := string(sl)

	return result
}
