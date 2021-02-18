package main

import "fmt"

/**

Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

 */


func main() {
	
	interval := []Interval{
		{
			Start: 30,
			End:   75,
		},
		{
			Start: 0,
			End:   50,
		},
		{
			Start: 20,
			End:   50,
		},
	}

	fmt.Printf("Number of rooms needed %d", solution(interval))

}

type Interval struct {
	Start int
	End   int
}

func solution(schedules []Interval) int {
	numberRooms := 0
	overlapping := map[int]int{}
	for _, schedule := range schedules {
		val := fillOverlapping(overlapping, schedule)
		if val > numberRooms {
			numberRooms = val
		}
	}
	return numberRooms
}

func fillOverlapping(overlapping map[int]int, interval Interval) int {
	currValue := 0
	for i := interval.Start; i <= interval.End; i += 5 {
		overlapping[i]++
		if currValue < overlapping[i] {
			currValue = overlapping[i]
		}
	}
	return currValue
}
