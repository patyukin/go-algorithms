package main

func main() {
	m := make([]int64, 0, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}

	clear(m)

	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}
	m = make([]int64, 0)

	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}
	m = nil
}
