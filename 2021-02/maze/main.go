package main

import (
	
	"fmt"
)

func main() {
	
	//maze := [][]int{{0,1},{0,3},{1,2},{1,0},{2,1},{3,0},{3,6},{3,4},{6,3},{4,3},{4,5},{4,7},{5,4},{7,4},{7,8},{8,7}}
	maze := map[int][]int{0:{1,3},1:{0,2},2:{1},3:{0,4,6},6:{3},4:{3,7,5},5:{4},7:{4,8},8:{7}}
	
	path := []int{}

	solution, path := solve(maze, 0,8, map[int]bool{})
	
	if solution {
		fmt.Printf("the solution is %+v", path)
	}
	
}


func solve(maze map[int][]int, start, end int, visited map[int]bool) (bool,  []int) {
	if visited[start] {
		return false, nil
	}
	
	if start == end {
		return true,[]int{end}
	}

	visited[start] = true
	for _,v := range maze[start] {
		sol, p :=  solve(maze,v, end, visited)
		if sol {
			p = append(p, start)
			
			return true,p
		}
	}
	
	 return false, nil
}