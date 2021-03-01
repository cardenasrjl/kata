package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{5,7,7,7}
	result := solve(input, 2)
	fmt.Printf("%t", result)
}


func solve(data []int, k int) bool {
	if !(len(data) >= k) {
		return false
	}
	
	sort.Ints(data)
	result := make([]int,k)
	index := 0
	possible := true
	
	retries := 0
	for i:=0;i<len(data);i++{
		if result[index] < data[i] {
			result[index] = data[i]
		} else {
			retries++
			i--
		}
		
		index++
		if index >= k {
			index=0
		}
		
		if retries >= k {
			possible = false
			break
		}
		
	}
	
	return possible
}
