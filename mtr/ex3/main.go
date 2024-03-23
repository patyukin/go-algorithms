package main

func main() {
	s := make([]string, 0, 4)
	s = append(s, "a", "b", "c")

	add(s, "d")

	update(s, 3, "e")
}

func add(s []string, a string) []string {
	s = append(s, a)

	return s
}

func update(s []string, idx int, item string) {
	s[idx] = item
}
