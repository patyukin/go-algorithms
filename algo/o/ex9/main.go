package main

import (
	"fmt"
	"sort"
)

// дан массив meetings, в котором каждый элемент meeting[i] - это пара чисел [start, end],
// объединить все накладывающиеся друг на друга встречи и вернуть массив с объединенными встречами, покрывающих те же временные слоты

// input meetings = [[0, 1], [3, 5], [4, 8], [10, 12], [9, 10]]
// output [[0, 1], [3, 8], [9, 12]]

// https://leetcode.com/problems/meeting-rooms-ii

func mergeMeetings(meetings [][]int) [][]int {
	if len(meetings) <= 1 {
		return meetings
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	mergedMeetings := [][]int{meetings[0]}

	for i := 1; i < len(meetings); i++ {
		currentMeeting := meetings[i]
		lastMerged := mergedMeetings[len(mergedMeetings)-1]

		if currentMeeting[0] <= lastMerged[1] {
			lastMerged[1] = max(currentMeeting[1], lastMerged[1])
		} else {
			mergedMeetings = append(mergedMeetings, currentMeeting)
		}
	}

	return mergedMeetings
}

func main() {
	meetings := [][]int{{0, 1}, {3, 5}, {4, 8}, {10, 12}, {9, 10}}
	fmt.Println(mergeMeetings(meetings))
}
