package main

import "sort"

/*
identify the combinations of elements in an array of 10 elements whose sum is equal to 18. 
 The findElementsWithSum method recursively tries to find the combination
 */

func main () {
	
	 arr := []int{1,4,7,8,3,9,2,4,1,8}
	 sort.Ints(arr)


	findElementsWithSum(arr, 18, 0, [][]int{},0, 0)
	
}

func findElementsWithSum(elements []int, sumEqual int, currIteration int, combinations [][]int, currSum,  i int)  {
	//go back
	if currSum + elements[i] > sumEqual {
		combinations  = combinations[:currIteration-1]
		return 
	}

	if currSum + elements[i] == sumEqual {
		combinations[currIteration][i] = elements[i]
		return 
	}
	
	
	
	
	
}



