package main

import (
	"fmt"
)

const read = "read"
const write = "write"
const del = "delete"

type SafeChannelMap struct {
	operationType string
	key           string
	value         int
}

func safeMapOperations(ch chan SafeChannelMap, mapValue chan int) {
	for {
		select {
		case msg := <-ch:
			if msg.operationType == write {
				m[msg.key] = msg.value
			} else if msg.operationType == del {
				delete(m, msg.key)
			} else if msg.operationType == read {
				mapValue <- m[msg.key]
			}
		}
	}
}

var m = make(map[string]int)

func main() {
	ch := make(chan SafeChannelMap)
	mapValue := make(chan int)
	go safeMapOperations(ch, mapValue)

	ch <- SafeChannelMap{
		operationType: write,
		key:           "key",
		value:         1,
	}

	ch <- SafeChannelMap{
		operationType: read,
		key:           "key",
	}

	fmt.Println(<-mapValue)

	ch <- SafeChannelMap{
		operationType: del,
		key:           "key",
	}

	close(ch)
	close(mapValue)
}
