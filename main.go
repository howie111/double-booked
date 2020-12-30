package main

import (
	"log"
	"sort"
	"time"
)

type Event struct {
	Name  string
	Start time.Time
	End   time.Time
}

type OverlappedEvents []Event

/* sort event by start time in ascending order */
type SortByStartTime []Event

func (s SortByStartTime) Len() int {
	return len(s)
}

func (s SortByStartTime) Less(i, j int) bool {
	return s[i].Start.Sub(s[j].Start) < 0
}

func (s SortByStartTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// initialize time struct
func NewTime(year, month, day, hour, minute int) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
}

func isOverlapping(curEvent, nextEvent Event) bool {

	return curEvent.End.Sub(nextEvent.Start) > 0

}

func GetOverlappingEvents(events []Event) []OverlappedEvents {

	overlappedEvents := []OverlappedEvents{}

	if len(events) == 0 {
		return overlappedEvents
	}

	sort.Sort(SortByStartTime(events))

	for i, _ := range events {
		for j := i + 1; j < len(events); j++ {
			pairedEvent := []Event{}
			if isOverlapping(events[i], events[j]) {
				pairedEvent = append(pairedEvent, events[i])
				pairedEvent = append(pairedEvent, events[j])
			}
			if len(pairedEvent) > 0 {
				overlappedEvents = append(overlappedEvents, pairedEvent)
			}
		}
	}

	return overlappedEvents
}

func main() {

	events := []Event{
		Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
		Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 30)},
		Event{"event3", NewTime(2020, 12, 21, 10, 0), NewTime(2020, 12, 21, 11, 30)},
		Event{"event4", NewTime(2020, 12, 21, 11, 0), NewTime(2020, 12, 21, 12, 30)},
		Event{"event5", NewTime(2020, 12, 22, 10, 0), NewTime(2020, 12, 22, 11, 30)},
		Event{"event6", NewTime(2020, 12, 22, 11, 0), NewTime(2020, 12, 22, 12, 30)},
	}

	overlappedEvents := GetOverlappingEvents(events)

	if len(overlappedEvents) != 3 {
		log.Fatal()
	}

}
