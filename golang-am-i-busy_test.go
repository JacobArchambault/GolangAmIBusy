package main

import (
	"slices"
	"testing"
)

type Meeting struct {
	StartTime float64
	EndTime   float64
}

// ScheduleSimplifier sorts meetings in place by start time and merges
// overlapping or back-to-back meetings.
func ScheduleSimplifier(meetings []Meeting) []Meeting {
	if len(meetings) == 0 || len(meetings) == 1 {
		return meetings
	}

	// Sort meetings by start time
	slices.SortFunc(meetings, func(a, b Meeting) int {
		if a.StartTime < b.StartTime {
			return -1
		}
		if a.StartTime > b.StartTime {
			return 1
		}
		return 0
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
	if len(simplified) != 1 {
		t.Errorf("Expected 1 meeting for single meeting input")
	}
}

func TestOrderedNonOverlappingMeetings(t *testing.T) {
	meetings := []Meeting{
		{StartTime: 9.0, EndTime: 10.0},
		{StartTime: 11.0, EndTime: 12.0},
	}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 2 {
		t.Errorf("Expected 2 meetings for non-overlapping input")
	}
	if simplified[0].StartTime != 9.0 || simplified[1].StartTime != 11.0 || simplified[0].EndTime != 10.0 || simplified[1].EndTime != 12.0 {
		t.Errorf("Expected meetings to remain unchanged for non-overlapping input")
	}
}

func TestUnorderedNonOverlappingMeetings(t *testing.T) {
	meetings := []Meeting{
		{StartTime: 11.0, EndTime: 12.0},
		{StartTime: 9.0, EndTime: 10.0},
	}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 2 {
		t.Errorf("Expected 2 meetings for non-overlapping input")
	}
	if simplified[0].StartTime != 9.0 || simplified[1].StartTime != 11.0 || simplified[0].EndTime != 10.0 || simplified[1].EndTime != 12.0 {
		t.Errorf("Expected meetings to be sorted for unordered non-overlapping input")
	}
}

func TestOverlappingMeetings(t *testing.T) {
	meetings := []Meeting{
		{StartTime: 9.0, EndTime: 11.0},
		{StartTime: 10.0, EndTime: 12.0},
	}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 1 {
		t.Errorf("Expected 1 meeting for overlapping input")
	}
	if simplified[0].StartTime != 9.0 || simplified[0].EndTime != 12.0 {
		t.Errorf("Expected merged meeting to have start time 9.0 and end time 12.0 for overlapping input")
	}
}

func TestBackToBackMeetings(t *testing.T) {
	meetings := []Meeting{
		{StartTime: 10.0, EndTime: 11.0},
		{StartTime: 9.0, EndTime: 10.0},
	}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 1 {
		t.Errorf("Expected 1 meeting for back-to-back input")
	}
	if simplified[0].StartTime != 9.0 || simplified[0].EndTime != 11.0 {
		t.Errorf("Expected merged meeting to have start time 9.0 and end time 11.0 for back-to-back input")
	}
}

func TestComplexOverlappingMeetings(t *testing.T) {
	meetings := []Meeting{
		{StartTime: 10.5, EndTime: 12.0},
		{StartTime: 9.0, EndTime: 10.0},
		{StartTime: 9.5, EndTime: 11.0},
	}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 1 {
		t.Errorf("Expected 1 meeting for complex overlapping input")
	}
	if simplified[0].StartTime != 9.0 || simplified[0].EndTime != 12.0 {
		t.Errorf("Expected merged meeting to have start time 9.0 and end time 12.0 for complex overlapping input")
	}
}

func TestMultiplePartiallyOverlappingMeetings(t *testing.T) {
	meetings := []Meeting{
		{StartTime: 9.0, EndTime: 10.0},
		{StartTime: 9.5, EndTime: 11.0},
		{StartTime: 8.5, EndTime: 8.9},
	}
	simplified := ScheduleSimplifier(meetings)
	if len(simplified) != 2 {
		t.Errorf("Expected 2 meetings for multiple partially overlapping input")
	}
	if simplified[0].StartTime != 8.5 || simplified[0].EndTime != 8.9 {
		t.Errorf("Expected first meeting to have start time 8.5 and end time 8.9 for multiple partially overlapping input")
	}
	if simplified[1].StartTime != 9.0 || simplified[1].EndTime != 11.0 {
		t.Errorf("Expected second meeting to have start time 9.0 and end time 11.0 for multiple partially overlapping input")
	}
}
