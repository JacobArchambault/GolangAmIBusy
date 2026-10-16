package main

import (
	"sort"
	"testing"
)

type Meeting struct {
	StartTime float64
	EndTime   float64
}

func ScheduleSimplifier(meetings []Meeting) []Meeting {
	if len(meetings) == 0 || len(meetings) == 1 {
		return []Meeting{}
	}

	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].StartTime < meetings[j].StartTime
	})

	simplified := []Meeting{meetings[0]}

	for i := 1; i < len(meetings); i++ {
		last := &simplified[len(simplified)-1]
		current := meetings[i]

		if current.StartTime <= last.EndTime {
			// Merge overlapping meetings
			if current.EndTime > last.EndTime {
				last.EndTime = current.EndTime
			}
		} else {
			simplified = append(simplified, current)
		}
	}

	return simplified
}

func TestEmptySchedule(t *testing.T) {
	meetings := []Meeting{}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 0 {
		t.Errorf("Expected empty schedule for empty input")
	}
}

func TestSingleMeeting(t *testing.T) {
	meetings := []Meeting{{StartTime: 9.0, EndTime: 10.0}}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 0 {
		t.Errorf("Expected empty schedule for single meeting input")
	}
}
