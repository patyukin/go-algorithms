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
	reader        chan int
}

func safeMapOperations(ch chan SafeChannelMap) {
	for {
		select {
		case msg := <-ch:
			if msg.operationType == write {
				m[msg.key] = msg.value
			} else if msg.operationType == del {
				delete(m, msg.key)
			} else if msg.operationType == read {
				msg.reader <- m[msg.key]
			}
		}
	}
}

var m = make(map[string]int)

func main() {
	ch := make(chan SafeChannelMap)
	go safeMapOperations(ch)

	writer := SafeChannelMap{
		operationType: write,
		key:           "key",
		value:         100,
	}

	reader := SafeChannelMap{
		operationType: read,
		key:           "key",
		reader:        make(chan int),
	}

	writer.value = 1000
	ch <- writer
	ch <- reader
	fmt.Println(<-reader.reader)

	writer.value = 1 << 12
	ch <- writer
	ch <- reader
	fmt.Println(<-reader.reader)

	ch <- SafeChannelMap{
		operationType: del,
		key:           "key",
	}

	ch <- reader
	fmt.Println(<-reader.reader)

	close(ch)
}
