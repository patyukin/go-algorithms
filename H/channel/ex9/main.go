package main

import (
	"fmt"
	"time"
)

func main() {
	var input, e1, e2 string
	fmt.Scanln(&e1, &e2)
	input = e1 + " " + e2
	layout := "2006-01-02 15:04:05"
	eventTime, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println("Проблема с парсингом даты.")
		return
	}

	if eventTime.Hour() >= 13 {
		eventTime = eventTime.AddDate(0, 0, 1)
	}

	fmt.Println(eventTime.Format(layout))
}
