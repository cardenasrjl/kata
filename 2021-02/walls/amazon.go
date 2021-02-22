package main

import (
	"fmt"
	"math"
)

/*
You are given an M by N matrix consisting of booleans that represents a board. Each True boolean represents a wall. Each False boolean represents a tile you can walk on.

Given this matrix, a start coordinate, and an end coordinate, return the minimum number of steps required to reach the end coordinate from the start. If there is no possible path, then return null. You can move up, left, down, and right. You cannot move through walls. You cannot wrap around the edges of the board.

For example, given the following board:

[[f, f, f, f],
[t, t, f, t],
[f, f, f, f],
[f, f, f, f]]
and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number of steps required to reach the end is 7, since we would need to go through (1, 2) because there is a wall everywhere else on the second row.
*/

 
func main() {

	

	m := [][]bool{
		{true, true, true, true},
		{true, true, true, true},
		{true, false, true, false},
		{true, true, true, false},
	}

	var visited [4][4]bool
	mi, found := walk(m, 0, 3, 3, 2, visited)

	fmt.Printf("found %t, minimum %+d", found, mi)

}

func walk(mtx [][]bool, currRow, currCol, desiredRow, desiredCol int, visited [4][4]bool) (int, bool) {

	if !(currRow >= 0 && currRow < len(mtx) && currCol >= 0 && currCol < len(mtx[0]) && mtx[currRow][currCol] && !visited[currRow][currCol]) {
		return 0, false
	}

	visited[currRow][currCol] = true

	if currCol == desiredCol && currRow == desiredRow {
		return 1, true
	}

	a, u := walk(mtx, currRow, currCol-1, desiredRow, desiredCol, visited)
	b, d := walk(mtx, currRow, currCol+1, desiredRow, desiredCol, visited)
	c, l := walk(mtx, currRow-1, currCol, desiredRow, desiredCol, visited)
	e, r := walk(mtx, currRow+1, currCol, desiredRow, desiredCol, visited)

	result := []int{}
	if u {
		result = append(result, a)
	}

	if d {
		result = append(result, b)
	}

	if l {
		result = append(result, c)
	}

	if r {
		result = append(result, e)
	}
	min := math.MinInt64
	for i, e := range result {
		if i == 0 || e < min {
			min = e
		}
	}

	return min + 1, u || d || l || r

}
