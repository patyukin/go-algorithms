package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomMatrix(n, m int) [][]int {
	var res [][]int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	mp := make(map[int]struct{})

	var generateNumber func() int
	generateNumber = func() int {
		randNumber := r.Int()
		if _, ok := mp[randNumber]; ok {
			return generateNumber()
		}

		mp[randNumber] = struct{}{}
		return randNumber
	}

	for i := 0; i < n; i++ {
		res = append(res, make([]int, m))
		for j := 0; j < m; j++ {
			res[i][j] = generateNumber()
		}
	}

	return res
}

func main() {
	fmt.Println(generateRandomMatrix(3, 3))
}
