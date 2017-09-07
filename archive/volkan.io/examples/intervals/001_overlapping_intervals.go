package main

import (
	"sort"
	"fmt"
)

type Interval []int

type ByStartTime []Interval

func (t ByStartTime) Len() int {
	return len(t)
}

func (t ByStartTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByStartTime) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	if intervals == nil {
		return false
	}

	sort.Sort(ByStartTime(intervals))

	for i, interval := range intervals {
		if i > 0 {
			if interval[0] < intervals[i-1][1] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(
		canAttendMeetings([]Interval{}),
	)
}