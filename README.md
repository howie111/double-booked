# Double-booked
Given a sequence of events, each having a start and end time, write a program that will return the sequence of all pairs of overlapping events.

## General Idea
1. Sort events's start time in ascending order
2. Use Brute-force search to check if one event is conflicting with other events


## Unit test
Wrote different test cases to cover all scenarios:
1. Given empty events
2. Only one event in given events 
3. No overlapping events in given events
4. Normal multiple pairs of overlapping events(1-to-1 conflict)
5. Multiple pairs of overlapping events(1-to-2 conflict)

## Installation

Clone this repo and Run command
```
git clone https://github.com/howie111/double-booked
```

Run the main file
```bash
go run main.go
```




