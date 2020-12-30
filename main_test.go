package main

import (
	"sort"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMain(t *testing.T) {

	t.Run("SortEvent", func(t *testing.T) {

		events := []Event{
			Event{"event1", NewTime(2020, 12, 20, 12, 0), NewTime(2020, 12, 20, 12, 30)},
			Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 11, 30)},
		}

		expected := []Event{
			Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 11, 30)},
			Event{"event1", NewTime(2020, 12, 20, 12, 0), NewTime(2020, 12, 20, 12, 30)},
		}

		sort.Sort(SortByStartTime(events))

		assert.Equal(t, expected, events)

	})

	t.Run("EmptyEvent", func(t *testing.T) {
		events := []Event{}

		overlappedEvents := GetOverlappingEvents(events)
		assert.Equal(t, 0, len(overlappedEvents))

	})
	t.Run("OneEvent", func(t *testing.T) {
		events := []Event{
			Event{"event1", NewTime(2020, 12, 20, 12, 0), NewTime(2020, 12, 20, 12, 30)},
		}

		overlappedEvents := GetOverlappingEvents(events)

		assert.Equal(t, 0, len(overlappedEvents))
	})

	t.Run("NoOverlappingEvents", func(t *testing.T) {

		events := []Event{
			Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
			Event{"event2", NewTime(2020, 12, 20, 14, 0), NewTime(2020, 12, 20, 14, 30)},
		}

		overlappedEvents := GetOverlappingEvents(events)

		assert.Equal(t, 0, len(overlappedEvents))

	})
	t.Run("OnePairGetOverlappingEvents", func(t *testing.T) {

		events := []Event{
			Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
			Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 30)},
		}

		overlappedEvents := GetOverlappingEvents(events)

		assert.Equal(t, 1, len(overlappedEvents))

	})

	t.Run("MultiplePairsOfGetOverlappingEventsOneToOne", func(t *testing.T) {

		events := []Event{
			Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
			Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 30)},
			Event{"event3", NewTime(2020, 12, 21, 10, 0), NewTime(2020, 12, 21, 11, 30)},
			Event{"event4", NewTime(2020, 12, 21, 11, 0), NewTime(2020, 12, 21, 12, 30)},
			Event{"event5", NewTime(2020, 12, 22, 10, 0), NewTime(2020, 12, 22, 11, 30)},
			Event{"event6", NewTime(2020, 12, 22, 11, 0), NewTime(2020, 12, 22, 12, 30)},
		}

		expected := []OverlappedEvents{
			[]Event{
				Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
				Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 30)},
			},
			[]Event{
				Event{"event3", NewTime(2020, 12, 21, 10, 0), NewTime(2020, 12, 21, 11, 30)},
				Event{"event4", NewTime(2020, 12, 21, 11, 0), NewTime(2020, 12, 21, 12, 30)},
			},
			[]Event{
				Event{"event5", NewTime(2020, 12, 22, 10, 0), NewTime(2020, 12, 22, 11, 30)},
				Event{"event6", NewTime(2020, 12, 22, 11, 0), NewTime(2020, 12, 22, 12, 30)},
			},
		}

		overlappedEvents := GetOverlappingEvents(events)

		assert.Equal(t, expected, overlappedEvents)

	})

	t.Run("MultiplePairsOfGetOverlappingEventsOneToTwo", func(t *testing.T) {

		events := []Event{
			Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
			Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 0)},
			Event{"event3", NewTime(2020, 12, 20, 11, 15), NewTime(2020, 12, 20, 13, 30)},
		}

		expected := []OverlappedEvents{
			[]Event{
				Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
				Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 0)},
			},
			[]Event{
				Event{"event1", NewTime(2020, 12, 20, 10, 0), NewTime(2020, 12, 20, 11, 30)},
				Event{"event3", NewTime(2020, 12, 20, 11, 15), NewTime(2020, 12, 20, 13, 30)},
			},
			[]Event{
				Event{"event2", NewTime(2020, 12, 20, 11, 0), NewTime(2020, 12, 20, 12, 0)},
				Event{"event3", NewTime(2020, 12, 20, 11, 15), NewTime(2020, 12, 20, 13, 30)},
			},
		}

		overlappedEvents := GetOverlappingEvents(events)

		assert.Equal(t, expected, overlappedEvents)

	})

}
